package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/way-platform/vin-go/internal/vpic"
	"github.com/way-platform/vin-go/internal/wmi"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	countryFile := flag.String("country", "", "vPIC country CSV file")
	makeModelFile := flag.String("make-model", "", "vPIC make-model CSV file")
	makeFile := flag.String("make", "", "vPIC make CSV file")
	manufacturerMakeFile := flag.String("manufacturer-make", "", "vPIC manufacturer-make CSV file")
	manufacturerFile := flag.String("manufacturer", "", "vPIC manufacturer CSV file")
	modelFile := flag.String("model", "", "vPIC model CSV file")
	wmiMakeFile := flag.String("wmi-make", "", "vPIC wmi-make CSV file")
	wmiFile := flag.String("wmi", "", "vPIC wmi CSV file")
	outputFile := flag.String("o", "-", "Output JSONL file path (defaults to stdout)")
	flag.Parse()
	if err := run(&InputFiles{
		CountryFile:          *countryFile,
		MakeModelFile:        *makeModelFile,
		MakeFile:             *makeFile,
		ManufacturerMakeFile: *manufacturerMakeFile,
		ManufacturerFile:     *manufacturerFile,
		ModelFile:            *modelFile,
		WmiMakeFile:          *wmiMakeFile,
		WmiFile:              *wmiFile,
	}, *outputFile); err != nil {
		log.Fatal(err)
	}
}

type InputFiles struct {
	CountryFile          string
	MakeModelFile        string
	MakeFile             string
	ManufacturerMakeFile string
	ManufacturerFile     string
	ModelFile            string
	WmiMakeFile          string
	WmiFile              string
}

type CountryRecord struct {
	Id       uint16
	Name     string
	WmiCount uint16
}

