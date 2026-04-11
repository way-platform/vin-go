package manvin

import (
	"github.com/way-platform/vin-go/internal/iso3779"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported MAN VIN.
// Reference: docs/deep-research/man.md
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]

	// 1. Determine Brand from WMI
	brand, knownWMI := getBrandFromWMI(wmi)
	if !knownWMI {
		return nil, false
	}

	// 2. Decode VDS (Vehicle Descriptor Section, positions 4-9)
	model, vehicleType, fuelTypes, axleCount := decodeVDS(vin)

	// 3. Decode VIS (Vehicle Identifier Section, positions 10-17)
	var year int32
	if y, ok := iso3779.Year(vin[9]); ok {
		year = int32(y)
	}

	// 4. Assemble Vehicle object
	hasData := brand != vinv1.Brand_BRAND_UNSPECIFIED ||
		model != vinv1.Model_MODEL_UNSPECIFIED ||
		vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED ||
		len(fuelTypes) > 0 ||
		axleCount > 0 ||
		year > 0

	if !hasData {
		return nil, false
	}

	builder := vinv1.Vehicle_builder{
		DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
	}
	if brand != vinv1.Brand_BRAND_UNSPECIFIED {
		builder.Brand = new(brand)
	}
	if model != vinv1.Model_MODEL_UNSPECIFIED {
		builder.Model = new(model)
	}
	if vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
		builder.Type = new(vehicleType)
	}
	if len(fuelTypes) > 0 {
		builder.FuelTypes = fuelTypes
	}
	if axleCount > 0 {
		builder.AxleCount = new(axleCount)
	}
	if year > 0 {
		builder.ModelYear = new(year)
	}

	return builder.Build(), true
}
