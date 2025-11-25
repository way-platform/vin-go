package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/way-platform/vin-go/internal/kba"
	"github.com/way-platform/vin-go/internal/vpic"
	"github.com/way-platform/vin-go/internal/wmi"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// Record represents a record in the WMI index.
type Record struct {
	WMI1         string
	WMI1Base36   uint16
	WMI2         string
	WMI2Base36   uint16
	DataSource   vinv1.DataSource
	Manufacturer string
	Country      vinv1.Country
	Region       vinv1.Region
	Brand        vinv1.Brand
	Category     vinv1.Category
}

func zeroToEmpty(s string) string {
	if s == "0" {
		return ""
	}
	return s
}

// CSV returns the record as a CSV record.
func (r *Record) CSV() []string {
	countryStr := ""
	if r.Country != vinv1.Country_COUNTRY_UNSPECIFIED {
		countryStr = r.Country.String()
	}
	regionStr := ""
	if r.Region != vinv1.Region_REGION_UNSPECIFIED {
		regionStr = r.Region.String()
	}
	brandStr := ""
	if r.Brand != vinv1.Brand_BRAND_UNSPECIFIED {
		brandStr = r.Brand.String()
	}
	categoryStr := ""
	if r.Category != vinv1.Category_CATEGORY_UNSPECIFIED {
		categoryStr = r.Category.String()
	}
	return []string{
		r.WMI1,
		strconv.FormatUint(uint64(r.WMI1Base36), 10),
		r.WMI2,
		zeroToEmpty(strconv.FormatUint(uint64(r.WMI2Base36), 10)),
		r.DataSource.String(),
		r.Manufacturer,
		countryStr,
		regionStr,
		brandStr,
		categoryStr,
	}
}

type VPICRecord struct {
	ID               int32
	WMI              string
	ManufacturerID   int32
	ManufacturerName string
	MakeID           int32
	MakeName         string
	CountryID        int32
	CountryName      string
}

