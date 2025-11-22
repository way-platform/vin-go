package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

type KbaRecord struct {
	KBA    int
	Record []string
}

func main() {
	inputDir := flag.String("d", "data/kba/pages", "Path to the directory containing KBA CSV files")
	outputFile := flag.String("o", "-", "Path to the output CSV file, or '-' for stdout")
	flag.Parse()

	var records []KbaRecord
	pagesDir := *inputDir

	files, err := filepath.Glob(filepath.Join(pagesDir, "*.csv"))
	if err != nil {
		log.Fatalf("failed to glob csv files: %v", err)
	}

	if len(files) == 0 {
		log.Fatalf("no csv files found in %s", pagesDir)
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Printf("failed to open file %s: %v", file, err)
			continue
		}

		r := csv.NewReader(f)
		isHeader := true
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("failed to read csv record from %s: %v", file, err)
				continue
			}
			if isHeader {
				isHeader = false
				continue
			}

			if len(record) > 2 && (record[1] != "" || record[2] != "") {
				kba, err := strconv.Atoi(record[0])
				if err != nil {
					// Skip records with invalid KBA
					continue
				}
				records = append(records, KbaRecord{KBA: kba, Record: record})
			}
		}
		f.Close()
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].KBA < records[j].KBA
	})

	var outputWriter io.Writer
	if *outputFile == "-" {
		outputWriter = os.Stdout
	} else {
		f, err := os.Create(*outputFile)
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}
		defer f.Close()
		outputWriter = f
	}

	csvWriter := csv.NewWriter(outputWriter)
	defer csvWriter.Flush()

	header := []string{"KBA", "WMI1", "WMI2", "Name", "FullName", "Location"}
	if err := csvWriter.Write(header); err != nil {
		log.Fatalf("failed to write header: %v", err)
	}

	for _, rec := range records {
		if err := csvWriter.Write(rec.Record); err != nil {
			log.Printf("failed to write record: %v", err)
		}
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Fatalf("failed to write CSV: %v", err)
	}
	fmt.Fprintf(os.Stderr, "Collated %d records.\n", len(records))
}
