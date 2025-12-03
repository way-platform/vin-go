package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// decodeFuelTypeUS decodes fuel type from US 2-digit model code.
// code2 is the 2-digit code at VIN positions 3-5 (0-indexed: 3,4).
// Returns a slice to support future multi-fuel vehicles.
func decodeFuelTypeUS(code2 string) []vinv1.FuelType {
	switch code2 {
	// Gasoline codes
	case "40", "70": // M274 2.0L Gasoline
		return []vinv1.FuelType{vinv1.FuelType_GASOLINE}
	// Metris (447) US Code
	case "V0":
		return []vinv1.FuelType{vinv1.FuelType_GASOLINE}

	// Diesel codes
	case "4D", "5D", "8D", "9D": // OM651 2.2L Diesel
		return []vinv1.FuelType{vinv1.FuelType_DIESEL}
	case "4E", "5E", "8E", "9E": // OM642 3.0L Diesel
		return []vinv1.FuelType{vinv1.FuelType_DIESEL}
	case "4K", "5K", "4N", "5N", "8K", "8N", "9K", "9N": // OM654 2.0L Diesel
		return []vinv1.FuelType{vinv1.FuelType_DIESEL}

	// Electric code
	case "4V": // eSprinter
		return []vinv1.FuelType{vinv1.FuelType_ELECTRIC}
	}

	return nil
}



