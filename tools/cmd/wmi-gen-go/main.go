package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/way-platform/vin-go/internal/codegen"
	"github.com/way-platform/vin-go/internal/wmi"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// Manufacturer is the input struct, matching the output of manufacturers-merge
type Manufacturer struct {
	WMI1         string   `json:"wmi1"`
	WMI2         string   `json:"wmi2,omitempty"`
	DisplayName  string   `json:"displayName"`
	Country      string   `json:"country,omitempty"`
	Region       string   `json:"region,omitempty"`
	Brands       []string `json:"brands,omitempty"`
	VehicleTypes []string `json:"vehicleTypes,omitempty"`
	KBAID        int      `json:"kbaId,omitempty"`     // Not directly used in Proto, but might be useful
	VPICID       int      `json:"vpicId,omitempty"`    // Not directly used in Proto
	LowVolume    bool     `json:"lowVolume,omitempty"` // Handled by WMI1[2] == '9'
	DataSources  []string `json:"dataSources,omitempty"`
	URL          string   `json:"url,omitempty"`
	Note         string   `json:"note,omitempty"`
}

// IndexEntry for standard manufacturers
type StandardIndexEntry struct {
	K uint16 // Base36(WMI1) - Key
	O uint32 // Offset in wmi.bin
}

// IndexEntry for LVMs
type LVMIndexEntry struct {
	K uint32 // Packed Key (WMI1 << 16 | WMI2)
	O uint32 // Offset in lvm.bin
}

func main() {
	inputFile := flag.String("input", "data/manufacturers.jsonl", "Path to the merged manufacturers JSONL file")
	outputWMIFile := flag.String("output-wmi", "wmi.gen.go", "Path to the output Go file for WMI index")
	outputLVMFile := flag.String("output-lvm", "lvm.gen.go", "Path to the output Go file for LVM index")
	outputWMIBin := flag.String("output-wmi-bin", "wmi.bin", "Path to the output binary file for standard WMIs")
	outputLVMBin := flag.String("output-lvm-bin", "lvm.bin", "Path to the output binary file for LVMs")
	packageName := flag.String("package", "vin", "Package name for generated Go code")
	flag.Parse()

	if err := run(*inputFile, *outputWMIFile, *outputLVMFile, *outputWMIBin, *outputLVMBin, *packageName); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func run(inputFile, outputWMIFile, outputLVMFile, outputWMIBin, outputLVMBin, packageName string) error {
	manufacturers, err := readManufacturers(inputFile)
	if err != nil {
		return fmt.Errorf("reading manufacturers: %w", err)
	}

	standardManufacturers, lvmManufacturers := splitManufacturers(manufacturers)

	// Generate binary blobs
	_, standardIndex, err := generateStandardBlobAndIndex(standardManufacturers, outputWMIBin)
	if err != nil {
		return fmt.Errorf("generating WMI blob: %w", err)
	}

	_, lvmIndex, err := generateLVMBlobAndIndex(lvmManufacturers, outputLVMBin)
	if err != nil {
		return fmt.Errorf("generating LVM blob: %w", err)
	}

	// Generate separate Go files
	if err := generateWMICode(outputWMIFile, packageName, standardIndex, outputWMIBin); err != nil {
		return fmt.Errorf("generating WMI code: %w", err)
	}

	if err := generateLVMCode(outputLVMFile, packageName, lvmIndex, outputLVMBin); err != nil {
		return fmt.Errorf("generating LVM code: %w", err)
	}

	return nil
}

func readManufacturers(filename string) ([]*Manufacturer, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var manufacturers []*Manufacturer
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		var m Manufacturer
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			return nil, fmt.Errorf("unmarshalling JSONL line: %w", err)
		}
		manufacturers = append(manufacturers, &m)
	}
	return manufacturers, scanner.Err()
}

func splitManufacturers(manufacturers []*Manufacturer) ([]*Manufacturer, []*Manufacturer) {
	var standard []*Manufacturer
	var lvm []*Manufacturer

	for _, m := range manufacturers {
		if m.WMI2 == "" {
			standard = append(standard, m)
		} else {
			lvm = append(lvm, m)
		}
	}

	// Sort standard by Base36(WMI1)
	sort.Slice(standard, func(i, j int) bool {
		keyI, _ := wmi.ToBase36(standard[i].WMI1)
		keyJ, _ := wmi.ToBase36(standard[j].WMI1)
		return keyI < keyJ
	})

	// Sort LVM by Packed Key
	sort.Slice(lvm, func(i, j int) bool {
		kI, _ := wmi.Pack(lvm[i].WMI1, lvm[i].WMI2)
		kJ, _ := wmi.Pack(lvm[j].WMI1, lvm[j].WMI2)
		return kI < kJ
	})

	return standard, lvm
}

