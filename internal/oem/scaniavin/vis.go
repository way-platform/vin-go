package scaniavin

import (
	"github.com/way-platform/vin-go/internal/iso3779"
)

// Position 11: Plant of Manufacture
// Reference: docs/deep-research/scania.md Section 4.2
var plantCodes = map[byte]string{
	'S': "Södertälje, Sweden",     // Scania Global HQ
	'R': "Zwolle, Netherlands",     // Scania Production Zwolle (largest assembly plant)
	'9': "Angers, France",          // Scania Production Angers
	'A': "Angers, France",          // Alternative code for Angers
	'B': "São Bernardo, Brazil",    // Scania Latin America
}

// getPlantCode extracts the plant code from position 11.
// Returns the plant location name and true if a plant code was identified.
func getPlantCode(pos11 byte) (string, bool) {
	if plant, ok := plantCodes[pos11]; ok {
		return plant, true
	}
	return "", false
}

// decodeVIS decodes the Vehicle Identifier Section (positions 10-17).
// Returns model year and plant code.
// Reference: docs/deep-research/scania.md Section 4
func decodeVIS(vin string) (int32, string) {
	var year int32
	var plantCode string

	if len(vin) < 17 {
		return year, plantCode
	}

	// Position 10: Model Year (ISO 3779)
	if y, ok := iso3779.Year(vin[9]); ok {
		year = int32(y)
	}

	// Position 11: Plant of Manufacture
	if plant, ok := getPlantCode(vin[10]); ok {
		plantCode = plant
	}

	// Positions 12-17: Sequential Production Number
	// Extracted for future use (recall tracking, component matching)
	// serialNumber := vin[11:17]

	return year, plantCode
}