func (r *VPICRecord) UnmarshalCSV(record []string) error {
	if len(record) != 8 {
		return fmt.Errorf("expected 8 columns, got %d", len(record))
	}
	var err error
	id, err := strconv.ParseInt(record[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	r.ID = int32(id)
	r.WMI = record[1]
	if len(r.WMI) != 3 && len(r.WMI) != 6 {
		return fmt.Errorf("WMI must be 3 or 6 characters, got %d: %s", len(r.WMI), r.WMI)
	}
	manufacturerID, err := strconv.ParseInt(record[2], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid ManufacturerID: %w", err)
	}
	r.ManufacturerID = int32(manufacturerID)
	r.ManufacturerName = record[3]
	if record[4] == "NULL" || record[4] == "" {
		r.MakeID = 0
	} else {
		makeID, err := strconv.ParseInt(record[4], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid MakeID: %w", err)
		}
		r.MakeID = int32(makeID)
	}
	r.MakeName = record[5]
	if record[6] == "NULL" || record[6] == "" {
		r.CountryID = 0
	} else {
		countryID, err := strconv.ParseInt(record[6], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid CountryID: %w", err)
		}
		r.CountryID = int32(countryID)
	}
	r.CountryName = record[7]
	return nil
}

type KBARecord struct {
	KBAManufacturerID    string
	WMI1                 string
	WMI2                 string
	ManufacturerName     string
	ManufacturerFullName string
	Location             string
}

func (r *KBARecord) UnmarshalCSV(record []string) error {
	if len(record) != 6 {
		return fmt.Errorf("expected 6 columns, got %d", len(record))
	}
	r.KBAManufacturerID = record[0]
	r.WMI1 = record[1]
	// Allow empty WMI1 (will be skipped in processing)
	if r.WMI1 != "" && len(r.WMI1) != 3 {
		return fmt.Errorf("WMI1 must be exactly 3 characters, got %d: %s", len(r.WMI1), r.WMI1)
	}
	r.WMI2 = record[2]
	if r.WMI2 != "" && len(r.WMI2) != 3 {
		return fmt.Errorf("WMI2 must be empty or exactly 3 characters, got %d: %s", len(r.WMI2), r.WMI2)
	}
	r.ManufacturerName = record[3]
	r.ManufacturerFullName = record[4]
	r.Location = record[5]
	return nil
}

type WikipediaRecord struct {
	WMI1         string
	Manufacturer string
	URL          string
}

func (r *WikipediaRecord) UnmarshalCSV(record []string) error {
	if len(record) != 3 {
		return fmt.Errorf("expected 3 columns, got %d", len(record))
	}
	r.WMI1 = record[0]
	r.Manufacturer = record[1]
	r.URL = record[2]
	return nil
}

type WikibooksRecord struct {
	WMI1         string
	WMI2         string
	Manufacturer string
}

func (r *WikibooksRecord) UnmarshalCSV(record []string) error {
	if len(record) != 3 {
		return fmt.Errorf("expected 3 columns, got %d", len(record))
	}
	r.WMI1 = record[0]
	// Allow non-3-character WMIs (will be skipped in processing, not an error)
	r.WMI2 = record[1]
	if r.WMI2 != "" && len(r.WMI2) != 3 {
		return fmt.Errorf("WMI2 must be empty or exactly 3 characters, got %d: %s", len(r.WMI2), r.WMI2)
	}
	r.Manufacturer = record[2]
	return nil
}

func main() {
	vpicFile := flag.String("vpic", "data/vpic/wmi.csv", "Path to the vPIC WMI CSV file")
	kbaFile := flag.String("kba", "data/kba/wmi.csv", "Path to the KBA WMI CSV file")
	wikipediaFile := flag.String("wikipedia", "data/wikipedia/wmi.csv", "Path to the Wikipedia WMI CSV file")
	wikibooksFile := flag.String("wikibooks", "data/wikibooks/wmi.csv", "Path to the Wikibooks WMI CSV file")
	vpicWMIMakeFile := flag.String("vpic-wmi-make", "data/vpic/wmi-make.csv", "Path to the vPIC WMI-Make lookup CSV file")
	outputFile := flag.String("o", "data/wmi.csv", "Path to the output CSV file")
	flag.Parse()
	if err := run(context.Background(), *vpicFile, *kbaFile, *wikipediaFile, *wikibooksFile, *vpicWMIMakeFile, *outputFile); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func run(ctx context.Context, vpicFile, kbaFile, wikipediaFile, wikibooksFile, vpicWMIMakeFile, outputFile string) error {
	// Map to store records keyed by WMI1+WMI2 (WMI2 empty for non-LVMs)
	index := make(map[string]*Record)

	// Load WMI-Make mapping
	wmiMakeMapping, err := loadWMIMakeMapping(vpicWMIMakeFile)
	if err != nil {
		return fmt.Errorf("loading WMI-Make mapping: %w", err)
	}

	// Step 1: Process VPIC (base layer)
	if err := processVPIC(ctx, vpicFile, wmiMakeMapping, index); err != nil {
		return fmt.Errorf("processing VPIC: %w", err)
	}

	// Step 2: Process KBA (add missing entries)
	if err := processKBA(ctx, kbaFile, index); err != nil {
		return fmt.Errorf("processing KBA: %w", err)
	}

	// Step 3: Process Wikipedia (add missing entries)
	if err := processWikipedia(ctx, wikipediaFile, index); err != nil {
		return fmt.Errorf("processing Wikipedia: %w", err)
	}

	// Step 4: Process Wikibooks (add missing entries)
	if err := processWikibooks(ctx, wikibooksFile, index); err != nil {
		return fmt.Errorf("processing Wikibooks: %w", err)
	}

	// Step 5: Write output CSV
	if err := writeOutput(outputFile, index); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}

// loadWMIMakeMapping loads the WMI-Make lookup table from CSV.
// Returns a map from WmiId to a slice of MakeIds (one WmiId can map to multiple MakeIds).
func loadWMIMakeMapping(filename string) (map[int32][]int32, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("empty WMI-Make file")
	}

	mapping := make(map[int32][]int32)

	// Skip header
	for i := 1; i < len(records); i++ {
		if len(records[i]) != 2 {
			return nil, fmt.Errorf("line %d: expected 2 columns, got %d", i+1, len(records[i]))
		}

		wmiId, err := strconv.ParseInt(records[i][0], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("line %d: invalid WmiId: %w", i+1, err)
		}

		makeId, err := strconv.ParseInt(records[i][1], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("line %d: invalid MakeId: %w", i+1, err)
		}

		wmiId32 := int32(wmiId)
		makeId32 := int32(makeId)

		// Append MakeId to the slice for this WmiId
		mapping[wmiId32] = append(mapping[wmiId32], makeId32)
	}

	return mapping, nil
}

func processVPIC(ctx context.Context, filename string, wmiMakeMapping map[int32][]int32, index map[string]*Record) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("empty VPIC file")
	}

	// Skip header
	for i := 1; i < len(records); i++ {
		var vpicRecord VPICRecord
		if err := vpicRecord.UnmarshalCSV(records[i]); err != nil {
			return fmt.Errorf("line %d: %w", i+1, err)
		}

		var wmi1, wmi2 string
		if len(vpicRecord.WMI) == 3 {
			wmi1 = vpicRecord.WMI
			wmi2 = ""
		} else if len(vpicRecord.WMI) == 6 {
			wmi1 = vpicRecord.WMI[0:3]
			wmi2 = vpicRecord.WMI[3:6]
		} else {
			return fmt.Errorf("line %d: invalid WMI length: %s", i+1, vpicRecord.WMI)
		}

		// Validate WMI1 is 3 characters
		if len(wmi1) != 3 {
			return fmt.Errorf("line %d: WMI1 must be 3 characters, got %d: %s", i+1, len(wmi1), wmi1)
		}

		// Validate WMI2 if present
		if wmi2 != "" && len(wmi2) != 3 {
			return fmt.Errorf("line %d: WMI2 must be 3 characters, got %d: %s", i+1, len(wmi2), wmi2)
		}

		// Convert to base36
		wmi1Base36, ok := wmi.ToBase36(wmi1)
		if !ok {
			return fmt.Errorf("line %d: invalid WMI1 for base36 conversion: %s", i+1, wmi1)
		}

		var wmi2Base36 uint16
		if wmi2 != "" {
			var ok bool
			wmi2Base36, ok = wmi.ToBase36(wmi2)
			if !ok {
				return fmt.Errorf("line %d: invalid WMI2 for base36 conversion: %s", i+1, wmi2)
			}
		}

		// Map MakeID to Brand (try existing MakeID first)
		brand, _ := vpic.ResolveBrand(vpicRecord.MakeID)

		// If brand is still unspecified, try WMI-Make lookup table
		if brand == vinv1.Brand_BRAND_UNSPECIFIED {
			makeIds, exists := wmiMakeMapping[vpicRecord.ID]
			if exists && len(makeIds) == 1 {
				// Only resolve if exactly one MakeID is mapped to this WmiId
				brand, _ = vpic.ResolveBrand(makeIds[0])
			}
		}

		// Map CountryID to Country
		country, _ := vpic.ResolveCountry(vpicRecord.CountryID)
		// Resolve region from WMI1
		region, _ := wmi.ResolveRegion(wmi1)

		// Create record
		record := &Record{
			WMI1:         wmi1,
			WMI1Base36:   wmi1Base36,
			WMI2:         wmi2,
			WMI2Base36:   wmi2Base36,
			DataSource:   vinv1.DataSource_VPIC,
			Manufacturer: vpicRecord.ManufacturerName,
			Country:      country,
			Region:       region,
			Brand:        brand,
		}

		// Store in index (keyed by WMI1+WMI2, WMI2 empty for non-LVMs)
		key := wmi1 + wmi2
		index[key] = record
	}

	return nil
}

