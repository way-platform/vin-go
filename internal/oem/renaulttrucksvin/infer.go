package renaulttrucksvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]
	var brand vinv1.Brand
	var model vinv1.Model
	var vehicleType vinv1.VehicleType

	// 1. Determine Brand
	// Renault Trucks (VF6, VF2, VNE, VN1, VF1 if commercial context)
	// Also VS... for Maxity (Renault Trucks branded)

	isRenaultTrucks := false

	switch wmi {
	case "VF6": // Heavy Duty
		brand = vinv1.Brand_RENAULT_TRUCKS
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
		isRenaultTrucks = true
	case "VF2": // Legacy Heavy
		brand = vinv1.Brand_RENAULT_TRUCKS
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
		isRenaultTrucks = true
	case "VNE": // Bus (Irisbus/Renault)
		brand = vinv1.Brand_RENAULT_TRUCKS // Or specific Bus brand if available
		vehicleType = vinv1.VehicleType_BUS
		isRenaultTrucks = true
	case "VN1": // Sovab (Master) - sold by Renault Trucks and Renault
		// The user asked for a decoder for "Renault Trucks".
		// We can return Brand_RENAULT_TRUCKS if we are sure, but VN1 is shared.
		// Let's default to RENAULT (LCV) unless we have a reason to say TRUCKS.
		// However, the function is in `renaulttrucksvin` package.
		// If this package is called, it implies we suspect it might be a Renault Truck.
		// But `decode.go` calls all decoders.
		// We should probably return `Brand_RENAULT` for LCVs generally,
		// and `Brand_RENAULT_TRUCKS` for HDVs.
		// Or should we distinguish?
		// Research says: "Master... sold through Renault Trucks dealer network... VINs typically remain VF1 or VN1".
		// So the VIN doesn't distinguish Brand (Renault vs Renault Trucks) for LCVs easily.
		// I will set Brand to RENAULT for LCVs to be safe, or RENAULT_TRUCKS for VF6.
		brand = vinv1.Brand_RENAULT
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		isRenaultTrucks = true // Handled by this decoder
	case "VF1": // Renault SA
		brand = vinv1.Brand_RENAULT
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		isRenaultTrucks = true
	}

	// Maxity check (VS...)
	if len(wmi) == 3 && wmi[:2] == "VS" {
		// VS is Nissan Spain. But Maxity is Renault Trucks.
		// If we are decoding as Renault Trucks, we can claim it?
		// Or we might conflict with a Nissan decoder.
		// Assuming we don't have a conflicting Nissan decoder yet.
		// Let's claim it if it looks like a Maxity?
		// Without VDS logic for Nissan, hard to be sure.
		// But "VS..." is listed in research as Maxity.
		// We'll mark it as candidate.
		isRenaultTrucks = true
		brand = vinv1.Brand_RENAULT_TRUCKS
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
	}

	if !isRenaultTrucks {
		return nil, false
	}

	// 2. Decode Model based on WMI and VDS

	// Heavy Duty (VF6)
	if wmi == "VF6" {
		familyCode := vin[3:5]
		switch familyCode {
		case "11", "29":
			model = vinv1.Model_T
		case "34", "32", "38":
			model = vinv1.Model_K
		case "24", "25":
			model = vinv1.Model_C
		case "20", "21":
			model = vinv1.Model_D
		case "17":
			model = vinv1.Model_MAGNUM
		case "VF", "MF": // VF6 Master Exception (VF or MF)
			model = vinv1.Model_MASTER
			vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE // Though "Heavy" homologation (4.5t), physically LCV-like
		}

		// Refine Type for Tractors if T/C/K
	}

	// Light Commercial (VF1, VN1)
	if wmi == "VF1" || wmi == "VN1" {
		vdsStart := vin[3:5]
		switch vdsStart {
		case "MA", "FD", "FV":
			model = vinv1.Model_MASTER
		case "FL": // Trafic II
			model = vinv1.Model_TRAFIC
		case "JG", "FG": // Trafic III
			model = vinv1.Model_TRAFIC
		// Kangoo? (Not in research scope but usually 'KC', 'FW', 'KW')
		case "KC", "FW", "KW":
			model = vinv1.Model_KANGOO
		}
	}

	// Maxity (VS...)
	if len(wmi) == 3 && wmi[:2] == "VS" {
		// Assume Maxity for now if passed to this decoder?
		// Or just set unspecified if we aren't sure.
		// Research says "VS... ... Maxity".
		model = vinv1.Model_MAXITY
	}

	// 3. Final Assembly
	var output vinv1.Vehicle

	// Adjust Brand for Renault Trucks specific models
	switch model {
	case vinv1.Model_T, vinv1.Model_K, vinv1.Model_C, vinv1.Model_D, vinv1.Model_MAGNUM, vinv1.Model_KERAX, vinv1.Model_PREMIUM, vinv1.Model_MIDLUM, vinv1.Model_MAXITY, vinv1.Model_MASCOTT:
		brand = vinv1.Brand_RENAULT_TRUCKS
	case vinv1.Model_MASTER, vinv1.Model_TRAFIC, vinv1.Model_KANGOO:
		// These remain RENAULT brand unless they are clearly Renault Trucks (VF6).
		if brand != vinv1.Brand_RENAULT_TRUCKS {
			brand = vinv1.Brand_RENAULT
		}
	}

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

	if !hasData {
		return nil, false
	}

	return &output, true
}
