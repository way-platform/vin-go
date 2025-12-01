package fordvin

import (
	"fmt"
	"strings"
)

// YearMap maps the VIN character (Position 11) to the Year.
// Based on Table 5.
var yearMap = map[byte]int{
	'Y': 2000,
	'1': 2001, '2': 2002, '3': 2003, '4': 2004, '5': 2005,
	'6': 2006, '7': 2007, '8': 2008, '9': 2009,
	'A': 2010, 'B': 2011, 'C': 2012, 'D': 2013, 'E': 2014,
	'F': 2015, 'G': 2016, 'H': 2017, 'J': 2018, 'K': 2019,
	'L': 2020, 'M': 2021, 'N': 2022, 'P': 2023, 'R': 2024,
	'S': 2025,
	// Extrapolated based on pattern (skipping I, O, Q, U, Z)
	'T': 2026, 'V': 2027, 'W': 2028, 'X': 2029,
}

// monthSequence is the rotational cipher sequence.
// C K D E L Y S T J U M P B R A G
const monthSequence = "CKDELYSTJUMPBRAG"

// DecodeDate decodes the Year and Month from the VIN.
// Expects the character at Position 11 (yearCode) and Position 12 (monthCode).
func DecodeDate(yearCode byte, monthCode byte) (int, int, error) {
	year, ok := yearMap[yearCode]
	if !ok {
		return 0, 0, fmt.Errorf("unknown Ford year code: %c", yearCode)
	}

	// Determine Base Offset for the Year.
	// Reference: 2025 is Year 2 in the example, starting with 'C' (Index 0).
	// Offset decreases by 4 each year.
	// Offset = (0 - 4 * (year - 2025)) % 16
	// We handle negative results correctly for modulo.

	diff := year - 2025
	offset := (-4 * diff) % 16
	if offset < 0 {
		offset += 16
	}

	// Find index of monthCode in sequence
	idx := strings.IndexByte(monthSequence, monthCode)
	if idx == -1 {
		return year, 0, fmt.Errorf("invalid Ford month code: %c", monthCode)
	}

	// Calculate Month Index (0 = Jan, 11 = Dec)
	// MonthIndex = (idx - offset) % 16
	monthIdx := (idx - offset) % 16
	if monthIdx < 0 {
		monthIdx += 16
	}

	if monthIdx >= 12 {
		return year, 0, fmt.Errorf("month code %c is not valid for year %d", monthCode, year)
	}

	return year, monthIdx + 1, nil
}
