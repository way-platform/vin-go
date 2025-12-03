package mercedesvin

import (
	"github.com/way-platform/vin-go/internal/iso3779"
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
	var year int32
	var axleCount int32

	if isUSCommercialWMI(wmi) {
		code2 := vin[3:5]
		usModel := decodeModelUS(code2)
		if usModel != vinv1.Model_MODEL_UNSPECIFIED {
			model = usModel
			// Extract fuel type from 2-digit code
			fuelTypes = decodeFuelTypeUS(code2)

			// Extract Year (Position 10) - North American specific
			if y, ok := iso3779.Year(vin[9]); ok {
				year = int32(y)
			}

			// Extract Axle Count (Position 7) - North American specific
			// Codes: 3, 6 (Class 3 4x2/4x4); D-H (Class D-H 4x2); R-V (Class D-H 4x4)
			// All these represent 2-axle vehicles.
			switch vin[6] {
			case '3', '6',
				'D', 'E', 'F', 'G', 'H',
				'R', 'S', 'T', 'U', 'V':
				axleCount = 2
			}
		}
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
	if year > 0 {
		output.SetYear(year)
	}
	if axleCount > 0 {
		output.SetAxleCount(axleCount)
	}
	hasData := brand != vinv1.Brand_BRAND_UNSPECIFIED ||
		model != vinv1.Model_MODEL_UNSPECIFIED ||
		vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED ||
		len(fuelTypes) > 0 ||
		year > 0 ||
		axleCount > 0
	return &output, hasData
}
