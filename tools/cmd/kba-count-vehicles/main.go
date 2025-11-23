package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputFile := flag.String("i", "", "Input CSV file")
	outputFile := flag.String("o", "-", "Output CSV file (defaults to stdout)")
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("Input file is required")
	}

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// reader.FieldsPerRecord = -1

	// Read header
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	manufacturerTotals := make(map[string]int)
	manufacturerNames := make(map[string]string)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(record) < 2 {
			continue
		}

		manufacturerID := record[0]
		count, err := strconv.Atoi(record[len(record)-1])
		if err != nil {
			// Skip lines where the last column is not a valid number
			continue
		}

		manufacturerTotals[manufacturerID] += count
		if len(record) > 2 {
			manufacturerNames[manufacturerID] = record[1]
		}
	}

	type kv struct {
		Key   string
		Name  string
		Value int
	}

	var sortedTotals []kv
	for k, v := range manufacturerTotals {
		sortedTotals = append(sortedTotals, kv{k, manufacturerNames[k], v})
	}

	sort.Slice(sortedTotals, func(i, j int) bool {
		return sortedTotals[i].Value > sortedTotals[j].Value
	})

	var writer *csv.Writer
	if *outputFile == "-" {
		writer = csv.NewWriter(os.Stdout)
	} else {
		outFile, err := os.Create(*outputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
		writer = csv.NewWriter(outFile)
	}
	defer writer.Flush()

	writer.Write([]string{"KBAManufacturerID", "KBAManufacturerName", "TotalCount"})

	for _, kv := range sortedTotals {
		writer.Write([]string{kv.Key, kv.Name, strconv.Itoa(kv.Value)})
	}
}
