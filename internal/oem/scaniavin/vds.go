package scaniavin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// Position 4: Model Series (Cab Architecture)
// Reference: docs/deep-research/scania.md Section 3.1
var seriesMap = map[byte]vinv1.Model{
	// Truck Series
	'S': vinv1.Model_S_SERIES, // S-Series (NTG Flagship, flat floor, introduced 2016)
	'R': vinv1.Model_R_SERIES, // R-Series (High forward-control cab, long-haul)
	'G': vinv1.Model_G_SERIES, // G-Series (Medium forward-control cab, construction/regional)
	'P': vinv1.Model_P_SERIES, // P-Series (Low forward-control cab, distribution/construction)
	'L': vinv1.Model_L_SERIES, // L-Series (Low-entry cab, urban, introduced 2017)
	'T': vinv1.Model_T_SERIES, // T-Series (Torpedo/bonneted cab, legacy, production ceased 2005)

	// Bus Series
	'K': vinv1.Model_K_SERIES, // K-Series (Longitudinally mounted engine, intercity coaches)
	'N': vinv1.Model_N_SERIES, // N-Series (Transversely mounted engine, city transit)
	'F': vinv1.Model_F_SERIES, // F-Series (Front-engine chassis, developing markets/school buses)
}

// decodeSeries decodes the model series from VIN position 4.
// Returns the model and true if a series was identified.
func decodeSeries(pos4 byte) (vinv1.Model, bool) {
	if model, ok := seriesMap[pos4]; ok {
		return model, true
	}
	return vinv1.Model_MODEL_UNSPECIFIED, false
}

// Position 5: Cab Type and Chassis Adaptation
// Reference: docs/deep-research/scania.md Section 3.2
// Chassis Adaptation Codes:
//   A: Articulated (Tractor unit)
//   B: Basic (Rigid truck)
// Duty Class Codes:
//   M: Medium duty
//   H: Heavy duty (high chassis height, construction/off-road)
//   E: Extra Heavy duty
//   L: Low chassis height (volume transport)
//   N: Normal chassis height

// isHeavyGoodsVehicleChassis checks if position 5 indicates a heavy goods vehicle chassis.
// This is used to distinguish trucks from buses when series code is unknown.
func isHeavyGoodsVehicleChassis(pos5 byte) bool {
	return pos5 == 'A' || pos5 == 'B'
}

// Positions 5-7: Drive Configuration and Axle Arrangement
// Reference: docs/deep-research/scania.md Section 3.3
// Common configurations: 4x2, 6x2, 6x4, 8x4
// The VIN encodes these configurations, but the exact mapping requires
// access to Scania's manufacturing database.
// However, real-world data (e.g., rFMS output) confirms that some Scania VINs
// explicitly encode the configuration in positions 5-7 (e.g., "4X2").

// getAxleConfiguration extracts the axle configuration code from positions 5-7.
// Returns the axle count if determinable.
func getAxleConfiguration(vin string) int32 {
	if len(vin) < 8 {
		return 0
	}

	// Check for explicit "AxB" or "AXB" pattern in positions 5-7 (Indices 4, 5, 6)
	// Example: YS2R4X2... -> Pos 5='4', Pos 6='X', Pos 7='2'
	// This indicates a "4x2" configuration.
	// Note: 4x2 = 2 axles, 6x2/6x4 = 3 axles, 8x4 = 4 axles.

	char1 := vin[4] // Position 5
	char2 := vin[5] // Position 6
	// char3 := vin[6] // Position 7 - Used for drive wheels, not needed for axle count

	if char2 == 'x' || char2 == 'X' {
		switch char1 {
		case '4':
			return 2
		case '6':
			return 3
		case '8':
			return 4
		}
	}

	return 0
}

// Position 8: Engine Family (EU) or Safety Systems (NA)
// Reference: docs/deep-research/scania.md Section 3.4
// In EU markets, Position 8 identifies the engine hardware platform:
//   - DC09 (9-liter inline-5)
//   - DC13 (13-liter inline-6)
//   - DC16 (16-liter V8)
// In NA markets, Position 8 may encode safety restraint systems.
// The exact mapping requires manufacturer data. For now, we extract
// the character for future decoding.

// getEngineFamilyCode extracts the engine family code from position 8.
// Returns the character and true if position is valid.
func getEngineFamilyCode(vin string) (byte, bool) {
	if len(vin) < 9 {
		return 0, false
	}
	return vin[7], true
}

// decodeVDS decodes the Vehicle Descriptor Section (positions 4-9).
// Returns model, vehicle type (if determinable from series), and axle count (if determinable).
// Reference: docs/deep-research/scania.md Section 3
func decodeVDS(vin string) (vinv1.Model, vinv1.VehicleType, int32) {
	var model vinv1.Model
	var vehicleType vinv1.VehicleType
	var axleCount int32

	if len(vin) < 9 {
		return model, vehicleType, axleCount
	}

	// Position 4: Series
	if m, ok := decodeSeries(vin[3]); ok {
		model = m
		// Determine vehicle type from series
		switch vin[3] {
		case 'K', 'N', 'F':
			vehicleType = vinv1.VehicleType_BUS
		case 'S', 'R', 'G', 'P', 'L', 'T':
			vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
		}
	}

	// Positions 5-7: Axle configuration
	axleCount = getAxleConfiguration(vin)

	// Position 8: Engine family (extracted for future use)
	// Future: decode engine family from this position
	_, _ = getEngineFamilyCode(vin)

	// Position 9: Check digit (handled by existing checkdigit package)

	return model, vehicleType, axleCount
}
