package volvotrucksvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// determineVehicleType determines the vehicle type based on WMI, model, and Position 4.
// Reference: docs/deep-research/volvo-trucks.md Section 2.1
func determineVehicleType(wmi string, model vinv1.Model, pos4 byte) vinv1.VehicleType {
	// Buses
	if wmi == "YV3" {
		return vinv1.VehicleType_BUS
	}

	// Incomplete vehicles: 4V5 WMI indicates incomplete/chassis-cab
	// Exception: VNR Electric (Position 4 = 'W') is a complete vehicle
	if wmi == "4V5" && pos4 != 'W' {
		return vinv1.VehicleType_INCOMPLETE_VEHICLE
	}

	// All other Volvo Trucks are Heavy Goods Vehicles
	return vinv1.VehicleType_HEAVY_GOODS_VEHICLE
}

