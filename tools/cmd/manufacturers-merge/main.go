package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	mergedData = make(map[string]*vinv1.Manufacturer)
)

func main() {
	wikipediaFile := flag.String("wikipedia", "", "Path to Wikipedia manufacturers.jsonl")
	wikibooksFile := flag.String("wikibooks", "", "Path to Wikibooks manufacturers.jsonl")
	vpicFile := flag.String("vpic", "", "Path to vPIC manufacturers.jsonl")
	kbaFile := flag.String("kba", "", "Path to KBA manufacturers.jsonl")
	outputFile := flag.String("output", "data/manufacturers.jsonl", "Path to the output JSONL file")
	flag.Parse()

	if *wikipediaFile == "" || *wikibooksFile == "" || *vpicFile == "" || *kbaFile == "" {
		log.Fatal("All input files (wikipedia, wikibooks, vpic, kba) must be specified.")
	}

	// Load sources in precedence order (Highest -> Lowest)
	// 1. KBA (Highest)
	if err := load(*kbaFile, vinv1.DataSource_KBA); err != nil {
		log.Fatalf("Failed to load KBA data: %v", err)
	}

	// 2. VPIC
	if err := load(*vpicFile, vinv1.DataSource_VPIC); err != nil {
		log.Fatalf("Failed to load vPIC data: %v", err)
	}

	// 3. Wikibooks
	if err := load(*wikibooksFile, vinv1.DataSource_WIKIBOOKS); err != nil {
		log.Fatalf("Failed to load Wikibooks data: %v", err)
	}

	// 4. Wikipedia (Lowest)
	if err := load(*wikipediaFile, vinv1.DataSource_WIKIPEDIA); err != nil {
		log.Fatalf("Failed to load Wikipedia data: %v", err)
	}

	// Write output
	if err := writeOutput(*outputFile); err != nil {
		log.Fatalf("Failed to write output: %v", err)
	}

	log.Printf("Successfully merged %d manufacturers to %s", len(mergedData), *outputFile)
}

func load(path string, source vinv1.DataSource) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		var record vinv1.Manufacturer
		if err := protojson.Unmarshal([]byte(line), &record); err != nil {
			log.Printf("Skipping invalid JSON in %s: %v", path, err)
			continue
		}

		// *** VALIDATION ***
	wmi1 := record.GetWmi1()
		if !isValidWMI(wmi1) {
			return fmt.Errorf("WMI1 '%s' in %s is not normalized (must be 3 uppercase alphanumeric chars)", wmi1, path)
		}
	wmi2 := record.GetWmi2()
		if wmi2 != "" && !isValidWMI(wmi2) {
			return fmt.Errorf("WMI2 '%s' in %s is not normalized (must be 3 uppercase alphanumeric chars)", wmi2, path)
		}

		// Ensure new record has the correct data source set (if missing)
	hasSource := false
		dataSources := record.GetDataSources()
		for _, s := range dataSources {
			if s == source {
				hasSource = true
				break
			}
		}
		if !hasSource {
			dataSources = append(dataSources, source)
			record.SetDataSources(dataSources)
		}

		key := getManufacturerKey(&record)
		existing, exists := mergedData[key]

		if !exists {
			mergedData[key] = &record
			count++
		} else {
			mergeProto(existing, &record, source)
		}
	}

	log.Printf("Loaded %d records from %s", count, source)
	return scanner.Err()
}

// isValidWMI checks if a WMI string is 3 uppercase alphanumeric characters
func isValidWMI(wmi string) bool {
	if len(wmi) != 3 {
		return false
	}
	for _, r := range wmi {
		if !((r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

func getManufacturerKey(m *vinv1.Manufacturer) string {
	wmi2 := m.GetWmi2()
	if wmi2 != "" {
		return m.GetWmi1() + ":" + wmi2
	}
	return m.GetWmi1()
}

func mergeProto(existing, new *vinv1.Manufacturer, source vinv1.DataSource) {
	// Snapshot for comparison - create a clone
	snapshot := proto.Clone(existing).(*vinv1.Manufacturer)

	// Update scalars (fill gaps only - existing takes precedence)
	if existing.GetDisplayName() == "" && new.GetDisplayName() != "" {
		existing.SetDisplayName(new.GetDisplayName())
	}
	if existing.GetCountry() == vinv1.Country_COUNTRY_UNSPECIFIED && new.GetCountry() != vinv1.Country_COUNTRY_UNSPECIFIED {
		existing.SetCountry(new.GetCountry())
	}
	if existing.GetRegion() == vinv1.Region_REGION_UNSPECIFIED && new.GetRegion() != vinv1.Region_REGION_UNSPECIFIED {
		existing.SetRegion(new.GetRegion())
	}
	if existing.GetKbaId() == 0 && new.GetKbaId() != 0 {
		existing.SetKbaId(new.GetKbaId())
	}
	if existing.GetVpicId() == 0 && new.GetVpicId() != 0 {
		existing.SetVpicId(new.GetVpicId())
	}
	// For bools, we only set if currently false and new is true
	if !existing.GetLowVolume() && new.GetLowVolume() {
		existing.SetLowVolume(true)
	}

	// Merge repeated fields (Additive)
	existing.SetBrands(mergeBrands(existing.GetBrands(), new.GetBrands()))
	existing.SetVehicleTypes(mergeVehicleTypes(existing.GetVehicleTypes(), new.GetVehicleTypes()))

	// Check for changes *excluding* DataSources
	// We temporarily clear DataSources on both to compare content
	existingDS := existing.GetDataSources()
	existing.SetDataSources(nil)
	snapshot.SetDataSources(nil)

	changed := !proto.Equal(existing, snapshot)

	// Restore DataSources
	existing.SetDataSources(existingDS)

	if changed {
		// If content changed, ensure the source is added
		addSource(existing, source)
	}
}

func addSource(m *vinv1.Manufacturer, source vinv1.DataSource) {
	dataSources := m.GetDataSources()
	for _, s := range dataSources {
		if s == source {
			return
		}
	}
	dataSources = append(dataSources, source)
	m.SetDataSources(dataSources)
}

func mergeBrands(a, b []vinv1.Brand) []vinv1.Brand {
	seen := make(map[vinv1.Brand]bool)
	var res []vinv1.Brand
	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			res = append(res, v)
		}
	}
	for _, v := range b {
		if !seen[v] {
			seen[v] = true
			res = append(res, v)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func mergeVehicleTypes(a, b []vinv1.VehicleType) []vinv1.VehicleType {
	seen := make(map[vinv1.VehicleType]bool)
	var res []vinv1.VehicleType
	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			res = append(res, v)
		}
	}
	for _, v := range b {
		if !seen[v] {
			seen[v] = true
			res = append(res, v)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func writeOutput(path string) error {
	// Sort manufacturers by key for deterministic output
	keys := make([]string, 0, len(mergedData))
	for k := range mergedData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	marshalOpts := protojson.MarshalOptions{Multiline: false}
	for _, k := range keys {
		m := mergedData[k]

		// 2. Omit manufacturers (wmi and lvm) that don't have any brands inferred
		if len(m.GetBrands()) == 0 {
			continue
		}

		// 1. Omit the manufacturer name (we don't have much use for it anyway, we only care about the structured metadata)
		m.ClearDisplayName()

		data, err := marshalOpts.Marshal(m)
		if err != nil {
			return err
		}
		if _, err := writer.Write(data); err != nil {
			return err
		}
		if _, err := writer.WriteString("\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}
