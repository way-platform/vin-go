package volvotrucksvin

import (
	"github.com/way-platform/vin-go/internal/iso3779"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// Position 4: Model Series
// Reference: docs/deep-research/volvo-trucks.md Section 3.1, Table 1
var naModelSeriesMap = map[byte]vinv1.Model{
	'N': vinv1.Model_VNL,              // VNL (Long-Haul)
	'M': vinv1.Model_VNM,               // VNM (Regional Haul, Legacy)
	'R': vinv1.Model_VNR,               // VNR / VAH (Regional & Auto Hauler)
	'K': vinv1.Model_VHD,               // VHD (Vocational Heavy Duty)
	'X': vinv1.Model_VNX,               // VNX (Extreme Heavy Haul)
	'W': vinv1.Model_VNR_ELECTRIC,      // VNR Electric (BEV)
	'A': vinv1.Model_ACL,               // ACL (Autocar / Legacy Vocational)
	'S': vinv1.Model_VHD,               // VHD (Early, Legacy)
}

// Position 5: Chassis and Brake Architecture
// Reference: docs/deep-research/volvo-trucks.md Section 3.2
// Returns axle count based on chassis configuration
func getNAAxleCountFromPos5(pos5 byte) int32 {
	switch pos5 {
	case '1', '3':
		// 1 = 4x2 Class 7, 3 = 4x2 Class 8
		return 2
	case 'B':
		// B = 6x2 (steer + two rear, one driven)
		return 3
	case 'C':
		// C = 6x4 (steer + two rear, both driven)
		return 3
	case '9':
		// 9 = Multi-axle (8x4, 8x6, etc.)
		// Return minimum of 4 axles
		return 4
	default:
		return 0
	}
}

// decodeNAVDS decodes the North American Vehicle Descriptor Section (Positions 4-9).
// Returns model, fuel types, axle count, and year.
// Reference: docs/deep-research/volvo-trucks.md Section 3
func decodeNAVDS(vin string) (vinv1.Model, []vinv1.FuelType, int32, int32) {
	var model vinv1.Model
	var fuelTypes []vinv1.FuelType
	var axleCount int32
	var year int32

	if len(vin) < 17 {
		return model, fuelTypes, axleCount, year
	}

	// Position 4: Model Series
	if m, ok := naModelSeriesMap[vin[3]]; ok {
		model = m
	}

	// Position 5: Chassis/Brake Architecture -> Axle Count
	axleCount = getNAAxleCountFromPos5(vin[4])

	// Position 7: Engine
	// Position 8: Power
	// Position 4: Model (for VNR Electric)
	fuelTypes = inferFuelTypeFromNAEngineCode(vin[6], vin[7], vin[3])

	// Position 10: Model Year (ISO 3779)
	if y, ok := iso3779.Year(vin[9]); ok {
		year = int32(y)
	}

	return model, fuelTypes, axleCount, year
}

