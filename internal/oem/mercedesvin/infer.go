package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported Mercedes-Benz/Daimler VIN.
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

	// 2. Determine Model
	var model vinv1.Model
	var fuelTypes []vinv1.FuelType

	// Strategy A: North American Commercial (Sprinter/Metris) 2-digit codes
	// Check US commercial WMIs first for 2-digit code patterns
	if isUSCommercialWMI(wmi) {
		code2 := vin[3:5]
		model = decodeModelUS(code2)
		// Extract fuel type from 2-digit code
		fuelTypes = decodeFuelTypeUS(code2)
	}

	// Strategy B: Standard Baumuster (3-digit) Decoding
	// If Model is still unspecified, try the EU Baumuster system
	if model == vinv1.Model_MODEL_UNSPECIFIED {
		model = decodeModelEU(vin, wmi)
	}

	// Special case: W1H is explicitly Freightliner Econic
	if wmi == "W1H" {
		model = vinv1.Model_E_ECONIC
	}

	// 3. Determine Vehicle Type
	vehicleType := determineVehicleType(model, wmi)

	// 4. Refine Vehicle Type for US passenger vans (position 6: F, G → PASSENGER_CAR)
	if len(vin) > 5 {
		vehicleType = refineVehicleTypeForUSPassengerVan(vehicleType, vin[5], model, wmi)
	}

	// 5. Build output
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
	if len(fuelTypes) > 0 {
		output.SetFuelTypes(fuelTypes)
	}
	hasData := brand != vinv1.Brand_BRAND_UNSPECIFIED ||
		model != vinv1.Model_MODEL_UNSPECIFIED ||
		vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED ||
		len(fuelTypes) > 0
	return &output, hasData
}
