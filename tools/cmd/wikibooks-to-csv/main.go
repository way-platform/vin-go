package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	inputFile := flag.String("i", "data/wikibooks/wmi.md", "Path to the input markdown file")
	outputFile := flag.String("o", "-", "Path to the output CSV file, or '-' for stdout")
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

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

	header := []string{"WMI", "WMI2", "Manufacturer"}
	if err := csvWriter.Write(header); err != nil {
		log.Fatalf("failed to write header: %v", err)
	}

	scanner := bufio.NewScanner(file)
	// Skip header and separator lines
	scanner.Scan() // Skip | WMI ... |
	scanner.Scan() // Skip | --- |

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}

		wmiPart := strings.TrimSpace(parts[1])
		manufacturerPart := strings.TrimSpace(parts[2])

		// Clean up manufacturer string
		manufacturerPart = strings.ReplaceAll(manufacturerPart, "<br>", " ")

		for strings.Contains(manufacturerPart, "[") {
			start := strings.Index(manufacturerPart, "[")
			endBracket := strings.Index(manufacturerPart, "]")
			if endBracket == -1 || endBracket < start {
				break
			}

			if endBracket+1 < len(manufacturerPart) && manufacturerPart[endBracket+1] == '(' {
				parenCount := 1
				endParen := -1
				for i := endBracket + 2; i < len(manufacturerPart); i++ {
					if manufacturerPart[i] == '(' {
						parenCount++
					} else if manufacturerPart[i] == ')' {
						parenCount--
						if parenCount == 0 {
							endParen = i
							break
						}
					}
				}

				if endParen != -1 {
					manufacturerPart = manufacturerPart[:start] + manufacturerPart[endParen+1:]
				} else {
					manufacturerPart = manufacturerPart[:start] + manufacturerPart[endBracket+1:]
				}
			} else {
				manufacturerPart = manufacturerPart[:start] + manufacturerPart[start+1:]
				manufacturerPart = strings.Replace(manufacturerPart, "]", "", 1)
			}
		}

		manufacturerPart = strings.TrimSpace(manufacturerPart)

		wmi1 := ""
		wmi2 := ""
		if strings.Contains(wmiPart, "/") {
			wmiSplit := strings.Split(wmiPart, "/")
			wmi1 = strings.TrimSpace(wmiSplit[0])
			if len(wmiSplit) > 1 {
				wmi2 = strings.TrimSpace(wmiSplit[1])
			}
		} else {
			wmi1 = wmiPart
		}

		record := []string{wmi1, wmi2, manufacturerPart}
		if err := csvWriter.Write(record); err != nil {
			log.Printf("failed to write record: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	fmt.Fprintln(os.Stderr, "Successfully converted markdown to CSV.")
}
