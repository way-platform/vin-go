package scaniavin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported Scania VIN.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]
	var brand vinv1.Brand
	var vehicleType vinv1.VehicleType
	var model vinv1.Model

	knownWMI := false

	// 1. Determine Brand from WMI
	switch wmi {
	case "YS2", "YS4", // Sweden
		"9BS",        // Brazil
		"3AX", "3BE", // Mexico
		"SZA", // Poland
		"VLU": // France
		brand = vinv1.Brand_SCANIA
		knownWMI = true
	}

	// Handle Netherlands XLE fallback if not exact match in switch above (since XLE is a prefix)
	if !knownWMI && len(wmi) == 3 && wmi[:2] == "XL" {
		// Scania Netherlands often uses XLER, XLEP, XLE
		// We check if it starts with XL. Though XL could be other manufacturers,
		// within the context of "is this a Scania", we might need to be stricter if this function is called for any VIN.
		// However, usually this is called after general WMI lookup or as part of a chain.
		// Let's assume for now we only claim it if we are sure.
		// The docs mention XLER, XLEP.
		if wmi == "XLE" || wmi == "XLR" || wmi == "XLP" || wmi == "XLER" || wmi == "XLEP" {
			// Note: wmi string is 3 chars, so XLER/XLEP are impossible as 3-char WMI.
			// The doc says "XLER" is the code. ISO WMI is 3 chars.
			// The 3rd char identifies manufacturer. If manufacturer produces < 500, it's 9.
			// Netherlands is XL-XR range.
			// "XLER" implies 4 chars? No, WMI is 3.
			// Likely WMI is XLE and 4th char (part of VDS) or it's just XLE.
			// Let's check the doc again carefully.
			// "Region Code 'S–Z' (Europe)... The letter range S through Z... XL is Netherlands."
			// Table 1 says: "XLER ... Scania Production Zwolle".
			// This might mean WMI=XLE, and R is something else?
			// Or maybe it's a typo in the doc and it means WMI "XLE".
			// Let's stick to "XLE" as the WMI for Scania Netherlands based on `data/wmi.csv`.
			// `data/wmi.csv`: "XLE ... SCANIA (S) ... THE_NETHERLANDS"
		}
	}
	
	// Double check XLE from CSV data
	if wmi == "XLE" {
		brand = vinv1.Brand_SCANIA
		knownWMI = true
	}

	if !knownWMI {
		return nil, false
	}

	// 2. Determine Model and Type from Position 4 (Series)
	// Position 4 is index 3.
	seriesCode := vin[3]

	switch seriesCode {
	// Truck Series
	case 'S':
		model = vinv1.Model_S_SERIES
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	case 'R':
		model = vinv1.Model_R_SERIES
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	case 'G':
		model = vinv1.Model_G_SERIES
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	case 'P':
		model = vinv1.Model_P_SERIES
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	case 'L':
		model = vinv1.Model_L_SERIES
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	case 'T':
		model = vinv1.Model_T_SERIES
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE

	// Bus Series
	case 'K':
		model = vinv1.Model_K_SERIES
		vehicleType = vinv1.VehicleType_BUS
	case 'N':
		model = vinv1.Model_N_SERIES
		vehicleType = vinv1.VehicleType_BUS
	case 'F':
		model = vinv1.Model_F_SERIES
		vehicleType = vinv1.VehicleType_BUS
	}

	// 3. Refine Vehicle Type based on WMI if Series didn't clarify
	if vehicleType == vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
		switch wmi {
		case "YS4", "3BE": // Explicit Bus WMIs
			vehicleType = vinv1.VehicleType_BUS
		case "YS2", "3AX", "XLE", "VLU", "SZA": // Mostly Trucks (SZA/VLU context from data implies trucks/buses, but mainly trucks for Scania production units generally?)
			// SZA is Scania Slupsk (Bus bodywork often).
			// Let's check SZA in data. `data/wikipedia/manufacturers.jsonl`: "Scania Production Slupsk" -> Brands: Scania.
			// `data/wikibooks/wmi.md`: "Scania Poland".
			// Slupsk is a bus factory.
			if wmi == "SZA" {
				vehicleType = vinv1.VehicleType_BUS
			} else {
				vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
			}
		case "9BS":
			// Brazil produces both.
			vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
		}
	}

	// 4. Refine Vehicle Type based on Position 5 (Chassis Adaptation)
	// Only if strictly necessary or to override?
	// If Model is known (e.g. P_SERIES), Type is already set to HGV.
	// If Model is unknown, we rely on WMI default.
	// Let's peek at Pos 5 to confirm HGV if ambiguous.
	if vehicleType == vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED || vehicleType == vinv1.VehicleType_HEAVY_GOODS_VEHICLE {
		chassisAdaptation := vin[4]
		if chassisAdaptation == 'A' || chassisAdaptation == 'B' {
			vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
		}
	}

			// Final assembly of the Vehicle object

			var output vinv1.Vehicle

			if brand != vinv1.Brand_BRAND_UNSPECIFIED {

				output.SetBrand(brand)

			}

			if model != vinv1.Model_MODEL_UNSPECIFIED {

				output.SetModel(model)

			}

			if vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {

				output.SetVehicleType(vehicleType)

			}

		

			hasData := brand != vinv1.Brand_BRAND_UNSPECIFIED ||

				model != vinv1.Model_MODEL_UNSPECIFIED ||

				vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED

		

			if !hasData {

				return nil, false

			}

		

			return &output, true

		}

		

	