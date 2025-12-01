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
	var brand vinv1.Brand
	var vehicleType vinv1.VehicleType
	var model vinv1.Model

	knownWMI := false

	// 1. Determine Brand and initial Type hints from WMI
	switch wmi {
	// Mercedes-Benz
	case "VSA", "WDB", "WDF", "WDD", "WDC", "W1T", "W1V", "W1W", "W1X", "W1Y", "W1Z",
		"W1K", "W1N", "WMX", "4JG", // Passenger/US additions
		"WD3", "WD4", "WDA", "WDZ", "1MB",
		"8AB", "8AC", "8BN", "8BR", "8BT", "8BU", "9BM", "NMB":
		brand = vinv1.Brand_MERCEDES_BENZ
		knownWMI = true

	// Freightliner
	case "W1H", "WD6", "WD7", "WDR", "WDP", "WDY", "WCD",
		"W2W", "W2X", "W2Y", "W2Z":
		brand = vinv1.Brand_FREIGHTLINER
		knownWMI = true

	// Dodge
	case "WD0", "WD8", "WDW", "WDX":
		brand = vinv1.Brand_DODGE
		knownWMI = true

	// Ambiguous (Dodge or Freightliner)
	case "WD1", "WD2", "WD5":
		brand = vinv1.Brand_DODGE
		knownWMI = true
	}

	if !knownWMI {
		return nil, false
	}

	// 2. Determine Model

	// Strategy A: North American Commercial (Sprinter/Metris) 2-digit codes
	// Applicable primarily to W1W, W1Y, WD3, WD4, WDA... and potentially Dodge/Freightliner variants.
	// We check the 2-digit code first if it matches known patterns.

	isUSCommercialCandidate := false
	switch wmi {
	case "W1V", "W1W", "W1Y", "W1X", "WD3", "WD4", "WDA", "WDZ", "WD0", "WD8", "WDW", "WDX", "WD1", "WD2", "WD5", "W2W", "W2X":
		isUSCommercialCandidate = true
	}

	if isUSCommercialCandidate {
		code2 := vin[3:5]
		switch code2 {
		// VS30 Sprinter Codes (907)
		case "40", "70": // Gas
			model = vinv1.Model_SPRINTER
		case "4D", "5D", "8D", "9D": // OM651 Diesel
			model = vinv1.Model_SPRINTER
		case "4E", "5E", "8E", "9E": // OM642 Diesel
			model = vinv1.Model_SPRINTER
		case "4K", "5K", "4N": // OM654 Diesel
			model = vinv1.Model_SPRINTER
		case "3H": // Likely Sprinter 1500/2500 Gas or newer code
			model = vinv1.Model_SPRINTER
		case "4V": // eSprinter
			model = vinv1.Model_E_SPRINTER

		// Metris (447) US Code
		case "V0":
			model = vinv1.Model_METRIS
		}
	}

	// Strategy B: Standard Baumuster (3-digit) Decoding
	// If Model is still unspecified, try the standard 3-digit lookup.
	if model == vinv1.Model_MODEL_UNSPECIFIED {
		modelCode := vin[3:6]

		switch modelCode {
		case "447":
			// 447 is Vito/V-Class or Metris (US/Canada)
			// US Market
			if wmi == "W1W" || wmi == "WD4" || wmi == "WDA" || wmi == "WD3" {
				model = vinv1.Model_METRIS
			} else if wmi == "W1K" || wmi == "WMX" {
				// W1K is typical for US/Global Passenger Car V-Class
				model = vinv1.Model_V_CLASS
			} else {
				// WDF, W1V, VSA are typically Vito (Commercial) or V-Class (Passenger)
				// Research says: "V-Class (Passenger): Often W1V 447... or W1K 447..."
				// "Vito (Commercial): Often WDF 447... or W1V 447..."
				// Since W1V overlaps, we need a tiebreaker.
				// Often V-Class has distinct subtypes or is MPV.
				// Without subtype data, we default to Vito for WDF/VSA.
				// For W1V, it's ambiguous. Defaulting to Vito is safer for a commercial decoder.
				model = vinv1.Model_VITO
			}
		case "638", "639":
			model = vinv1.Model_VITO
		case "901", "902", "903", "904", "905", "906":
			model = vinv1.Model_SPRINTER
		case "907":
			model = vinv1.Model_SPRINTER
		case "910":
			// VS30 FWD
			// Check for eSprinter (Gen 1)
			// Subtype is at VIN positions 6-8 (0-based check: 6,7,8)
			// Research: "eSprinter typically occupies the 910.6xx range"
			// VIN: W1V 910 6 33... -> 6 at index 6.
			if len(vin) > 6 && vin[6] == '6' {
				model = vinv1.Model_E_SPRINTER
			} else {
				model = vinv1.Model_SPRINTER
			}

		case "415":
			model = vinv1.Model_CITAN
		case "420":
			// Citan / T-Class
			model = vinv1.Model_CITAN

		// Heavy Trucks
		case "930", "932", "933", "934": // Actros MP1-3
			model = vinv1.Model_ACTROS
		case "963": // Actros MP4/5 / Antos
			model = vinv1.Model_ACTROS

		case "964":
			model = vinv1.Model_AROCS

		case "950", "952", "953", "954":
			model = vinv1.Model_AXOR

		case "956", "957", "958":
			model = vinv1.Model_E_ECONIC

		case "967", "970", "972", "974", "975", "976":
			model = vinv1.Model_ATEGO

		case "949", "959":
			model = vinv1.Model_ZETROS

		// Unimog (400 series, exclude 447)
		case "404", "405", "406", "416", "424", "425", "435", "437":
			model = vinv1.Model_UNIMOG
		}
	}

	// W1H is explicitly Freightliner Econic
	if wmi == "W1H" {
		model = vinv1.Model_E_ECONIC
	}

	// 3. Determine Vehicle Type
	// We prefer LCV/HGV based on the Model family.

	// Model-based defaults
	switch model {
	case vinv1.Model_SPRINTER, vinv1.Model_VITO, vinv1.Model_METRIS, vinv1.Model_E_SPRINTER, vinv1.Model_CITAN, vinv1.Model_T_CLASS:
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
	case vinv1.Model_V_CLASS:
		vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
	case vinv1.Model_ACTROS, vinv1.Model_AROCS, vinv1.Model_ATEGO, vinv1.Model_E_ECONIC,
		vinv1.Model_AXOR, vinv1.Model_ZETROS, vinv1.Model_UNIMOG:
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	}

	// If Model didn't determine type (unmapped model code), fallback to WMI hints
	if vehicleType == vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
		switch wmi {
		// HGV Context
		case "W1T", "W1H", "WD6", "WDY", "W2Y", "1MB", "9BM", "NMB":
			vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
		// LCV Context (Vans/Trucks)
		case "W1V", "W1Y", "WD2", "WD3", "8AC", "8BU":
			vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		// Bus
		case "W1Z", "WDZ", "WDW", "WCD", "W2Z", "8BR":
			vehicleType = vinv1.VehicleType_BUS
		// MPV / SUV
		case "W1W", "WD4", "WD5", "WD8", "WDR", "W2W", "8BT", "W1N", "WDC", "4JG":
			vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
		// Passenger Car
		case "WDD", "W1K", "WMX":
			vehicleType = vinv1.VehicleType_PASSENGER_CAR
		// Incomplete
		case "W1X", "WDA", "WD1", "WD7", "WDX", "WDP", "W2X", "8BN":
			// Ideally, incomplete Sprinters are still LCVs logic-wise, but they are legally incomplete.
			// The user requested preferring LCV/HGV.
			// If we know it's a Sprinter (via model), it's already LCV above.
			// If model is unknown but WMI is incomplete, we default to Incomplete
			// unless we want to force LCV? Sticking to Incomplete if unknown model is safer,
			// or we could assume LCV for the smaller WMIs.
			// Let's map unknown-model Incomplete to LCV if it's a "Sprinter" WMI range?
			// For now, leave as Incomplete if model is unknown.
			vehicleType = vinv1.VehicleType_INCOMPLETE_VEHICLE
		}
	}

	// Override: If WMI is explicitly BUS, we might want to keep it as BUS even if Model is Sprinter?
	// The user said "prefer LCV/HGV types".
	// However, a Sprinter Bus is distinct from a Sprinter Van.
	// But if the Model is SPRINTER, we set LCV above.
	// Let's check if we should override LCV back to BUS if WMI demands it.
	// Use case: Regulatory vs Commercial grouping.
	// If I leave it as LCV, it's safe for "commercial vehicles".
	// I will respect the explicit WMI "Bus" types to override the default LCV assignment
	// because a Bus is not a "Goods Vehicle".
	switch wmi {
	case "W1Z", "WDZ", "WDW", "WCD", "W2Z", "8BR":
		vehicleType = vinv1.VehicleType_BUS
	}

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
	hasData := brand != vinv1.Brand_BRAND_UNSPECIFIED ||
		model != vinv1.Model_MODEL_UNSPECIFIED ||
		vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED
	return &output, hasData
}
