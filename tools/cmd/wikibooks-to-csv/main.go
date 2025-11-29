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

	header := []string{"WMI", "WMI2", "Manufacturer", "URL"}
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

		// Skip lines where WMI is less than 3 characters (these are section headers)
		// Check the original wmiPart before parsing to catch header lines like "1G"
		if len(wmiPart) < 3 {
			continue
		}

		// Clean up manufacturer string
		manufacturerPart = strings.ReplaceAll(manufacturerPart, "<br>", " ")

		// Clean up manufacturer string and extract URLs
		manufacturerPart, urlPart := cleanMarkdownAndExtractURL(manufacturerPart)

		// Parse and expand multiple WMIs
		wmis := parseWMIs(wmiPart)

		// Skip if any parsed WMI is less than 3 characters (after normalization)
		skipLine := false
		for _, wmi := range wmis {
			wmiNormalized := normalizeWMI(strings.TrimSpace(wmi))
			// Extract WMI1 if it contains "/"
			if strings.Contains(wmiNormalized, "/") {
				wmiSplit := strings.Split(wmiNormalized, "/")
				if len(wmiSplit[0]) < 3 {
					skipLine = true
					break
				}
			} else if len(wmiNormalized) < 3 {
				skipLine = true
				break
			}
		}
		if skipLine {
			continue
		}

		for _, wmi := range wmis {
			wmi1 := ""
			wmi2 := ""
			if strings.Contains(wmi, "/") {
				wmiSplit := strings.Split(wmi, "/")
				wmi1 = normalizeWMI(strings.TrimSpace(wmiSplit[0]))
				if len(wmiSplit) > 1 {
					wmi2 = normalizeWMI(strings.TrimSpace(wmiSplit[1]))
				}
			} else {
				wmi1 = normalizeWMI(strings.TrimSpace(wmi))
			}

			if wmi1 == "" {
				continue
			}

			record := []string{wmi1, wmi2, manufacturerPart, urlPart}
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

// normalizeWMI normalizes WMI characters by converting similar-looking Unicode characters
// to their ASCII equivalents (e.g., Greek Η to Latin H).
func normalizeWMI(wmi string) string {
	var result strings.Builder
	for _, r := range wmi {
		switch r {
		case 'Η': // Greek capital letter eta (U+0397)
			result.WriteRune('H')
		case 'η': // Greek small letter eta (U+03B7)
			result.WriteRune('H')
		case 'Ι': // Greek capital letter iota (U+0399)
			result.WriteRune('I')
		case 'ι': // Greek small letter iota (U+03B9)
			result.WriteRune('I')
		case 'Ο': // Greek capital letter omicron (U+039F)
			result.WriteRune('O')
		case 'ο': // Greek small letter omicron (U+03BF)
			result.WriteRune('O')
		case 'Α': // Greek capital letter alpha (U+0391)
			result.WriteRune('A')
		case 'α': // Greek small letter alpha (U+03B1)
			result.WriteRune('A')
		case 'Β': // Greek capital letter beta (U+0392)
			result.WriteRune('B')
		case 'β': // Greek small letter beta (U+03B2)
			result.WriteRune('B')
		case 'Ε': // Greek capital letter epsilon (U+0395)
			result.WriteRune('E')
		case 'ε': // Greek small letter epsilon (U+03B5)
			result.WriteRune('E')
		case 'Ζ': // Greek capital letter zeta (U+0396)
			result.WriteRune('Z')
		case 'ζ': // Greek small letter zeta (U+03B6)
			result.WriteRune('Z')
		case 'Κ': // Greek capital letter kappa (U+039A)
			result.WriteRune('K')
		case 'κ': // Greek small letter kappa (U+03BA)
			result.WriteRune('K')
		case 'Μ': // Greek capital letter mu (U+039C)
			result.WriteRune('M')
		case 'μ': // Greek small letter mu (U+03BC)
			result.WriteRune('M')
		case 'Ν': // Greek capital letter nu (U+039D)
			result.WriteRune('N')
		case 'ν': // Greek small letter nu (U+03BD)
			result.WriteRune('N')
		case 'Ρ': // Greek capital letter rho (U+03A1)
			result.WriteRune('P')
		case 'ρ': // Greek small letter rho (U+03C1)
			result.WriteRune('P')
		case 'Τ': // Greek capital letter tau (U+03A4)
			result.WriteRune('T')
		case 'τ': // Greek small letter tau (U+03C4)
			result.WriteRune('T')
		case 'Υ': // Greek capital letter upsilon (U+03A5)
			result.WriteRune('Y')
		case 'υ': // Greek small letter upsilon (U+03C5)
			result.WriteRune('Y')
		case 'Χ': // Greek capital letter chi (U+03A7)
			result.WriteRune('X')
		case 'χ': // Greek small letter chi (U+03C7)
			result.WriteRune('X')
		default:
			result.WriteRune(r)
		}
	}
	return result.String()
}

// parseWMIs parses a WMI string that may contain multiple WMIs separated by commas,
// spaces, and/or <br> tags, and expands ranges like "JHF-JHG" or "JH1-JH5".
// Returns a slice of individual WMI strings.
func parseWMIs(wmiStr string) []string {
	var result []string

	// Replace <br> tags with spaces for easier parsing
	wmiStr = strings.ReplaceAll(wmiStr, "<br>", " ")
	wmiStr = strings.ReplaceAll(wmiStr, "<BR>", " ")

	// Replace commas with spaces (we'll handle both separators uniformly)
	wmiStr = strings.ReplaceAll(wmiStr, ",", " ")

	// Split by whitespace (handles multiple spaces, tabs, etc.)
	parts := strings.Fields(wmiStr)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// Check if this is a range (e.g., "JHF-JHG" or "JH1-JH5")
		if strings.Contains(part, "-") {
			expanded := expandWMIRange(part)
			result = append(result, expanded...)
		} else {
			result = append(result, part)
		}
	}

	return result
}

// expandWMIRange expands a WMI range like "JHF-JHG" or "JH1-JH5" into individual WMIs.
// Assumes WMIs are 3 characters and the range is on the last character.
func expandWMIRange(rangeStr string) []string {
	var result []string

	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		// Not a valid range, return as-is
		return []string{rangeStr}
	}

	start := strings.TrimSpace(parts[0])
	end := strings.TrimSpace(parts[1])

	if len(start) != 3 || len(end) != 3 {
		// WMIs should be 3 characters, return as-is if not
		return []string{rangeStr}
	}

	// Check if first two characters match
	if start[:2] != end[:2] {
		// Range spans different prefixes, return as-is
		return []string{rangeStr}
	}

	prefix := start[:2]
	startChar := start[2]
	endChar := end[2]

	// Check if range is valid (start <= end)
	if startChar > endChar {
		// Invalid range, return as-is
		return []string{rangeStr}
	}

	// Expand the range
	for c := startChar; c <= endChar; c++ {
		expandedWMI := prefix + string(c)
		// Normalize the expanded WMI to handle any Greek letters
		result = append(result, normalizeWMI(expandedWMI))
	}

	return result
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
