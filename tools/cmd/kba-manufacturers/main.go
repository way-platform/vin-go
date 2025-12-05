package main

import (
	"bufio"
	"context"
	_ "embed"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	"google.golang.org/genai"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//go:embed system-prompt.txt
var baseSystemPrompt string

var bracketRegex = regexp.MustCompile(`\((.*?)\)`)

var vehicleTypeKeywords = []struct {
	Type     vinv1.VehicleType
	Keywords []string
}{
	{
		Type:     vinv1.VehicleType_TRAILER,
		Keywords: []string{"anhaenger", "anh.", "anhänger", "trailer", "caravan", "wohnwagen"},
	},
	{
		Type:     vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
		Keywords: []string{"lkw", "nutzfahrzeug", "nutzfahrzeuge", "zugmaschinen", "truck", "trucks"},
	},
	{
		Type:     vinv1.VehicleType_BUS,
		Keywords: []string{"bus", "omnibus", "kleinbus", "kleinbusse"},
	},
	{
		Type:     vinv1.VehicleType_MOTORCYCLE,
		Keywords: []string{"kraftrad", "kraftraeder", "krad", "kräder", "motorcycle", "motorcycles", "zweirad"},
	},
	{
		Type:     vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
		Keywords: []string{"van", "pick up", "pickup"},
	},
}

func normalizeForSearch(s string) string {
	var b strings.Builder
	// Pad with space to handle start of string
	b.WriteRune(' ')
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
		} else {
			b.WriteRune(' ')
		}
	}
	// Pad with space to handle end of string
	b.WriteRune(' ')
	return b.String()
}

func inferVehicleTypes(fullName string) []vinv1.VehicleType {
	matches := bracketRegex.FindAllStringSubmatch(fullName, -1)
	var types []vinv1.VehicleType
	seen := make(map[vinv1.VehicleType]bool)

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		content := match[1]
		normalized := normalizeForSearch(content)

		for _, matcher := range vehicleTypeKeywords {
			for _, keyword := range matcher.Keywords {
				if strings.Contains(normalized, " "+keyword+" ") {
					if !seen[matcher.Type] {
						types = append(types, matcher.Type)
						seen[matcher.Type] = true
					}
					break
				}
			}
		}
	}
	return types
}

func init() {
	baseSystemPrompt = strings.ReplaceAll(baseSystemPrompt, "{{COUNTRIES}}", getEnumValuesForPrompt(vinv1.Country_COUNTRY_UNSPECIFIED.Descriptor()))
	baseSystemPrompt = strings.ReplaceAll(baseSystemPrompt, "{{BRANDS}}", getEnumValuesForPrompt(vinv1.Brand_BRAND_UNSPECIFIED.Descriptor()))
}

func main() {
	inputFile := flag.String("i", "", "Input CSV file path")
	outputFile := flag.String("o", "-", "Output CSV file path (defaults to stdout)")
	flag.Parse()
	if err := run(context.Background(), *inputFile, *outputFile); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}