func processKBA(ctx context.Context, filename string, index map[string]*Record) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("empty KBA file")
	}

	// Skip header
	for i := 1; i < len(records); i++ {
		var kbaRecord KBARecord
		if err := kbaRecord.UnmarshalCSV(records[i]); err != nil {
			return fmt.Errorf("line %d: %w", i+1, err)
		}

		// Skip records without WMI1 (KBA has its own manufacturer identification)
		if kbaRecord.WMI1 == "" {
			continue
		}

		// Skip if WMI1+WMI2 already exists
		key := kbaRecord.WMI1 + kbaRecord.WMI2
		if _, exists := index[key]; exists {
			continue
		}

		// Resolve country from WMI1
		country, _ := wmi.ResolveCountry(kbaRecord.WMI1)

		// Resolve region from WMI1
		region, _ := wmi.ResolveRegion(kbaRecord.WMI1)

		// Convert to base36
		wmi1Base36, ok := wmi.ToBase36(kbaRecord.WMI1)
		if !ok {
			return fmt.Errorf("line %d: invalid WMI1 for base36 conversion: %s", i+1, kbaRecord.WMI1)
		}

		manufacturerID, err := strconv.ParseInt(kbaRecord.KBAManufacturerID, 10, 32)
		if err != nil {
			return fmt.Errorf("line %d: invalid KBA manufacturer ID: %s", i+1, kbaRecord.KBAManufacturerID)
		}
		brand, _ := kba.ResolveBrand(int32(manufacturerID))

		var category vinv1.Category
	SpecLoop:
		for _, spec := range []struct {
			Category   vinv1.Category
			Substrings []string
		}{
			{
				Category: vinv1.Category_PASSENGER_CAR,
				Substrings: []string{
					"(PKW-MINI)",
					"(PERSONENWAGEN)",
					"(PASSENGER CARS)",
					"/PASSENGER CAR)",
				},
			},

			{
				Category: vinv1.Category_MULTIPURPOSE_PASSENGER_VEHICLE,
				Substrings: []string{
					"/MP-PASSENGER VEHICLE)",
					"/MULTIPURPOSE PASSENGER VEH)",
					"(MULTI PURPOSE VEHICLES)",
					"(MULTIPURPOSE.VEHICLE)",
				},
			},

			{
				Category: vinv1.Category_SPECIAL_PURPOSE_VEHICLE,
				Substrings: []string{
					"(SPECIAL PURPOSE VEHICLE)",
				},
			},

			{
				Category: vinv1.Category_MOTORCYCLE,
				Substrings: []string{
					"(KRAFTRAEDER)",
					"(MOTORCYCLE)",
				},
			},
			{
				Category: vinv1.Category_TRAILER,
				Substrings: []string{
					"(ANHAENGER)",
					"(COMMERCIAL TRAILER)",
					"(TRAILER)",
				},
			},
		} {
			for _, substring := range spec.Substrings {
				if strings.Contains(kbaRecord.ManufacturerFullName, substring) {
					category = spec.Category
					break SpecLoop
				}
			}
		}

		var wmi2Base36 uint16
		if kbaRecord.WMI2 != "" {
			var ok bool
			wmi2Base36, ok = wmi.ToBase36(kbaRecord.WMI2)
			if !ok {
				return fmt.Errorf("line %d: invalid WMI2 for base36 conversion: %s", i+1, kbaRecord.WMI2)
			}
		}

		// Create record
		record := &Record{
			WMI1:         kbaRecord.WMI1,
			WMI1Base36:   wmi1Base36,
			WMI2:         kbaRecord.WMI2,
			WMI2Base36:   wmi2Base36,
			DataSource:   vinv1.DataSource_KBA,
			Manufacturer: kbaRecord.ManufacturerName,
			Country:      country,
			Region:       region,
			Brand:        brand,
			Category:     category,
		}

		// Store in index
		index[key] = record
	}

	return nil
}