func generateStandardBlobAndIndex(manufacturers []*Manufacturer, outputPath string) ([]byte, []StandardIndexEntry, error) {
	var buf bytes.Buffer
	var standardIndex []StandardIndexEntry

	for _, m := range manufacturers {
		protoMan := toProto(m)
		data, err := proto.Marshal(protoMan)
		if err != nil {
			return nil, nil, fmt.Errorf("marshalling proto: %w", err)
		}

		// Record offset before writing
		currentOffset := uint32(buf.Len())

		key, ok := wmi.ToBase36(m.WMI1)
		if !ok {
			return nil, nil, fmt.Errorf("invalid WMI1 for base36 conversion: %s", m.WMI1)
		}
		standardIndex = append(standardIndex, StandardIndexEntry{K: key, O: currentOffset})

		if _, err := buf.Write(data); err != nil {
			return nil, nil, fmt.Errorf("writing to buffer: %w", err)
		}
	}

	// Write blob to file
	if err := os.WriteFile(outputPath, buf.Bytes(), 0o644); err != nil {
		return nil, nil, fmt.Errorf("writing binary blob to %s: %w", outputPath, err)
	}

	return buf.Bytes(), standardIndex, nil
}

func generateLVMBlobAndIndex(manufacturers []*Manufacturer, outputPath string) ([]byte, []LVMIndexEntry, error) {
	var buf bytes.Buffer
	var lvmIndex []LVMIndexEntry

	for _, m := range manufacturers {
		protoMan := toProto(m)
		data, err := proto.Marshal(protoMan)
		if err != nil {
			return nil, nil, fmt.Errorf("marshalling proto: %w", err)
		}

		// Record offset before writing
		currentOffset := uint32(buf.Len())

		k, ok := wmi.Pack(m.WMI1, m.WMI2)
		if !ok {
			return nil, nil, fmt.Errorf("invalid WMI for base36 conversion: %s%s", m.WMI1, m.WMI2)
		}
		lvmIndex = append(lvmIndex, LVMIndexEntry{K: k, O: currentOffset})

		if _, err := buf.Write(data); err != nil {
			return nil, nil, fmt.Errorf("writing to buffer: %w", err)
		}
	}

	// Write blob to file
	if err := os.WriteFile(outputPath, buf.Bytes(), 0o644); err != nil {
		return nil, nil, fmt.Errorf("writing binary blob to %s: %w", outputPath, err)
	}

	return buf.Bytes(), lvmIndex, nil
}

func toProto(m *Manufacturer) *vinv1.Manufacturer {
	protoMan := &vinv1.Manufacturer{}

	protoMan.SetDisplayName(m.DisplayName)
	protoMan.SetWmi1(m.WMI1)
	protoMan.SetLowVolume(m.LowVolume || m.WMI2 != "")

	if m.WMI2 != "" {
		protoMan.SetWmi2(m.WMI2)
	}

	if m.VPICID != 0 {
		protoMan.SetVpicId(int32(m.VPICID))
	}

	if m.KBAID != 0 {
		protoMan.SetKbaId(int32(m.KBAID))
	}

	if m.Country != "" {
		protoMan.SetCountry(convertCountry(m.Country))
	}

	if m.Region != "" {
		protoMan.SetRegion(convertRegion(m.Region))
	}

	// Convert Brands from string to enum slice
	var brands []vinv1.Brand
	for _, b := range m.Brands {
		if brand := convertBrand(b); brand != vinv1.Brand_BRAND_UNSPECIFIED {
			brands = append(brands, brand)
		}
	}
	if len(brands) > 0 {
		protoMan.SetBrands(brands)
	}

	// Convert VehicleTypes from string to enum slice
	var vehicleTypes []vinv1.VehicleType
	for _, vt := range m.VehicleTypes {
		if vtEnum := convertVehicleType(vt); vtEnum != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
			vehicleTypes = append(vehicleTypes, vtEnum)
		}
	}
	if len(vehicleTypes) > 0 {
		protoMan.SetVehicleTypes(vehicleTypes)
	}

	// Convert DataSources from string to enum slice
	var dataSources []vinv1.DataSource
	for _, ds := range m.DataSources {
		if dsEnum := convertDataSource(ds); dsEnum != vinv1.DataSource_DATA_SOURCE_UNSPECIFIED {
			dataSources = append(dataSources, dsEnum)
		}
	}
	if len(dataSources) > 0 {
		protoMan.SetDataSources(dataSources)
	}

	// Models field is currently empty in input data, but prepared for future use
	// No conversion needed as Models would come as enum strings if present

	return protoMan
}

