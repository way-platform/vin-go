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
	"strings"
	"sync"
	"time"

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

func init() {
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
		if len(record) != 3 {
			return fmt.Errorf("invalid CSV record: %v", record)
		}
		wmi, manufacturer, url := record[0], record[1], record[2]
		if i > 0 && len(wmi) != 3 {
			return fmt.Errorf("invalid WMI: %s", wmi)
		}
		if i > 0 && manufacturer == "" {
			return fmt.Errorf("empty manufacturer")
		}
		// URL is optional, so we don't validate it
		_ = url
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
			wmi, manufacturer, url := record[0], record[1], record[2]

			// Normalize WMI value for matching
			wmiNormalized := strings.ToUpper(strings.TrimSpace(wmi))
			wmiKey := makeWMIKey(wmiNormalized, "")

			var manufacturerProto *vinv1.Manufacturer

			// Check if we have an existing record for this WMI
			existingMu.Lock()
			existing, ok := existingByWMI[wmiKey]
			existingMu.Unlock()

			if ok {
				slog.Info("reusing existing manufacturer", "wmiKey", wmiKey, "wmi", wmiNormalized)
				manufacturerProto = existing
			} else {
				// Create a unique key for caching/singleflight that includes the URL
				// This ensures that if the same manufacturer name appears with different URLs,
				// they are treated as distinct inference requests.
				cacheKey := manufacturer + "|" + url

				// Check memory cache first to avoid singleflight overhead
				cacheMu.Lock()
				cached, hit := cache[cacheKey]
				cacheMu.Unlock()

				if hit {
					manufacturerProto = cached
				} else {
					// Use singleflight to coalesce concurrent requests for the same manufacturer+url
					result, err, _ := requestGroup.Do(cacheKey, func() (interface{}, error) {
						// Double-check memory cache inside singleflight to avoid re-generation
						cacheMu.Lock()
						if cached, hit := cache[cacheKey]; hit {
							cacheMu.Unlock()
							return cached, nil
						}
						cacheMu.Unlock()

						// Generate from API
						startTime := time.Now()
						proto, err := generateManufacturer(gctx, client, header, record)
						slog.Info("generated manufacturer", "manufacturer", manufacturer, "wmiKey", wmiKey, "time", time.Since(startTime))
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
						return fmt.Errorf("failed to get or create manufacturer: %w", err)
					}
					manufacturerProto = result.(*vinv1.Manufacturer)
				}
			}

			// Clone to ensure we don't modify the cached/existing instance which might be shared
			manufacturerProto = proto.Clone(manufacturerProto).(*vinv1.Manufacturer)

			// Update manufacturer fields from input record
			manufacturerProto.SetWmi1(wmiNormalized)
			manufacturerProto.SetDisplayName(manufacturer)

			// Set data source
			manufacturerProto.SetDataSources([]vinv1.DataSource{vinv1.DataSource_WIKIPEDIA})

			// Write manufacturer to file
			data, err := protojson.MarshalOptions{}.Marshal(manufacturerProto)
			if err != nil {
				return fmt.Errorf("failed to marshal manufacturer: %w", err)
			}

			outputMu.Lock()
			_, err = fmt.Fprintln(out, string(data))
			if err == nil && outputFile != "-" {
				// Sync to ensure data is flushed to disk immediately
				// We do this inside the lock to ensure integrity, though frequent Sync might be slow.
				// Given the low throughput (limit 5), this is acceptable for safety.
				err = out.Sync()
			}
			outputMu.Unlock()

			if err != nil {
				return fmt.Errorf("failed to write manufacturer: %w", err)
			}

			// Update existingByWMI map to track processed records
			existingMu.Lock()
			existingByWMI[wmiKey] = manufacturerProto
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
