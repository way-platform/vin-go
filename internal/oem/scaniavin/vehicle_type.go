package scaniavin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// determineVehicleType determines the vehicle type based on WMI, model series, and chassis adaptation.
// Reference: docs/deep-research/scania.md Sections 2.2, 3.1
func determineVehicleType(wmi string, model vinv1.Model, pos4 byte, pos5 byte) vinv1.VehicleType {
	// If model is known, vehicle type is already determined from series
	if model != vinv1.Model_MODEL_UNSPECIFIED {
		switch pos4 {
		case 'K', 'N', 'F':
			return vinv1.VehicleType_BUS
		case 'S', 'R', 'G', 'P', 'L', 'T':
			return vinv1.VehicleType_HEAVY_GOODS_VEHICLE
		}
	}

	// If model is unknown, infer from WMI
	if isBusWMI(wmi) {
		return vinv1.VehicleType_BUS
	}

	// Check position 5 (chassis adaptation) for HGV indicators
	if isHeavyGoodsVehicleChassis(pos5) {
		return vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	}

	// Default based on WMI characteristics
	// Bus-specific WMIs
	switch wmi {
	case "YS4", "3BE", "SZA":
		return vinv1.VehicleType_BUS
	}

	// Default to HGV for truck-centric WMIs
	// YS2, 3AX, XLE, VLU are primarily truck WMIs
	// 9BS produces both, but defaults to HGV
	return vinv1.VehicleType_HEAVY_GOODS_VEHICLE
}

