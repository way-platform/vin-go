package main

import (
	"bufio"
	"bytes"
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
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

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

func readManufacturers(filename string) ([]*vinv1.Manufacturer, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var manufacturers []*vinv1.Manufacturer
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		var m vinv1.Manufacturer
		if err := protojson.Unmarshal([]byte(line), &m); err != nil {
			return nil, fmt.Errorf("unmarshalling JSONL line: %w", err)
		}
		manufacturers = append(manufacturers, &m)
	}
	return manufacturers, scanner.Err()
}

func splitManufacturers(manufacturers []*vinv1.Manufacturer) ([]*vinv1.Manufacturer, []*vinv1.Manufacturer) {
	var standard []*vinv1.Manufacturer
	var lvm []*vinv1.Manufacturer

	for _, m := range manufacturers {
		if m.GetWmi2() == "" {
			standard = append(standard, m)
		} else {
			lvm = append(lvm, m)
		}
	}

	// Sort standard by Base36(WMI1)
	sort.Slice(standard, func(i, j int) bool {
		keyI, _ := wmi.ToBase36(standard[i].GetWmi1())
		keyJ, _ := wmi.ToBase36(standard[j].GetWmi1())
		return keyI < keyJ
	})

	// Sort LVM by Packed Key
	sort.Slice(lvm, func(i, j int) bool {
		kI, _ := wmi.Pack(lvm[i].GetWmi1(), lvm[i].GetWmi2())
		kJ, _ := wmi.Pack(lvm[j].GetWmi1(), lvm[j].GetWmi2())
		return kI < kJ
	})

	return standard, lvm
}

func generateStandardBlobAndIndex(manufacturers []*vinv1.Manufacturer, outputPath string) ([]byte, []StandardIndexEntry, error) {
	var buf bytes.Buffer
	var standardIndex []StandardIndexEntry

	for _, m := range manufacturers {
		// Record offset before writing
		currentOffset := uint32(buf.Len())

		key, ok := wmi.ToBase36(m.GetWmi1())
		if !ok {
			return nil, nil, fmt.Errorf("invalid WMI1 for base36 conversion: %s", m.GetWmi1())
		}
		standardIndex = append(standardIndex, StandardIndexEntry{K: key, O: currentOffset})

		// Clone and strip fields that can be inferred from the key
		clone := proto.Clone(m).(*vinv1.Manufacturer)
		clone.ClearWmi1()
		clone.ClearWmi2()
		clone.ClearLowVolume()

		data, err := proto.Marshal(clone)
		if err != nil {
			return nil, nil, fmt.Errorf("marshalling proto: %w", err)
		}

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

func generateLVMBlobAndIndex(manufacturers []*vinv1.Manufacturer, outputPath string) ([]byte, []LVMIndexEntry, error) {
	var buf bytes.Buffer
	var lvmIndex []LVMIndexEntry

	for _, m := range manufacturers {
		// Record offset before writing
		currentOffset := uint32(buf.Len())

		k, ok := wmi.Pack(m.GetWmi1(), m.GetWmi2())
		if !ok {
			return nil, nil, fmt.Errorf("invalid WMI for base36 conversion: %s%s", m.GetWmi1(), m.GetWmi2())
		}
		lvmIndex = append(lvmIndex, LVMIndexEntry{K: k, O: currentOffset})

		// Clone and strip fields that can be inferred from the key
		clone := proto.Clone(m).(*vinv1.Manufacturer)
		clone.ClearWmi1()
		clone.ClearWmi2()
		clone.ClearLowVolume()

		data, err := proto.Marshal(clone)
		if err != nil {
			return nil, nil, fmt.Errorf("marshalling proto: %w", err)
		}

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