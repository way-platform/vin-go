package mercedesvin

import (
	"github.com/way-platform/vin-go/internal/iso3779"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
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
	var year int32
	var axleCount int32

	// Compute a year candidate from position 10 for potential US-spec VINs.
	// We do NOT assign it to `year` yet — dual-market WMIs (W1V, W1K, W1N, WDD)
	// are shared between US and EU vehicles. EU Mercedes uses position 10 for
	// steering orientation (1=LHD, 2=RHD), not model year. We only promote the
	// candidate to `year` after a US-specific decoding strategy succeeds.
	var usYearCandidate int32
	if isUSVehicleWMI(wmi) {
		if y, ok := iso3779.Year(vin[9]); ok {
			usYearCandidate = int32(y)
		}

		// Extract Axle Count from VIN Position 7 - North American specific
		// Codes: 3, 6 (Class 3 4x2/4x4); D-H (Class D-H 4x2); R-V (Class D-H 4x4)
		// All these represent 2-axle vehicles.
		switch vin[6] {
		case '3', '6',
			'D', 'E', 'F', 'G', 'H',
			'R', 'S', 'T', 'U', 'V':
			axleCount = 2
		}
	}

	// EU Axle Count extraction for Mercedes-Benz trucks.
	// Tier 1: alphabetic 4th digit encodes drive configuration directly.
	// Tier 2: numeric 4th digit requires full Baumuster suffix lookup.
	if !isUSVehicleWMI(wmi) {
		digit4 := vin[3]
		if digit4 >= 'A' && digit4 <= 'Z' {
			if ac, ok := decodeEUAxleFromDriveCode(digit4); ok {
				axleCount = ac
			}
		} else if digit4 >= '0' && digit4 <= '9' {
			if ac, ok := decodeEUAxleFromBaumuster(vin[3:9]); ok {
				axleCount = ac
			}
		}
	}

	// Strategy A: North American Commercial (Sprinter/Metris) 2-digit codes
	// This is only attempted if the WMI is specifically for US Commercial vehicles.
	if isUSCommercialWMI(wmi) {
		code2 := vin[3:5]
		usModel := decodeModelUS(code2)
		if usModel != vinv1.Model_MODEL_UNSPECIFIED {
			model = usModel
			fuelTypes = decodeFuelTypeUS(code2)
			year = usYearCandidate
		}
	}

	// Strategy B: Standard Baumuster (3-digit) Decoding (for EU commercial/cars, and some US if Strategy A fails)
	// This is tried if no model was found by Strategy A.
	if model == vinv1.Model_MODEL_UNSPECIFIED {
		model = decodeModelEU(vin, wmi)
		// If a model was successfully decoded via Baumuster, try to infer fuel type
		if model != vinv1.Model_MODEL_UNSPECIFIED && len(vin) >= 9 {
			series := vin[3:6]
			subtype := vin[6:9]
			fuelTypes = decodeFuelTypeBaumuster(series, subtype)
		}
	}

	// Strategy C: W1V Attribute Logic (Mercedes-Benz Vans Germany)
	// This is tried if no model was found by previous strategies and WMI is W1V.
	// This handles European Sprinters (VS30) and Vitos (W447) that use attribute codes instead of Baumuster.
	if model == vinv1.Model_MODEL_UNSPECIFIED && wmi == "W1V" {
		var w1vFuel []vinv1.FuelType
		model, w1vFuel = decodeAttributesW1V(vin)
		if len(w1vFuel) > 0 {
			fuelTypes = w1vFuel
		}
	}

	// Strategy D: US Passenger Car (W1K, W1N, WDD with US VDS rules)
	// This runs if model is still unspecified, and we know it's a potential US-spec
	// vehicle. If the US passenger model decode succeeds, we confirm position 10
	// as a valid year code.
	if model == vinv1.Model_MODEL_UNSPECIFIED && isUSVehicleWMI(wmi) && usYearCandidate > 0 {
		usModel := decodeUSPassengerCarModel(vin[3], usYearCandidate)
		if usModel != vinv1.Model_MODEL_UNSPECIFIED {
			model = usModel
			year = usYearCandidate
		}
	}

	// Special case: W1H is explicitly Freightliner Econic (always tried regardless of previous results if WMI matches)
	// This might override a generic model, but often W1H would not decode a model via other strategies.
	if wmi == "W1H" {
		model = vinv1.Model_E_ECONIC
	}

	// 3. Determine Vehicle Type (depends on model and WMI)
	vehicleType := determineVehicleType(model, wmi)

	// 4. Refine Vehicle Type for US passenger vans (position 6: F, G → PASSENGER_CAR)
	// This must run after model and vehicleType are initially set.
	if len(vin) > 5 {
		vehicleType = refineVehicleTypeForUSPassengerVan(vehicleType, vin[5], model, wmi)
	}

	// 5. Final Fuel Type Backup: Use model annotations if no fuel type was determined yet.
	if len(fuelTypes) == 0 && model != vinv1.Model_MODEL_UNSPECIFIED {
		enumValueDescriptor := model.Descriptor()
		if enumValueDescriptor != nil {
			if opts := enumValueDescriptor.Options(); opts != nil {
				if proto.HasExtension(opts, vinv1.E_FuelType) {
					ext := proto.GetExtension(opts, vinv1.E_FuelType)
					if fts, ok := ext.([]vinv1.FuelType); ok && len(fts) > 0 {
						fuelTypes = fts
					}
				}
			}
		}
	}

	// 6. Build output
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
		output.SetModelYear(year)
	}
	if axleCount > 0 {
		output.SetAxleCount(axleCount)
	}
	output.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})
	hasData := brand != vinv1.Brand_BRAND_UNSPECIFIED ||
		model != vinv1.Model_MODEL_UNSPECIFIED ||
		vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED ||
		len(fuelTypes) > 0 ||
		year > 0 ||
		axleCount > 0
	return &output, hasData
}