func run(ctx context.Context, inputFile, outputFile string) error {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		return fmt.Errorf("GOOGLE_CLOUD_PROJECT environment variable not set")
	}
	location := os.Getenv("GOOGLE_CLOUD_LOCATION")
	if location == "" {
		return fmt.Errorf("GOOGLE_CLOUD_LOCATION environment variable not set")
	}
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  projectID,
		Location: location,
		Backend:  genai.BackendVertexAI,
	})
	if err != nil {
		return fmt.Errorf("failed to create Vertex AI client: %w", err)
	}
	records, err := readInputFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}
	slog.Info("read input file", "inputFile", inputFile, "records", len(records))

	// Read existing output file if it exists
	existingByWMI, err := readExistingOutput(outputFile)
	if err != nil {
		return fmt.Errorf("failed to read existing output: %w", err)
	}

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
	// Validate records.
	for i, record := range records {
		slog.Info("validating record", "record", record, "index", i)
		if len(record) != 6 {
			return fmt.Errorf("invalid CSV record: %v", record)
		}
		kba, wmi1, wmi2, name := record[0], record[1], record[2], record[3]
		if i > 0 && len(wmi1) != 3 {
			return fmt.Errorf("invalid WMI1: %s", wmi1)
		}
		if i > 0 && len(wmi2) > 0 && len(wmi2) != 3 {
			return fmt.Errorf("invalid WMI2: %s", wmi2)
		}
		if name == "" {
			return fmt.Errorf("empty name")
		}
		if i > 0 {
			if _, err := strconv.ParseInt(kba, 10, 64); err != nil {
				return fmt.Errorf("invalid KBA: %w", err)
			}
		}
	}
	slog.Info("validated records", "records", len(records))
	header := records[0]
	records = records[1:]

	// Synchronization primitives
	var (
		existingMu   sync.Mutex // Protects existingByWMI
		outputMu     sync.Mutex // Protects writing to out
		cacheMu      sync.Mutex // Protects cache
		cache        = make(map[string]*vinv1.Manufacturer)
		requestGroup singleflight.Group
	)

	// Create errgroup with limit of 5 concurrent operations
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(5)

	// Process records concurrently
	for _, record := range records {
		g.Go(func() error {
			kba, wmi1, wmi2, name, fullName, location := record[0], record[1], record[2], record[3], record[4], record[5]

			// Normalize WMI values for matching
			wmi1Normalized := strings.ToUpper(strings.TrimSpace(wmi1))
			wmi2Normalized := strings.ToUpper(strings.TrimSpace(wmi2))
			wmiKey := makeWMIKey(wmi1Normalized, wmi2Normalized)

			var manufacturer *vinv1.Manufacturer

			// Check if we have an existing record for this WMI combination
			existingMu.Lock()
			existing, ok := existingByWMI[wmiKey]
			existingMu.Unlock()

			if ok {
				slog.Info("reusing existing manufacturer", "wmiKey", wmiKey, "wmi1", wmi1Normalized, "wmi2", wmi2Normalized)
				manufacturer = proto.Clone(existing).(*vinv1.Manufacturer)
			} else {
				// Check cache first (by name/fullName/location)
				cacheKey := makeCacheKey(name, fullName, location)
				cacheMu.Lock()
				cached, hit := cache[cacheKey]
				cacheMu.Unlock()

				if hit {
					slog.Info("using cached manufacturer", "cacheKey", cacheKey)
					manufacturer = proto.Clone(cached).(*vinv1.Manufacturer)
				} else {
					// Use singleflight to coalesce concurrent requests
					result, err, _ := requestGroup.Do(cacheKey, func() (interface{}, error) {
						// Double-check memory cache
						cacheMu.Lock()
						if cached, hit := cache[cacheKey]; hit {
							cacheMu.Unlock()
							return cached, nil
						}
						cacheMu.Unlock()

						// Perform AI inference
						startTime := time.Now()
						proto, err := generateManufacturer(gctx, client, header, record)
						slog.Info("generated manufacturer", "cacheKey", cacheKey, "wmiKey", wmiKey, "time", time.Since(startTime))
						if err != nil {
							return nil, err
						}

						// Store in cache
						cacheMu.Lock()
						cache[cacheKey] = proto
						cacheMu.Unlock()

						return proto, nil
					})

					if err != nil {
						return fmt.Errorf("failed to generate manufacturer: %w", err)
					}
					manufacturer = result.(*vinv1.Manufacturer)
				}
			}

			// Update manufacturer fields from input record
			manufacturer = proto.Clone(manufacturer).(*vinv1.Manufacturer)
			kbaID, err := strconv.ParseInt(kba, 10, 32)
			if err != nil {
				return fmt.Errorf("invalid KBA: %w", err)
			}
			manufacturer.SetKbaId(int32(kbaID))
			manufacturer.SetWmi1(wmi1Normalized)
			if wmi2Normalized != "" {
				manufacturer.SetWmi2(wmi2Normalized)
			}
			if fullName != "" {
				manufacturer.SetDisplayName(fullName)
			} else {
				manufacturer.SetDisplayName(name)
			}

			// Infer vehicle types from bracketed text in FullName
			if types := inferVehicleTypes(fullName); len(types) > 0 {
				manufacturer.SetVehicleTypes(types)
			}

			// Set data source
			manufacturer.SetDataSources([]vinv1.DataSource{vinv1.DataSource_KBA})

			// Marshal and write immediately
			data, err := protojson.MarshalOptions{}.Marshal(manufacturer)
			if err != nil {
				return fmt.Errorf("failed to marshal manufacturer: %w", err)
			}

			outputMu.Lock()
			_, err = fmt.Fprintln(out, string(data))
			if err == nil && outputFile != "-" {
				// Sync to ensure data is flushed to disk immediately
				err = out.Sync()
			}
			outputMu.Unlock()

			if err != nil {
				return fmt.Errorf("failed to write manufacturer: %w", err)
			}

			// Update existingByWMI map to track processed records
			existingMu.Lock()
			existingByWMI[wmiKey] = manufacturer
			existingMu.Unlock()

			return nil
		})
	}

	// Wait for all goroutines to complete
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func getEnumValuesForPrompt(descriptor protoreflect.EnumDescriptor) string {
	valuesDesc := descriptor.Values()
	values := make([]string, 0, valuesDesc.Len()-1)
	for i := 0; i < valuesDesc.Len(); i++ {
		value := valuesDesc.Get(i)
		if value.Number() == 0 {
			continue // Skip UNSPECIFIED
		}
		values = append(values, string(value.Name()))
	}
	return strings.Join(values, ",")
}

