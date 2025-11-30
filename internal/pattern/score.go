package pattern

import (
	"strings"
)

// Matches returns true if the input VIN satisfies the pattern rules.
// The input 'vin' is expected to be a full 17-character VIN string.
func Matches(pattern, vin string) bool {
	segment, isPlant := extractSegment(pattern, vin)
	if segment == "" {
		return false
	}
	return matchesSegment(pattern, segment, isPlant)
}

// Score calculates the confidence level (0.0 to 1.0) of the match.
// The input 'vin' is expected to be a full 17-character VIN string.
// Returns 0.0 if the pattern does not match or vin is invalid.
func Score(pattern, vin string) float64 {
	segment, isPlant := extractSegment(pattern, vin)
	if segment == "" {
		return 0.0
	}
	
	// Ensure it's a match before calculating score
	if !matchesSegment(pattern, segment, isPlant) {
		return 0.0
	}

	return scoreSegment(pattern, segment, isPlant)
}

// extractSegment determines the correct substring of the VIN to use based on the pattern.
// Returns the segment and a boolean indicating if it's a Plant Code pattern.
func extractSegment(pattern, vin string) (string, bool) {
	if len(vin) != 17 {
		return "", false
	}

	if strings.Contains(pattern, "|") {
		// For VIS patterns involving plant code, we use the 11th character (index 10)
		return string(vin[10]), true
	}

	// For standard patterns, we use the VIN starting from position 4 (index 3)
	// This covers VDS + Check Digit + VIS
	return vin[3:], false
}

func matchesSegment(pattern, input string, isPlant bool) bool {
	if isPlant {
		before, after, found := strings.Cut(pattern, "|")
		if found && len(before) == 5 {
			visPattern := after
			
			if len(input) == 0 {
				return false
			}
			plantCodeChar := rune(input[0])

			// Check metadata part of pattern (e.g., "*U")
			if len(visPattern) >= 2 && visPattern[0] == '*' {
				expectedPlantCode := rune(visPattern[1])
				if expectedPlantCode == '*' || plantCodeChar == expectedPlantCode {
					return true
				}
			}
			return false
		}
	}

	if len(pattern) == 0 {
		return len(input) == 0
	}

	patternIndex := 0
	inputIndex := 0

	for patternIndex < len(pattern) && inputIndex < len(input) {
		patternChar := rune(pattern[patternIndex])
		inputChar := rune(input[inputIndex])

		if patternChar == '[' {
			// Character class match
			localCloseBracket := strings.Index(pattern[patternIndex:], "]")
			if localCloseBracket == -1 {
				return false
			}

			charClassEnd := patternIndex + localCloseBracket + 1
			if charClassEnd > len(pattern) {
				return false
			}

			charClass := pattern[patternIndex:charClassEnd]
			content := charClass[1 : len(charClass)-1]

			if !isCharInRange(inputChar, content) {
				return false
			}

			patternIndex = charClassEnd
			inputIndex++
		} else if patternChar == '*' {
			// Wildcard match
			// Optimization: if wildcard is the last character, it matches the rest of input
			if patternIndex == len(pattern)-1 {
				return true
			}
			patternIndex++
			inputIndex++
		} else {
			// Exact character match
			if inputChar != patternChar {
				return false
			}
			patternIndex++
			inputIndex++
		}
	}

	// If we have consumed the entire pattern, it is a match.
	// Note: The loop handles the case where the last pattern character is a wildcard.
	if patternIndex == len(pattern) {
		return true
	}
	
	return false
}

func scoreSegment(pattern, input string, isPlant bool) float64 {
	if isPlant {
		before, after, found := strings.Cut(pattern, "|")
		if found && len(before) == 5 {
			visPattern := after
			if len(input) == 0 {
				return 0.0
			}
			plantCodeChar := rune(input[0])

			if len(visPattern) >= 2 && visPattern[0] == '*' {
				expectedPlantCode := rune(visPattern[1])
				if expectedPlantCode == '*' {
					return 0.8
				}
				if plantCodeChar == expectedPlantCode {
					return 1.0
				}
			}
			return 0.0
		}
	}

	if len(pattern) == 0 {
		return 0.0
	}

	totalScore := 0.0
	comparisons := 0

	patternIndex := 0
	inputIndex := 0

	for patternIndex < len(pattern) && inputIndex < len(input) {
		patternChar := rune(pattern[patternIndex])
		inputChar := rune(input[inputIndex])

		if patternChar == '[' {
			localCloseBracket := strings.Index(pattern[patternIndex:], "]")
			if localCloseBracket == -1 {
				return 0.0
			}

			charClassEnd := patternIndex + localCloseBracket + 1
			if charClassEnd > len(pattern) {
				return 0.0
			}

			charClass := pattern[patternIndex:charClassEnd]
			content := charClass[1 : len(charClass)-1]

			// Check for match (matchesSegment already validated this generally)
			if isCharInRange(inputChar, content) {
				comparisons++
				if strings.ContainsRune(content, '-') {
					totalScore += 0.7
				} else {
					totalScore += 0.8
				}
			}

			patternIndex = charClassEnd
			inputIndex++
		} else if patternChar == '*' {
			comparisons++
			totalScore += 0.5
			patternIndex++
			inputIndex++
		} else {
			if inputChar == patternChar {
				comparisons++
				totalScore += 1.0
			}
			patternIndex++
			inputIndex++
		}
	}

	if comparisons == 0 {
		return 0.0
	}

	return totalScore / float64(comparisons)
}

func isCharInRange(char rune, content string) bool {
	for i := 0; i < len(content); {
		if i+2 < len(content) && content[i+1] == '-' {
			// Range like A-E
			start := rune(content[i])
			end := rune(content[i+2])
			if char >= start && char <= end {
				return true
			}
			i += 3
		} else {
			// Individual character
			if char == rune(content[i]) {
				return true
			}
			i++
		}
	}
	return false
}