func (r *CountryRecord) UnmarshalCSV(record []string) error {
	const expectedLength = 3
	if len(record) != expectedLength {
		return fmt.Errorf("unexpected country record length %d expected %d", len(record), expectedLength)
	}
	id, err := strconv.ParseUint(record[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	r.Id = uint16(id)
	r.Name = strings.TrimSpace(record[1])
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	wmiCount, err := strconv.ParseUint(record[2], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid WMI count: %w", err)
	}
	r.WmiCount = uint16(wmiCount)
	return nil
}

type MakeModelRecord struct {
	Id        uint16
	MakeId    uint16
	ModelId   uint16
	CreatedOn time.Time
	UpdatedOn time.Time
}

func (r *MakeModelRecord) UnmarshalCSV(record []string) error {
	const expectedLength = 5
	if len(record) != expectedLength {
		return fmt.Errorf("unexpected make-model record length %d expected %d", len(record), expectedLength)
	}
	id, err := strconv.ParseUint(record[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	r.Id = uint16(id)
	makeId, err := strconv.ParseUint(record[1], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid MakeId: %w", err)
	}
	r.MakeId = uint16(makeId)
	modelId, err := strconv.ParseUint(record[2], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ModelId: %w", err)
	}
	r.ModelId = uint16(modelId)
	r.CreatedOn, err = parseNullableTime(record[3])
	if err != nil {
		return fmt.Errorf("invalid CreatedOn: %w", err)
	}
	r.UpdatedOn, err = parseNullableTime(record[4])
	if err != nil {
		return fmt.Errorf("invalid UpdatedOn: %w", err)
	}
	return nil
}

type MakeRecord struct {
	Id        uint16
	Name      string
	CreatedOn time.Time
	UpdatedOn time.Time
}

func (r *MakeRecord) UnmarshalCSV(record []string) error {
	const expectedLength = 4
	if len(record) != expectedLength {
		return fmt.Errorf("unexpected make record length %d expected %d", len(record), expectedLength)
	}
	id, err := strconv.ParseUint(record[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	r.Id = uint16(id)
	r.Name = strings.TrimSpace(record[1])
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	r.CreatedOn, err = parseNullableTime(record[2])
	if err != nil {
		return fmt.Errorf("invalid CreatedOn: %w", err)
	}
	r.UpdatedOn, err = parseNullableTime(record[3])
	if err != nil {
		return fmt.Errorf("invalid UpdatedOn: %w", err)
	}
	return nil
}

type ManufacturerMakeRecord struct {
	Id             uint16
	ManufacturerId uint16
	MakeId         uint16
}

func (r *ManufacturerMakeRecord) UnmarshalCSV(record []string) error {
	const expectedLength = 3
	if len(record) != expectedLength {
		return fmt.Errorf("unexpected manufacturer-make record length %d expected %d", len(record), expectedLength)
	}
	id, err := strconv.ParseUint(record[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	r.Id = uint16(id)
	manufacturerId, err := strconv.ParseUint(record[1], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ManufacturerId: %w", err)
	}
	r.ManufacturerId = uint16(manufacturerId)
	makeId, err := strconv.ParseUint(record[2], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid MakeId: %w", err)
	}
	r.MakeId = uint16(makeId)
	return nil
}

type ManufacturerRecord struct {
	Id   uint16
	Name string
}

func (r *ManufacturerRecord) UnmarshalCSV(record []string) error {
	const expectedLength = 2
	if len(record) != expectedLength {
		return fmt.Errorf("unexpected manufacturer record length %d expected %d", len(record), expectedLength)
	}
	id, err := strconv.ParseUint(record[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	r.Id = uint16(id)
	r.Name = strings.TrimSpace(record[1])
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

type ModelRecord struct {
	Id        uint16
	Name      string
	CreatedOn time.Time
	UpdatedOn time.Time
}

func (r *ModelRecord) UnmarshalCSV(record []string) error {
	const expectedLength = 4
	if len(record) != expectedLength {
		return fmt.Errorf("unexpected model record length %d expected %d", len(record), expectedLength)
	}
	id, err := strconv.ParseUint(record[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	r.Id = uint16(id)
	r.Name = strings.TrimSpace(record[1])
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	r.CreatedOn, err = parseNullableTime(record[2])
	if err != nil {
		return fmt.Errorf("invalid CreatedOn: %w", err)
	}
	r.UpdatedOn, err = parseNullableTime(record[3])
	if err != nil {
		return fmt.Errorf("invalid UpdatedOn: %w", err)
	}
	return nil
}

type WmiMakeRecord struct {
	WmiId  uint16
	MakeId uint16
}

func (r *WmiMakeRecord) UnmarshalCSV(record []string) error {
	const expectedLength = 2
	if len(record) != expectedLength {
		return fmt.Errorf("unexpected wmi-make record length %d expected %d", len(record), expectedLength)
	}
	wmiId, err := strconv.ParseUint(record[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid WmiId: %w", err)
	}
	r.WmiId = uint16(wmiId)
	makeId, err := strconv.ParseUint(record[1], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid MakeId: %w", err)
	}
	r.MakeId = uint16(makeId)
	return nil
}

type WmiRecord struct {
	Id                     uint16
	Wmi                    string
	ManufacturerId         uint16
	MakeId                 uint16
	VehicleTypeId          uint16
	CreatedOn              time.Time
	UpdatedOn              time.Time
	CountryId              uint16
	PublicAvailabilityDate time.Time
	TruckTypeId            uint16
	ProcessedOn            time.Time
	NonCompliant           bool
	NonCompliantSetByOVSC  bool
}

func (r *WmiRecord) UnmarshalCSV(record []string) error {
	const expectedLength = 13
	if len(record) != expectedLength {
		return fmt.Errorf("unexpected wmi record length %d expected %d", len(record), expectedLength)
	}
	id, err := strconv.ParseUint(record[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}
	r.Id = uint16(id)
	r.Wmi = strings.TrimSpace(record[1])
	if r.Wmi == "" {
		return fmt.Errorf("wmi is required")
	}
	manufacturerId, err := parseNullableUint16(record[2])
	if err != nil {
		return fmt.Errorf("invalid ManufacturerId: %w", err)
	}
	r.ManufacturerId = manufacturerId
	makeId, err := parseNullableUint16(record[3])
	if err != nil {
		return fmt.Errorf("invalid MakeId: %w", err)
	}
	r.MakeId = makeId
	vehicleTypeId, err := parseNullableUint16(record[4])
	if err != nil {
		return fmt.Errorf("invalid VehicleTypeId: %w", err)
	}
	r.VehicleTypeId = vehicleTypeId
	r.CreatedOn, err = parseNullableTime(record[5])
	if err != nil {
		return fmt.Errorf("invalid CreatedOn: %w", err)
	}
	r.UpdatedOn, err = parseNullableTime(record[6])
	if err != nil {
		return fmt.Errorf("invalid UpdatedOn: %w", err)
	}
	countryId, err := parseNullableUint16(record[7])
	if err != nil {
		return fmt.Errorf("invalid CountryId: %w", err)
	}
	r.CountryId = countryId
	r.PublicAvailabilityDate, err = parseNullableTime(record[8])
	if err != nil {
		return fmt.Errorf("invalid PublicAvailabilityDate: %w", err)
	}
	truckTypeId, err := parseNullableUint16(record[9])
	if err != nil {
		return fmt.Errorf("invalid TruckTypeId: %w", err)
	}
	r.TruckTypeId = truckTypeId
	r.ProcessedOn, err = parseNullableTime(record[10])
	if err != nil {
		return fmt.Errorf("invalid ProcessedOn: %w", err)
	}
	r.NonCompliant, err = parseNullableBool(record[11])
	if err != nil {
		return fmt.Errorf("invalid NonCompliant: %w", err)
	}
	r.NonCompliantSetByOVSC, err = parseNullableBool(record[12])
	if err != nil {
		return fmt.Errorf("invalid NonCompliantSetByOVSC: %w", err)
	}
	return nil
}

type DB struct {
	Country          []CountryRecord
	MakeModel        []MakeModelRecord
	Make             []MakeRecord
	ManufacturerMake []ManufacturerMakeRecord
	Manufacturer     []ManufacturerRecord
	Model            []ModelRecord
	WmiMake          []WmiMakeRecord
	Wmi              []WmiRecord
}

type Indices struct {
	Manufacturers     map[uint16]*ManufacturerRecord
	ManufacturerMakes map[uint16]map[uint16]bool
	WmiMakes          map[uint16]map[uint16]bool
	MakeModels        map[uint16]map[uint16]bool
}

func run(inputFiles *InputFiles, outputFile string) error {
	db, err := loadDB(inputFiles)
	if err != nil {
		return err
	}

	indices, err := buildIndices(db)
	if err != nil {
		return fmt.Errorf("building indices: %w", err)
	}

	// Deduplicate by WMI key (wmi1+wmi2)
	manufacturersByWmi := make(map[string]*vinv1.Manufacturer)

	for _, wmiRecord := range db.Wmi {
		manufacturer, err := transformWmiToManufacturer(wmiRecord, indices)
		if err != nil {
			log.Printf("warning: skipping WMI record %d: %v", wmiRecord.Id, err)
			continue
		}
		if manufacturer == nil {
			continue
		}

		// Create deduplication key
		wmiKey := manufacturer.GetWmi1() + manufacturer.GetWmi2()

		// Merge with existing if found
		if existing, exists := manufacturersByWmi[wmiKey]; exists {
			mergeManufacturers(existing, manufacturer)
		} else {
			manufacturersByWmi[wmiKey] = manufacturer
		}
	}

	// Output JSONL
	var out *os.File
	if outputFile == "-" {
		out = os.Stdout
	} else {
		out, err = os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer out.Close()
	}
	marshaler := protojson.MarshalOptions{}
	for _, manufacturer := range manufacturersByWmi {
		data, err := marshaler.Marshal(manufacturer)
		if err != nil {
			return fmt.Errorf("failed to marshal manufacturer: %w", err)
		}
		var buf bytes.Buffer
		if err := json.Compact(&buf, data); err != nil {
			return fmt.Errorf("failed to compact JSON: %w", err)
		}
		fmt.Fprintln(out, buf.String())
	}
	return nil
}

func loadDB(inputFiles *InputFiles) (*DB, error) {
	var output DB
	var err error
	output.Country, err = loadRecords(inputFiles.CountryFile, (*CountryRecord).UnmarshalCSV)
	if err != nil {
		return nil, fmt.Errorf("loading country: %w", err)
	}
	output.MakeModel, err = loadRecords(inputFiles.MakeModelFile, (*MakeModelRecord).UnmarshalCSV)
	if err != nil {
		return nil, fmt.Errorf("loading make-model: %w", err)
	}
	output.Make, err = loadRecords(inputFiles.MakeFile, (*MakeRecord).UnmarshalCSV)
	if err != nil {
		return nil, fmt.Errorf("loading make: %w", err)
	}
	output.ManufacturerMake, err = loadRecords(inputFiles.ManufacturerMakeFile, (*ManufacturerMakeRecord).UnmarshalCSV)
	if err != nil {
		return nil, fmt.Errorf("loading manufacturer-make: %w", err)
	}
	output.Manufacturer, err = loadRecords(inputFiles.ManufacturerFile, (*ManufacturerRecord).UnmarshalCSV)
	if err != nil {
		return nil, fmt.Errorf("loading manufacturer: %w", err)
	}
	output.Model, err = loadRecords(inputFiles.ModelFile, (*ModelRecord).UnmarshalCSV)
	if err != nil {
		return nil, fmt.Errorf("loading model: %w", err)
	}
	output.WmiMake, err = loadRecords(inputFiles.WmiMakeFile, (*WmiMakeRecord).UnmarshalCSV)
	if err != nil {
		return nil, fmt.Errorf("loading wmi-make: %w", err)
	}
	output.Wmi, err = loadRecords(inputFiles.WmiFile, (*WmiRecord).UnmarshalCSV)
	if err != nil {
		return nil, fmt.Errorf("loading wmi: %w", err)
	}
	return &output, nil
}

func loadRecords[R any](inputFile string, fn func(*R, []string) error) ([]R, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	all, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	output := make([]R, 0, len(all))
	for i, record := range all {
		if i == 0 {
			continue
		}
		var r R
		if err := fn(&r, record); err != nil {
			return nil, err
		}
		output = append(output, r)
	}
	return output, nil
}

type Record interface {
	UnmarshalCSV([]string) error
}

// parseNullableTime parses a timestamp string or returns zero time for NULL/empty values
func parseNullableTime(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" || s == "NULL" {
		return time.Time{}, nil
	}
	// Try multiple date formats
	formats := []string{
		"2006-01-02 15:04:05.000",
		"2006-01-02 15:04:05",
		"2006-01-02",
		time.RFC3339,
	}
	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse time: %q", s)
}

// parseNullableUint16 parses a uint16 or returns 0 for NULL/empty values
func parseNullableUint16(s string) (uint16, error) {
	s = strings.TrimSpace(s)
	if s == "" || s == "NULL" {
		return 0, nil
	}
	val, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(val), nil
}

// parseNullableBool parses a boolean or returns false for NULL/empty values
func parseNullableBool(s string) (bool, error) {
	s = strings.TrimSpace(s)
	if s == "" || s == "NULL" {
		return false, nil
	}
	val, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return val, nil
}

// buildIndices creates efficient lookup maps from the loaded database
func buildIndices(db *DB) (*Indices, error) {
	indices := &Indices{
		Manufacturers:     make(map[uint16]*ManufacturerRecord),
		ManufacturerMakes: make(map[uint16]map[uint16]bool),
		WmiMakes:          make(map[uint16]map[uint16]bool),
		MakeModels:        make(map[uint16]map[uint16]bool),
	}

	// Build manufacturers index
	for i := range db.Manufacturer {
		indices.Manufacturers[db.Manufacturer[i].Id] = &db.Manufacturer[i]
	}

	// Build manufacturer-makes index
	for _, record := range db.ManufacturerMake {
		if indices.ManufacturerMakes[record.ManufacturerId] == nil {
			indices.ManufacturerMakes[record.ManufacturerId] = make(map[uint16]bool)
		}
		indices.ManufacturerMakes[record.ManufacturerId][record.MakeId] = true
	}

	// Build wmi-makes index
	for _, record := range db.WmiMake {
		if indices.WmiMakes[record.WmiId] == nil {
			indices.WmiMakes[record.WmiId] = make(map[uint16]bool)
		}
		indices.WmiMakes[record.WmiId][record.MakeId] = true
	}

	// Build make-models index
	for _, record := range db.MakeModel {
		if indices.MakeModels[record.MakeId] == nil {
			indices.MakeModels[record.MakeId] = make(map[uint16]bool)
		}
		indices.MakeModels[record.MakeId][record.ModelId] = true
	}

	return indices, nil
}

// parseWMI extracts wmi1, wmi2, and low_volume flag from a WMI string
func parseWMI(wmi string) (wmi1, wmi2 string, lowVolume bool, err error) {
	wmi = strings.TrimSpace(wmi)
	if len(wmi) == 3 {
		wmi1 = wmi
		wmi2 = ""
		lowVolume = false
		return
	}
	if len(wmi) == 6 {
		wmi1 = wmi[0:3]
		wmi2 = wmi[3:6]
		lowVolume = true
		return
	}
	return "", "", false, fmt.Errorf("invalid WMI length: expected 3 or 6 characters, got %d", len(wmi))
}

// transformWmiToManufacturer converts a WmiRecord to a Manufacturer proto message
func transformWmiToManufacturer(wmiRecord WmiRecord, indices *Indices) (*vinv1.Manufacturer, error) {
	// Parse WMI
	wmi1, wmi2, lowVolume, err := parseWMI(wmiRecord.Wmi)
	if err != nil {
		return nil, err
	}

	// Create manufacturer proto
	manufacturer := &vinv1.Manufacturer{}
	manufacturer.SetWmi1(wmi1)
	if wmi2 != "" {
		manufacturer.SetWmi2(wmi2)
	}
	if lowVolume {
		manufacturer.SetLowVolume(true)
	}

	// Set vpic_id
	if wmiRecord.ManufacturerId != 0 {
		manufacturer.SetVpicId(int32(wmiRecord.ManufacturerId))
	}

	// Set display_name from manufacturers index
	if wmiRecord.ManufacturerId != 0 {
		if mfr, exists := indices.Manufacturers[wmiRecord.ManufacturerId]; exists {
			manufacturer.SetDisplayName(mfr.Name)
		}
	}

	// Resolve country
	if wmiRecord.CountryId != 0 {
		if country, ok := vpic.ResolveCountry(int32(wmiRecord.CountryId)); ok {
			manufacturer.SetCountry(country)
		}
	}

	// Resolve region from WMI1
	if region, ok := wmi.ResolveRegion(wmi1); ok {
		manufacturer.SetRegion(region)
	}

	// Collect distinct MakeIDs
	makeIDs := make(map[uint16]bool)

	// 1. Direct MakeId from WMI record
	if wmiRecord.MakeId != 0 {
		makeIDs[wmiRecord.MakeId] = true
	}

	// 2. WMI-specific makes
	if wmiMakes, exists := indices.WmiMakes[wmiRecord.Id]; exists {
		for makeID := range wmiMakes {
			makeIDs[makeID] = true
		}
	}

	// 3. All manufacturer makes (broad approach)
	if wmiRecord.ManufacturerId != 0 {
		if mfrMakes, exists := indices.ManufacturerMakes[wmiRecord.ManufacturerId]; exists {
			for makeID := range mfrMakes {
				makeIDs[makeID] = true
			}
		}
	}

	// Resolve brands and collect models
	brands := make(map[vinv1.Brand]bool)
	models := make(map[vinv1.Model]bool)

	for makeID := range makeIDs {
		// Resolve brand
		if brand, ok := vpic.ResolveBrand(int32(makeID)); ok && brand != vinv1.Brand_BRAND_UNSPECIFIED {
			brands[brand] = true
		}

		// Collect models for this make
		if makeModelIDs, exists := indices.MakeModels[makeID]; exists {
			for modelID := range makeModelIDs {
				if model, ok := vpic.ResolveModel(int32(modelID)); ok && model != vinv1.Model_MODEL_UNSPECIFIED {
					models[model] = true
				}
			}
		}
	}

	// Convert brand set to slice
	brandSlice := make([]vinv1.Brand, 0, len(brands))
	for brand := range brands {
		brandSlice = append(brandSlice, brand)
	}
	manufacturer.SetBrands(brandSlice)

	// Convert model set to slice
	modelSlice := make([]vinv1.Model, 0, len(models))
	for model := range models {
		modelSlice = append(modelSlice, model)
	}
	manufacturer.SetModels(modelSlice)

	// Resolve vehicle types
	if wmiRecord.VehicleTypeId != 0 {
		if vehicleType, ok := vpic.ResolveVehicleType(int32(wmiRecord.VehicleTypeId)); ok && vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
			manufacturer.SetVehicleTypes([]vinv1.VehicleType{vehicleType})
		}
	}

	// Set data source
	manufacturer.SetDataSources([]vinv1.DataSource{vinv1.DataSource_VPIC})

	return manufacturer, nil
}

// mergeManufacturers merges fields from src into dst, deduplicating repeated fields
func mergeManufacturers(dst, src *vinv1.Manufacturer) {
	// Merge brands
	brandSet := make(map[vinv1.Brand]bool)
	for _, brand := range dst.GetBrands() {
		if brand != vinv1.Brand_BRAND_UNSPECIFIED {
			brandSet[brand] = true
		}
	}
	for _, brand := range src.GetBrands() {
		if brand != vinv1.Brand_BRAND_UNSPECIFIED {
			brandSet[brand] = true
		}
	}
	brandSlice := make([]vinv1.Brand, 0, len(brandSet))
	for brand := range brandSet {
		brandSlice = append(brandSlice, brand)
	}
	dst.SetBrands(brandSlice)

	// Merge vehicle types
	vehicleTypeSet := make(map[vinv1.VehicleType]bool)
	for _, vt := range dst.GetVehicleTypes() {
		if vt != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
			vehicleTypeSet[vt] = true
		}
	}
	for _, vt := range src.GetVehicleTypes() {
		if vt != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
			vehicleTypeSet[vt] = true
		}
	}
	vehicleTypeSlice := make([]vinv1.VehicleType, 0, len(vehicleTypeSet))
	for vt := range vehicleTypeSet {
		vehicleTypeSlice = append(vehicleTypeSlice, vt)
	}
	dst.SetVehicleTypes(vehicleTypeSlice)

	// Merge models
	modelSet := make(map[vinv1.Model]bool)
	for _, model := range dst.GetModels() {
		if model != vinv1.Model_MODEL_UNSPECIFIED {
			modelSet[model] = true
		}
	}
	for _, model := range src.GetModels() {
		if model != vinv1.Model_MODEL_UNSPECIFIED {
			modelSet[model] = true
		}
	}
	modelSlice := make([]vinv1.Model, 0, len(modelSet))
	for model := range modelSet {
		modelSlice = append(modelSlice, model)
	}
	dst.SetModels(modelSlice)

	// Keep first non-empty display_name
	if dst.GetDisplayName() == "" && src.GetDisplayName() != "" {
		dst.SetDisplayName(src.GetDisplayName())
	}

	// Merge data sources
	dataSourceSet := make(map[vinv1.DataSource]bool)
	for _, ds := range dst.GetDataSources() {
		if ds != vinv1.DataSource_DATA_SOURCE_UNSPECIFIED {
			dataSourceSet[ds] = true
		}
	}
	for _, ds := range src.GetDataSources() {
		if ds != vinv1.DataSource_DATA_SOURCE_UNSPECIFIED {
			dataSourceSet[ds] = true
		}
	}
	dataSourceSlice := make([]vinv1.DataSource, 0, len(dataSourceSet))
	for ds := range dataSourceSet {
		dataSourceSlice = append(dataSourceSlice, ds)
	}
	dst.SetDataSources(dataSourceSlice)
}
