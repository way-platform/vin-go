package volvotrucksvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported Volvo VIN.
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

	var model vinv1.Model
	var fuelTypes []vinv1.FuelType
	var axleCount int32
	var year int32

	// 2. Decode VDS based on region
	if isNorthAmericanWMI(wmi) {
		model, fuelTypes, axleCount, year = decodeNAVDS(vin)
	} else if isGlobalWMI(wmi) {
		model, fuelTypes, axleCount, year = decodeGlobalVDS(vin, wmi)
	}
	// For buses (YV3), model and other attributes remain unspecified

	// 3. Determine Vehicle Type
	pos4 := byte(0)
	if len(vin) > 3 {
		pos4 = vin[3]
	}
	vehicleType := determineVehicleType(wmi, model, pos4)

	// 4. Build output with all inferred fields
	// Return true if we have at least brand information
	hasData := brand != vinv1.Brand_BRAND_UNSPECIFIED ||
		model != vinv1.Model_MODEL_UNSPECIFIED ||
		vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED ||
		len(fuelTypes) > 0 ||
		axleCount > 0 ||
		year > 0

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

	return builder.Build(), hasData
}
