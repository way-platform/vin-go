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

type MissingSegment struct {
	Start         int
	End           int
	Size          int
	LowerNeighbor int
	LowerFile     string
	UpperNeighbor int
	UpperFile     string
}

func main() {
	inputDir := flag.String("d", "data/kba/pages", "Path to the directory containing KBA CSV files")
	flag.Parse()

	kbaIndex := make(map[int]string)
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

			if len(record) > 0 && record[0] != "" {
				kba, err := strconv.Atoi(record[0])
				if err != nil {
					log.Printf("failed to parse KBA number from record in file %s: %v", file, err)
					continue
				}
			kbaIndex[kba] = filepath.Base(file)
			}
		}
		f.Close()
	}

	var keys []int
	for k := range kbaIndex {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var missingSegments []MissingSegment
	maxKBA := 9970
	i := 1
	for i <= maxKBA {
		if _, exists := kbaIndex[i]; !exists {
			segmentStart := i
			segmentEnd := i

			for {
				if _, exists := kbaIndex[segmentEnd+1]; exists || (segmentEnd+1) > maxKBA {
					break
				}
			segmentEnd++
			}

			segment := MissingSegment{
				Start: segmentStart,
				End:   segmentEnd,
				Size:  segmentEnd - segmentStart + 1,
			}

			insertionPos := sort.SearchInts(keys, segmentStart)

			if insertionPos > 0 {
				lowerNeighborKBA := keys[insertionPos-1]
				segment.LowerNeighbor = lowerNeighborKBA
				segment.LowerFile = kbaIndex[lowerNeighborKBA]
			}

			if insertionPos < len(keys) {
				upperNeighborKBA := keys[insertionPos]
				segment.UpperNeighbor = upperNeighborKBA
				segment.UpperFile = kbaIndex[upperNeighborKBA]
			}
			missingSegments = append(missingSegments, segment)

			i = segmentEnd + 1
		} else {
			i++
		}
	}

	sort.Slice(missingSegments, func(i, j int) bool {
		return missingSegments[i].Size > missingSegments[j].Size
	})

	for _, segment := range missingSegments {
		if segment.Size == 1 {
			fmt.Printf("Missing KBA: %d (size: 1)\n", segment.Start)
		} else {
			fmt.Printf("Missing contiguous segment: %d-%d (size: %d)\n", segment.Start, segment.End, segment.Size)
		}

		if segment.LowerFile != "" {
			fmt.Printf("  Nearest smaller KBA: %d in file: %s\n", segment.LowerNeighbor, segment.LowerFile)
		} else {
			fmt.Printf("  No smaller KBA found.\n")
		}

		if segment.UpperFile != "" {
			fmt.Printf("  Nearest larger KBA: %d in file: %s\n", segment.UpperNeighbor, segment.UpperFile)
		} else {
			fmt.Printf("  No larger KBA found.\n")
		}
	}
}