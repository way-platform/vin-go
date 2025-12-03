package volvotrucksvin

import (
	"github.com/way-platform/vin-go/internal/iso3779"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// Position 4: Cab Type / Model Family
// Reference: docs/deep-research/volvo-trucks-yv2.md Section 3.1
// Modern definition (post-2000): T=FL, V=FE, R=FH, A=FM/FMX
var globalModelSeriesMap = map[byte]vinv1.Model{
	'T': vinv1.Model_FL, // FL (Low Tilt) - Modern definition
	'V': vinv1.Model_FE, // FE Series
	'R': vinv1.Model_FH, // FH Series (High Tilt)
	'A': vinv1.Model_FM, // FM / FMX Series
	// Note: Historical (pre-1998) T could be FH12, but we prioritize modern definition
}

// Position 8: Chassis Configuration (Axles/Brakes)
// Reference: docs/deep-research/volvo-trucks-yv2.md Section 3.3, Table 2
// Returns axle count based on chassis configuration and model
func getGlobalAxleCountFromPos8(pos8 byte, model vinv1.Model) int32 {
	switch pos8 {
	case 'A':
		// A = 4x2 (Pneumatic)
		return 2
	case 'B':
		// B = 4x4 (Pneumatic)
		return 2
	case 'C':
		// C = 6x2 (Pneumatic)
		return 3
	case 'D':
		// D = 6x4 (Pneumatic)
		return 3
	case 'F':
		// F = 8x2 (Pneumatic)
		return 4
	case 'G':
		// G = 8x4 (Pneumatic)
		return 4
	default:
		// Numbers 1-9 indicate hydraulic brake systems (Light/Medium Duty)
		// Typically found on FL/FE series.
		// Reference: docs/deep-research/volvo-trucks-yv2.md Section 3.3 "Light/Medium Duty (FL)"
		if pos8 >= '1' && pos8 <= '9' {
			if model == vinv1.Model_FL || model == vinv1.Model_FE {
				return 2
			}
		}
		return 0
	}
}

// decodeGlobalVDS decodes the Global/European Vehicle Descriptor Section (Positions 4-9).
// Returns model, fuel types, axle count, and year.
// Reference: docs/deep-research/volvo-trucks-yv2.md Section 3
func decodeGlobalVDS(vin string, wmi string) (vinv1.Model, []vinv1.FuelType, int32, int32) {
	var model vinv1.Model
	var fuelTypes []vinv1.FuelType
	var axleCount int32
	var year int32

	if len(vin) < 17 {
		return model, fuelTypes, axleCount, year
	}

	// Only decode for Volvo Trucks, not Buses
	if wmi == "YV3" {
		return model, fuelTypes, axleCount, year
	}

	// Position 4: Cab Type / Model Family
	if m, ok := globalModelSeriesMap[vin[3]]; ok {
		model = m
	}

	// Position 5-7: Engine codes (composite)
	// Extract positions 5, 6, 7 for engine code
	if len(vin) >= 8 {
		engineCode := vin[4:7]
		fuelTypes = inferFuelTypeFromGlobalEngineCode(engineCode)
	}

	// Position 8: Chassis Configuration -> Axle Count
	axleCount = getGlobalAxleCountFromPos8(vin[7], model)

	// Position 10: Model Year (ISO 3779)
	if y, ok := iso3779.Year(vin[9]); ok {
		year = int32(y)
	}

	return model, fuelTypes, axleCount, year
}

