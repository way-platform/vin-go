package checkdigit

import "strings"

// ShouldSkipValidation determines if check digit validation should be skipped
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
func ShouldSkipValidation(vin string) bool {
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
