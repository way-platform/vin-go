package stellantisvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported Stellantis (Citroën/Peugeot) VIN.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]
	var brand vinv1.Brand

	switch wmi {
	case "VF7", "VR7": // Citroën (France / Spain)
		brand = vinv1.Brand_CITROEN
	case "VF3", "VR3": // Peugeot (France / Spain)
		brand = vinv1.Brand_PEUGEOT
	default:
		return nil, false
	}

	var model vinv1.Model
	var vehicleType vinv1.VehicleType
	var fuelTypes []vinv1.FuelType

	// Position 4 (index 3) = model/platform family in PSA VDS structure.
	platformCode := vin[3]

	switch platformCode {
	case 'V': // K9 platform — Berlingo / Partner / Rifter
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		switch brand {
		case vinv1.Brand_CITROEN:
			model = vinv1.Model_BERLINGO
		case vinv1.Brand_PEUGEOT:
			model = vinv1.Model_PARTNER
		}
	case 'E': // K0 platform — Dispatch/Jumpy / Expert
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		switch brand {
		case vinv1.Brand_CITROEN:
			model = vinv1.Model_DISPATCH
		case vinv1.Brand_PEUGEOT:
			model = vinv1.Model_EXPERT
		}
	}

	// Fuel logic: position 6 (index 5) 'Z' = electric in PSA convention.
	if len(vin) > 5 && vin[5] == 'Z' {
		fuelTypes = []vinv1.FuelType{vinv1.FuelType_ELECTRIC}
	}

	builder := vinv1.Vehicle_builder{
		Brand:       new(brand),
		DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
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

	return builder.Build(), true
}
