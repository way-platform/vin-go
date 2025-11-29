package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	inputFile := flag.String("i", "data/wikipedia/wmi.md", "Path to the input markdown file")
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

	header := []string{"WMI", "Manufacturer", "URL"}
	if err := csvWriter.Write(header); err != nil {
		log.Fatalf("failed to write header: %v", err)
	}

	scanner := bufio.NewScanner(file)
	// Skip header and separator lines
	scanner.Scan() // Skip | WMI | Country | Manufacturer |
	scanner.Scan() // Skip | --- | --- | --- |

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}

		wmiPart := strings.TrimSpace(parts[1])

		// Skip section headers (rows starting with ###)
		if strings.HasPrefix(wmiPart, "###") {
			continue
		}

		// Skip separator rows (rows that are all dashes or empty)
		if wmiPart == "" || strings.Trim(wmiPart, "- ") == "" {
			continue
		}

		// Skip empty WMI
		if wmiPart == "" {
			continue
		}

		// Skip lines where WMI is less than 3 characters (these are section headers)
		// Check the original wmiPart before parsing to catch header lines like "JN"
		if len(wmiPart) < 3 {
			continue
		}

		// Determine manufacturer column
		// Some rows have Country in column 2, Manufacturer in column 3
		// Some rows have Manufacturer directly in column 2 (no Country)
		var manufacturerPart string
		if len(parts) >= 4 && strings.TrimSpace(parts[3]) != "" {
			manufacturerPart = strings.TrimSpace(parts[3])
		} else if len(parts) >= 3 {
			manufacturerPart = strings.TrimSpace(parts[2])
		} else {
			continue
		}

		// Clean up manufacturer string and extract URLs
		manufacturerPart, urlPart := cleanMarkdownAndExtractURL(manufacturerPart)

		// Handle multiple WMIs (space-separated or comma-separated)
		// Replace commas with spaces for uniform handling
		wmiPartNormalized := strings.ReplaceAll(wmiPart, ",", " ")
		wmis := strings.Fields(wmiPartNormalized)

		// Skip if any parsed WMI is less than 3 characters (after trimming)
		skipLine := false
		for _, wmi := range wmis {
			wmi = strings.TrimSpace(wmi)
			if wmi == "" {
				continue
			}
			if len(wmi) < 3 {
				skipLine = true
				break
			}
		}
		if skipLine {
			continue
		}

		for _, wmi := range wmis {
			wmi = strings.TrimSpace(wmi)
			if wmi == "" {
				continue
			}
			record := []string{wmi, manufacturerPart, urlPart}
			if err := csvWriter.Write(record); err != nil {
				log.Printf("failed to write record: %v", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	fmt.Fprintln(os.Stderr, "Successfully converted markdown to CSV.")
}

// cleanMarkdownAndExtractURL removes markdown links and citations from text,
// and returns the cleaned text and extracted URLs (semicolon-separated)
func cleanMarkdownAndExtractURL(text string) (string, string) {
	var urls []string

	// First, remove citation references: [[18]] or [[18]](url) -> empty string
	// Handle both standalone citations and citations followed by URLs
	// Process citations BEFORE markdown links to avoid conflicts
	for strings.Contains(text, "[[") {
		start := strings.Index(text, "[[")
		endBracket := strings.Index(text[start:], "]]")
		if endBracket == -1 {
			break
		}
		endBracket += start + 2 // Position after ]]

		// Check if followed by (url) and remove it too
		if endBracket < len(text) && text[endBracket] == '(' {
			// Find matching closing parenthesis (handle nested parentheses)
			parenCount := 1
			endParen := -1
			for i := endBracket + 1; i < len(text); i++ {
				if text[i] == '(' {
					parenCount++
				} else if text[i] == ')' {
					parenCount--
					if parenCount == 0 {
						endParen = i
						break
					}
				}
			}
			if endParen != -1 {
				// Citations don't count as URLs, just remove them
				text = text[:start] + text[endParen+1:]
			} else {
				// Malformed, just remove the citation
				text = text[:start] + text[endBracket:]
			}
		} else {
			// Just remove the citation
			text = text[:start] + text[endBracket:]
		}
	}

	// Extract markdown links: [text](url) -> text, and collect URLs
	for strings.Contains(text, "[") {
		start := strings.Index(text, "[")
		endBracket := strings.Index(text[start:], "]")
		if endBracket == -1 {
			break
		}
		endBracket += start

		// Check if followed by (url)
		if endBracket+1 < len(text) && text[endBracket+1] == '(' {
			// Find matching closing parenthesis
			parenCount := 1
			endParen := -1
			for i := endBracket + 2; i < len(text); i++ {
				if text[i] == '(' {
					parenCount++
				} else if text[i] == ')' {
					parenCount--
					if parenCount == 0 {
						endParen = i
						break
					}
				}
			}

			if endParen != -1 {
				// Extract link text (between [ and ])
				linkText := text[start+1 : endBracket]
				// Extract URL (between ( and ))
				url := text[endBracket+2 : endParen]
				// Remove URL parameters/queries if present (keep only the base URL)
				// URLs in markdown can have spaces and quotes, e.g., (https://example.com "title")
				if spaceIdx := strings.Index(url, " "); spaceIdx != -1 {
					url = url[:spaceIdx]
				}
				// Remove quotes and angle brackets if present
				url = strings.Trim(url, `"'<>`)
				if url != "" {
					urls = append(urls, url)
				}
				// Replace [text](url) with just text
				text = text[:start] + linkText + text[endParen+1:]
			} else {
				// Malformed, just remove the brackets
				text = text[:start] + text[start+1:]
				text = strings.Replace(text, "]", "", 1)
			}
		} else {
			// Simple [text] without URL - remove brackets but keep text
			linkText := text[start+1 : endBracket]
			text = text[:start] + linkText + text[endBracket+1:]
		}
	}

	// Clean up multiple spaces
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)

	// Join URLs with semicolon and space
	urlStr := strings.Join(urls, "; ")
	return text, urlStr
}