func processWikipedia(ctx context.Context, filename string, index map[string]*Record) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("empty Wikipedia file")
	}

	// Skip header
	for i := 1; i < len(records); i++ {
		var wikipediaRecord WikipediaRecord
		if err := wikipediaRecord.UnmarshalCSV(records[i]); err != nil {
			return fmt.Errorf("line %d: %w", i+1, err)
		}

		// Skip records without valid WMI1 (some entries are ranges or formatted text)
		if wikipediaRecord.WMI1 == "" || len(wikipediaRecord.WMI1) != 3 {
			continue
		}

		// Validate WMI1 can be converted to base36 (skip invalid WMIs like "2Gx")
		wmi1Base36, ok := wmi.ToBase36(wikipediaRecord.WMI1)
		if !ok {
			// Skip invalid WMIs silently (similar to Wikibooks handling)
			continue
		}

		// Skip if WMI1+WMI2 already exists (Wikipedia only has WMI1, so WMI2 is empty)
		key := wikipediaRecord.WMI1 + ""
		if _, exists := index[key]; exists {
			continue
		}

		// Resolve country from WMI1
		country, _ := wmi.ResolveCountry(wikipediaRecord.WMI1)

		// Resolve region from WMI1
		region, _ := wmi.ResolveRegion(wikipediaRecord.WMI1)

		// Create record
		record := &Record{
			WMI1:         wikipediaRecord.WMI1,
			WMI1Base36:   wmi1Base36,
			WMI2:         "",
			WMI2Base36:   0,
			DataSource:   vinv1.DataSource_WIKIPEDIA,
			Manufacturer: wikipediaRecord.Manufacturer,
			Country:      country,
			Region:       region,
			Brand:        vinv1.Brand_BRAND_UNSPECIFIED,
		}

		// Store in index
		index[key] = record
	}

	return nil
}