// Conversion functions convert enum strings to proto enum values.
// Enum names in JSONL are already in the correct format (e.g., "THE_NETHERLANDS", "BUS", "VPIC").
func convertCountry(s string) vinv1.Country {
	val, ok := vinv1.Country_value[s]
	if !ok {
		log.Printf("warning: unmappable country value: %s", s)
		return vinv1.Country_COUNTRY_UNSPECIFIED
	}
	return vinv1.Country(val)
}

func convertRegion(s string) vinv1.Region {
	val, ok := vinv1.Region_value[s]
	if !ok {
		log.Printf("warning: unmappable region value: %s", s)
		return vinv1.Region_REGION_UNSPECIFIED
	}
	return vinv1.Region(val)
}

func convertBrand(s string) vinv1.Brand {
	val, ok := vinv1.Brand_value[s]
	if !ok {
		log.Printf("warning: unmappable brand value: %s", s)
		return vinv1.Brand_BRAND_UNSPECIFIED
	}
	return vinv1.Brand(val)
}

func convertVehicleType(s string) vinv1.VehicleType {
	val, ok := vinv1.VehicleType_value[s]
	if !ok {
		log.Printf("warning: unmappable vehicle type value: %s", s)
		return vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED
	}
	return vinv1.VehicleType(val)
}

func convertDataSource(s string) vinv1.DataSource {
	val, ok := vinv1.DataSource_value[s]
	if !ok {
		log.Printf("warning: unmappable data source value: %s", s)
		return vinv1.DataSource_DATA_SOURCE_UNSPECIFIED
	}
	return vinv1.DataSource(val)
}

func generateWMICode(outputFile, packageName string, standardIndex []StandardIndexEntry, blobPath string) error {
	blobName := filepath.Base(blobPath)
	f := codegen.NewFile(outputFile, "github.com/way-platform/vin-go")
	f.P("// Code generated by tools/cmd/wmi-gen-go. DO NOT EDIT.")
	f.P()
	f.P("package ", packageName)
	f.P()

	f.P(`import _ "embed"`)
	f.P()

	f.P("//go:embed ", blobName)
	f.P("var wmiBlob []byte")
	f.P()
	f.P("var wmiIndex = []struct { K uint16; O uint32 } {")
	for _, entry := range standardIndex {
		f.P(fmt.Sprintf("\t{K: 0x%x, O: 0x%x},", entry.K, entry.O))
	}
	f.P("}")

	content, err := f.Content()
	if err != nil {
		return fmt.Errorf("generating file content: %w", err)
	}

	if err := os.WriteFile(outputFile, content, 0o644); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

func generateLVMCode(outputFile, packageName string, lvmIndex []LVMIndexEntry, blobPath string) error {
	blobName := filepath.Base(blobPath)
	f := codegen.NewFile(outputFile, "github.com/way-platform/vin-go")
	f.P("// Code generated by tools/cmd/wmi-gen-go. DO NOT EDIT.")
	f.P()
	f.P("package ", packageName)
	f.P()

	f.P(`import _ "embed"`)
	f.P()

	f.P("//go:embed ", blobName)
	f.P("var lvmBlob []byte")
	f.P()
	f.P("var lvmIndex = []struct { K uint32; O uint32 } {")
	for _, entry := range lvmIndex {
		f.P(fmt.Sprintf("\t{K: 0x%x, O: 0x%x},", entry.K, entry.O))
	}
	f.P("}")

	content, err := f.Content()
	if err != nil {
		return fmt.Errorf("generating file content: %w", err)
	}

	if err := os.WriteFile(outputFile, content, 0o644); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}
