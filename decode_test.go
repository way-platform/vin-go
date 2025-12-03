package vin

import (
	"bytes"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

var update = flag.Bool("update", false, "update golden files")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestDecodeGoldenFiles(t *testing.T) {
	// Find all .txt files in testdata directory
	txtFiles, err := filepath.Glob("testdata/*.txt")
	if err != nil {
		t.Fatalf("Failed to glob testdata/*.txt: %v", err)
	}

	if len(txtFiles) == 0 {
		t.Skip("No .txt files found in testdata directory")
	}

	for _, txtFile := range txtFiles {
		t.Run(filepath.Base(txtFile), func(t *testing.T) {
			// Read the .txt file
			content, err := os.ReadFile(txtFile)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", txtFile, err)
			}

			// Split into lines and decode each VIN
			lines := strings.Split(string(content), "\n")
			var decodedVINs []*vinv1.Vin

			for _, line := range lines {
				vinStr := strings.TrimSpace(line)
				if vinStr == "" {
					continue
				}

				decoded, err := Decode(vinStr)
				if err != nil {
					t.Errorf("Failed to decode VIN %s: %v", vinStr, err)
					continue
				}
				decodedVINs = append(decodedVINs, decoded)
			}

			// Marshal each VIN to JSON using protojson
			marshaler := protojson.MarshalOptions{}
			var jsonObjects []json.RawMessage
			for _, vin := range decodedVINs {
				vinJSON, err := marshaler.Marshal(vin)
				if err != nil {
					t.Fatalf("Failed to marshal VIN to JSON: %v", err)
				}
				jsonObjects = append(jsonObjects, vinJSON)
			}

			// Combine into a JSON array
			jsonBytes, err := json.Marshal(jsonObjects)
			if err != nil {
				t.Fatalf("Failed to marshal VIN array to JSON: %v", err)
			}

			// Format JSON with consistent indentation
			var formattedJSON bytes.Buffer
			if err := json.Indent(&formattedJSON, jsonBytes, "", "  "); err != nil {
				t.Fatalf("Failed to indent JSON: %v", err)
			}

			// Determine golden file path
			goldenFile := strings.TrimSuffix(txtFile, ".txt") + ".json"

			if *update {
				// Update golden file
				if err := os.WriteFile(goldenFile, formattedJSON.Bytes(), 0o644); err != nil {
					t.Fatalf("Failed to write golden file %s: %v", goldenFile, err)
				}
			} else {
				// Compare with golden file
				expected, err := os.ReadFile(goldenFile)
				if err != nil {
					t.Fatalf("Failed to read golden file %s: %v", goldenFile, err)
				}

				if diff := cmp.Diff(string(expected), formattedJSON.String()); diff != "" {
					t.Errorf("Golden file mismatch for %s (-want +got):\n%s", goldenFile, diff)
				}
			}
		})
	}
}
