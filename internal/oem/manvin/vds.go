package manvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// vdsInfo holds the decoded vehicle attributes from VDS positions 4-6.
type vdsInfo struct {
	model       vinv1.Model
	vehicleType vinv1.VehicleType
	fuelTypes   []vinv1.FuelType
	axleCount   int32
}

// VDS model codes (positions 4-6).
// Reference: docs/deep-research/man.md Section "The Vehicle Descriptor Section"
//
// Positions 7-8 are ZZ filler (European standard padding, not data).
// Position 9 is the check digit.
var vdsMap = map[string]vdsInfo{
	// MAN TGE / eTGE light commercial vehicle.
	// 2-axle platform (FWD, RWD, or 4MOTION).
	// Default fuel: DIESEL. eTGE (electric) shares the same code.
	"06K": {
		model:       vinv1.Model_TGE,
		vehicleType: vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
		fuelTypes:   []vinv1.FuelType{vinv1.FuelType_DIESEL},
		axleCount:   2,
	},

	// MAN heavy-duty container truck (TGS/TGX family).
	// Cannot distinguish TGS vs TGX from VIN alone — leave model unset.
	// Axle count varies (2-4) — leave unset.
	"28S": {
		vehicleType: vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
		fuelTypes:   []vinv1.FuelType{vinv1.FuelType_DIESEL},
	},
}

// decodeVDS decodes the Vehicle Descriptor Section (positions 4-9).
// Returns model, vehicle type, fuel types, and axle count where determinable.
func decodeVDS(vin string) (vinv1.Model, vinv1.VehicleType, []vinv1.FuelType, int32) {
	if len(vin) < 9 {
		return vinv1.Model_MODEL_UNSPECIFIED, vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED, nil, 0
	}

	code := vin[3:6] // Positions 4-6
	if info, ok := vdsMap[code]; ok {
		return info.model, info.vehicleType, info.fuelTypes, info.axleCount
	}

	return vinv1.Model_MODEL_UNSPECIFIED, vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED, nil, 0
}