func makeCacheKey(name, fullName, location string) string {
	return fmt.Sprintf("%s,%s,%s", name, fullName, location)
}

func generateManufacturer(ctx context.Context, client *genai.Client, header []string, row []string) (*vinv1.Manufacturer, error) {
	var prompt strings.Builder
	promptWriter := csv.NewWriter(&prompt)
	if err := promptWriter.WriteAll([][]string{header, row}); err != nil {
		return nil, fmt.Errorf("failed to write prompt: %w", err)
	}
	promptWriter.Flush()
	if promptWriter.Error() != nil {
		return nil, fmt.Errorf("failed to flush prompt: %w", promptWriter.Error())
	}
	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		[]*genai.Content{
			{Role: genai.RoleUser, Parts: []*genai.Part{{Text: prompt.String()}}},
		},
		&genai.GenerateContentConfig{
			SystemInstruction: &genai.Content{Role: genai.RoleUser, Parts: []*genai.Part{{Text: baseSystemPrompt}}},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content from Vertex AI API: %w", err)
	}
	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no content received from Vertex AI API")
	}
	var responseText strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		responseText.WriteString(part.Text)
	}
	// Clean up response text - remove markdown code fences if present
	jsonText := strings.TrimSpace(responseText.String())
	if cut, ok := strings.CutPrefix(jsonText, "```json"); ok {
		cut = strings.TrimSuffix(cut, "```")
		jsonText = strings.TrimSpace(cut)
	} else if cut, ok := strings.CutPrefix(jsonText, "```"); ok {
		cut = strings.TrimSuffix(cut, "```")
		jsonText = strings.TrimSpace(cut)
	}
	var manufacturer vinv1.Manufacturer
	if err := protojson.Unmarshal([]byte(jsonText), &manufacturer); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}
	return &manufacturer, nil
}

func readInputFile(inputFile string) ([][]string, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file %s: %w", inputFile, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	return csvReader.ReadAll()
}

// makeWMIKey creates a composite key from wmi1 and wmi2 for matching.
// Empty wmi2 is normalized to empty string.
func makeWMIKey(wmi1, wmi2 string) string {
	wmi1 = strings.ToUpper(strings.TrimSpace(wmi1))
	wmi2 = strings.ToUpper(strings.TrimSpace(wmi2))
	return wmi1 + wmi2
}

// readExistingOutput reads existing output file (JSONL format) and returns
// a map keyed by wmi1+wmi2. Malformed lines are logged and skipped.
func readExistingOutput(outputFile string) (map[string]*vinv1.Manufacturer, error) {
	if outputFile == "-" {
		// Can't read from stdout
		return make(map[string]*vinv1.Manufacturer), nil
	}

	existing := make(map[string]*vinv1.Manufacturer)

	f, err := os.Open(outputFile)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet, return empty map
			return existing, nil
		}
		return nil, fmt.Errorf("failed to open existing output file %s: %w", outputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var manufacturer vinv1.Manufacturer
		if err := protojson.Unmarshal([]byte(line), &manufacturer); err != nil {
			slog.Warn("skipping malformed JSON line in existing output", "line", lineNum, "error", err)
			continue
		}

		wmi1 := manufacturer.GetWmi1()
		wmi2 := manufacturer.GetWmi2()
		key := makeWMIKey(wmi1, wmi2)

		if key != "" {
			existing[key] = &manufacturer
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read existing output file: %w", err)
	}

	slog.Info("read existing output file", "outputFile", outputFile, "records", len(existing))
	return existing, nil
}
