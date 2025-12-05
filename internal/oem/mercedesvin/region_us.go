package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// decodeModelUS decodes the model using US 2-digit code strategy.
// This applies to North American Commercial vehicles (Sprinter/Metris).
// code2 is the 2-digit code at VIN positions 3-5 (0-indexed: 3,4).
func decodeModelUS(code2 string) vinv1.Model {
	switch code2 {
	// VS30 Sprinter Codes (907)
	case "40", "70": // Gas
		return vinv1.Model_SPRINTER
	case "4D", "5D", "8D", "9D": // OM651 Diesel
		return vinv1.Model_SPRINTER
	case "4E", "5E", "8E", "9E": // OM642 Diesel
		return vinv1.Model_SPRINTER
	case "4K", "5K", "4N", "5N", "8N", "9N": // OM654 Diesel
		return vinv1.Model_SPRINTER
	case "3H": // Likely Sprinter 1500/2500 Gas or newer code
		return vinv1.Model_SPRINTER
	case "4V": // eSprinter
		return vinv1.Model_E_SPRINTER

	// Metris (447) US Code
	case "V0":
		return vinv1.Model_METRIS
	}

	return vinv1.Model_MODEL_UNSPECIFIED
}
