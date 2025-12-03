package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// decodeModelEU decodes the model using EU Baumuster (3-digit) strategy.
// vin is the full VIN string.
// wmi is used for context-specific disambiguation.
func decodeModelEU(vin string, wmi string) vinv1.Model {
	if len(vin) < 9 {
		return vinv1.Model_MODEL_UNSPECIFIED
	}

	// Extract Baumuster series (positions 3-6, 0-indexed: 3,4,5)
	series := vin[3:6]

	// Extract subtype (positions 6-9, 0-indexed: 6,7,8)
	// Note: Position 6 overlaps with series end, so we need positions 6-9
	subtype := vin[6:9]

	return decodeBaumuster(series, subtype, wmi)
}

