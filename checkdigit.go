package vin

import (
	"fmt"
	"strings"
)

// vinCharValues maps each valid VIN character to its numeric value for check digit calculation
var vinCharValues = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8,
	'J': 1, 'K': 2, 'L': 3, 'M': 4, 'N': 5, 'P': 7, 'R': 9,
	'S': 2, 'T': 3, 'U': 4, 'V': 5, 'W': 6, 'X': 7, 'Y': 8, 'Z': 9,
}

// vinWeights are the position weights for check digit calculation (positions 1-17)
var vinWeights = []int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}

// shouldSkipCheckDigitValidation determines if check digit validation should be skipped
// based on the VIN's WMI (World Manufacturer Identifier) structure.
//
// Rules based on ISO 3779 and regional standards:
//   - Position 1 = '1', '4', '5': North America (USA, Canada, Mexico) - MUST validate
//   - Position 1 = 'A'-'C': Africa - SKIP validation (ISO 3779, non-FMVSS)
//   - Position 1 = 'D'-'G': Central/South America - SKIP validation (ISO 3779, non-FMVSS)
//   - Position 1 = 'H'-'R': Asia - SKIP validation (regional standards, ISO 3779)
//   - Position 1 = 'S'-'Z': Europe - SKIP validation (ECE/EU directives, ISO 3779)
//   - Position 3 = '9': Low volume manufacturers - SKIP validation (ISO small-volume rules)
//   - WMI = "6ZZ": Australia government-issued - SKIP validation
func shouldSkipCheckDigitValidation(vin string) bool {
	if len(vin) < 3 {
		return false
	}
	vin = strings.ToUpper(vin)
	pos1 := vin[0] // Position 1 (geographic area indicator)
	pos3 := vin[2] // Position 3 (manufacturer identifier)
	// North America (1, 4, 5) - MUST validate (FMVSS 115 / 49 CFR Part 565)
	if pos1 == '1' || pos1 == '4' || pos1 == '5' {
		// Exception: Low volume manufacturers (position 3 = '9') skip validation
		if pos3 == '9' {
			return true
		}
		return false
	}
	// Australia government-issued (6ZZ) - SKIP validation
	if len(vin) >= 3 && vin[0:3] == "6ZZ" {
		return true
	}
	// Low volume manufacturers (position 3 = '9') - SKIP validation
	if pos3 == '9' {
		return true
	}
	// Africa (A-C) - SKIP validation
	if pos1 >= 'A' && pos1 <= 'C' {
		return true
	}
	// Central/South America (D-G) - SKIP validation
	if pos1 >= 'D' && pos1 <= 'G' {
		return true
	}
	// Asia (H-R) - SKIP validation
	if pos1 >= 'H' && pos1 <= 'R' {
		return true
	}
	// Europe (S-Z) - SKIP validation
	if pos1 >= 'S' && pos1 <= 'Z' {
		return true
	}
	return false
}

// calculateCheckDigit calculates the expected check digit for a VIN
func calculateCheckDigit(vin string) (string, error) {
	if len(vin) != 17 {
		return "", fmt.Errorf("VIN must be 17 characters")
	}

	vin = strings.ToUpper(vin)
	sum := 0

	for i, char := range vin {
		value, ok := vinCharValues[char]
		if !ok {
			return "", fmt.Errorf("invalid character '%c' at position %d", char, i+1)
		}

		// Position 9 (index 8) has weight 0, so it doesn't contribute to the sum
		sum += value * vinWeights[i]
	}

	remainder := sum % 11
	if remainder == 10 {
		return "X", nil
	}
	return fmt.Sprintf("%d", remainder), nil
}

// validateCheckDigit validates the check digit in position 9 of the VIN.
// Returns true (valid) if the VIN is from a manufacturer that doesn't use
// check digit validation, or if the check digit matches the calculated value.
func validateCheckDigit(vin string) (bool, error) {
	if len(vin) != 17 {
		return false, fmt.Errorf("VIN must be 17 characters")
	}

	// Skip check digit validation for manufacturers that don't use it
	if shouldSkipCheckDigitValidation(vin) {
		return true, nil
	}

	vin = strings.ToUpper(vin)
	actualCheckDigit := string(vin[8]) // Position 9 (0-indexed as 8)

	expectedCheckDigit, err := calculateCheckDigit(vin)
	if err != nil {
		return false, err
	}

	return actualCheckDigit == expectedCheckDigit, nil
}
