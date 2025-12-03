package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// determineVehicleType determines the vehicle type from model and WMI.
// Model-based defaults take precedence, with WMI-based fallbacks when model is unknown.
// Bus WMIs override model-based defaults to respect explicit bus identification.
func determineVehicleType(model vinv1.Model, wmi string) vinv1.VehicleType {
	// Explicit bus WMI override - buses are not "Goods Vehicles"
	if isBusWMI(wmi) {
		return vinv1.VehicleType_BUS
	}

	// Model-based defaults
	switch model {
	case vinv1.Model_SPRINTER, vinv1.Model_VITO, vinv1.Model_METRIS,
		vinv1.Model_E_SPRINTER, vinv1.Model_CITAN, vinv1.Model_T_CLASS:
		return vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE

	case vinv1.Model_V_CLASS:
		return vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE

	case vinv1.Model_ACTROS, vinv1.Model_AROCS, vinv1.Model_ATEGO,
		vinv1.Model_E_ECONIC, vinv1.Model_AXOR, vinv1.Model_ZETROS,
		vinv1.Model_UNIMOG, vinv1.Model_E_ACTROS:
		return vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	}

	// WMI-based fallbacks when model is unknown
	if model == vinv1.Model_MODEL_UNSPECIFIED {
		switch wmi {
		// HGV Context
		case "W1T", "W1H", "WD6", "WDY", "W2Y", "1MB", "9BM", "NMB":
			return vinv1.VehicleType_HEAVY_GOODS_VEHICLE
		// LCV Context (Vans/Trucks)
		case "W1V", "W1Y", "WD2", "WD3", "8AC", "8BU":
			return vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		// Bus
		case "W1Z", "WDZ", "WDW", "WCD", "W2Z", "8BR":
			return vinv1.VehicleType_BUS
		// MPV / SUV
		case "W1W", "WD4", "WD5", "WD8", "WDR", "W2W", "8BT", "W1N", "WDC", "4JG":
			return vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
		// Passenger Car
		case "WDD", "W1K", "WMX":
			return vinv1.VehicleType_PASSENGER_CAR
		// Incomplete
		case "W1X", "WDA", "WD1", "WD7", "WDX", "WDP", "W2X", "8BN":
			return vinv1.VehicleType_INCOMPLETE_VEHICLE
		}
	}

	return vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED
}

// refineVehicleTypeForUSPassengerVan refines vehicle type for US passenger vans.
// For US Sprinter vans, passenger vans (position 6: F, G) should be classified as PASSENGER_CAR.
// pos6 is the character at VIN position 6 (0-indexed: index 5).
func refineVehicleTypeForUSPassengerVan(vehicleType vinv1.VehicleType, pos6 byte, model vinv1.Model, wmi string) vinv1.VehicleType {
	// Only apply refinement for US Sprinter vans
	if model == vinv1.Model_SPRINTER && isUSCommercialWMI(wmi) {
		// If position 6 indicates passenger van ("F", "G"), override to PASSENGER_CAR
		if pos6 == 'F' || pos6 == 'G' {
			return vinv1.VehicleType_PASSENGER_CAR
		}
	}
	return vehicleType
}

