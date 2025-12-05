package scaniavin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported Scania VIN.
// Reference: docs/deep-research/scania.md
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
	model, vdsVehicleType, axleCount := decodeVDS(vin)

	// 3. Decode VIS (Vehicle Identifier Section, positions 10-17)
	year, _ := decodeVIS(vin)

	// 4. Determine Vehicle Type (refine based on WMI and chassis adaptation)
	var pos4, pos5 byte
	if len(vin) >= 5 {
		pos4 = vin[3]
		pos5 = vin[4]
	}
	vehicleType := determineVehicleType(wmi, model, pos4, pos5)
	// If VDS already determined vehicle type from series, use that
	if vdsVehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
		vehicleType = vdsVehicleType
	}

	// 5. Assemble Vehicle object
	var output vinv1.Vehicle

	if brand != vinv1.Brand_BRAND_UNSPECIFIED {
		output.SetBrand(brand)
	}

	if model != vinv1.Model_MODEL_UNSPECIFIED {
		output.SetModel(model)
	}

	if vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
		output.SetType(vehicleType)
	}
	output.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})

	if year > 0 {
		output.SetYear(year)
	}

	if axleCount > 0 {
		output.SetAxleCount(axleCount)
	}

	// Return true if we have at least brand information
	hasData := brand != vinv1.Brand_BRAND_UNSPECIFIED ||
		model != vinv1.Model_MODEL_UNSPECIFIED ||
		vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED ||
		year > 0 ||
		axleCount > 0

	if !hasData {
		return nil, false
	}

	return &output, true
}