func processWikibooks(ctx context.Context, filename string, index map[string]*Record) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("empty Wikibooks file")
	}

	// Skip header
	for i := 1; i < len(records); i++ {
		var wikibooksRecord WikibooksRecord
		if err := wikibooksRecord.UnmarshalCSV(records[i]); err != nil {
			return fmt.Errorf("line %d: %w", i+1, err)
		}

		// Skip records without valid WMI1 (some entries are ranges or formatted text)
		if wikibooksRecord.WMI1 == "" || len(wikibooksRecord.WMI1) != 3 {
			continue
		}

		// Skip if WMI1+WMI2 already exists
		key := wikibooksRecord.WMI1 + wikibooksRecord.WMI2
		if _, exists := index[key]; exists {
			continue
		}

		// Resolve country from WMI1
		country, _ := wmi.ResolveCountry(wikibooksRecord.WMI1)

		// Resolve region from WMI1
		region, _ := wmi.ResolveRegion(wikibooksRecord.WMI1)

		// Convert to base36
		wmi1Base36, ok := wmi.ToBase36(wikibooksRecord.WMI1)
		if !ok {
			return fmt.Errorf("line %d: invalid WMI1 for base36 conversion: %s", i+1, wikibooksRecord.WMI1)
		}

		var wmi2Base36 uint16
		if wikibooksRecord.WMI2 != "" {
			var ok bool
			wmi2Base36, ok = wmi.ToBase36(wikibooksRecord.WMI2)
			if !ok {
				return fmt.Errorf("line %d: invalid WMI2 for base36 conversion: %s", i+1, wikibooksRecord.WMI2)
			}
		}

		// Create record
		record := &Record{
			WMI1:         wikibooksRecord.WMI1,
			WMI1Base36:   wmi1Base36,
			WMI2:         wikibooksRecord.WMI2,
			WMI2Base36:   wmi2Base36,
			DataSource:   vinv1.DataSource_WIKIBOOKS,
			Manufacturer: wikibooksRecord.Manufacturer,
			Country:      country,
			Region:       region,
			Brand:        vinv1.Brand_BRAND_UNSPECIFIED,
		}

		// Store in index
		index[key] = record
	}

	return nil
}

func writeOutput(filename string, index map[string]*Record) error {
	// Convert map to slice and sort
	records := make([]*Record, 0, len(index))
	for _, record := range index {
		records = append(records, record)
	}

	sort.Slice(records, func(i, j int) bool {
		if records[i].WMI1 != records[j].WMI1 {
			return records[i].WMI1 < records[j].WMI1
		}
		return records[i].WMI2 < records[j].WMI2
	})

	var writer *csv.Writer
	if filename == "-" {
		writer = csv.NewWriter(os.Stdout)
	} else {
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		writer = csv.NewWriter(f)
	}
	defer writer.Flush()

	// Write header
	header := []string{"WMI1", "WMI1Base36", "WMI2", "WMI2Base36", "DataSource", "Manufacturer", "Country", "Region", "Brand", "Category"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write records
	for _, record := range records {
		if err := writer.Write(record.CSV()); err != nil {
			return err
		}
	}

	return nil
}
