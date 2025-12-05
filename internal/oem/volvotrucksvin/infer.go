package volvotrucksvin

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

	isNA := false
	isGlobal := false

	switch wmi {
	// North America
	case "4V1", "4V2", "4V3", "4V4", "4V5":
		brand = vinv1.Brand_VOLVO_TRUCKS
		isNA = true
	// Global / Europe / South America
	case "YV2", "YB3", "9BV":
		brand = vinv1.Brand_VOLVO_TRUCKS
		isGlobal = true
	// Buses (potentially mixed use)
	case "YV3":
		brand = vinv1.Brand_VOLVO_BUSES
		isGlobal = true
	default:
		return nil, false
	}

	// Decode Model
	if isNA {
		// North American VDS (Position 4)
		seriesCode := vin[3]
		switch seriesCode {
		case 'N':
			model = vinv1.Model_VNL
		case 'M':
			model = vinv1.Model_VNM
		case 'R':
			model = vinv1.Model_VNR // Also VAH? Doc says "R = VNR / VAH".
			// VAH is specialized. Can we distinguish?
			// Maybe deeper decoding? But doc says "R" covers both.
			// Let's default to VNR unless we find a specific VAH marker (not clear in provided text).
			// Actually, wait, "VAH ... due to shared chassis".
			// If I have to pick one, VNR is likely the primary.
			// However, I can't distinguish easily without more data.
			// I'll set VNR for now.
		case 'K':
			model = vinv1.Model_VHD
		case 'X':
			model = vinv1.Model_VNX
		case 'W':
			model = vinv1.Model_VNR_ELECTRIC
		case 'A':
			model = vinv1.Model_ACL
		case 'S':
			// "VHD (Early)"
			model = vinv1.Model_VHD
		}
	} else if isGlobal {
		// European / Global VDS (Position 4)
		// Based on docs/deep-research/volvo-trucks-yv2.md
		// Only apply this logic to Volvo Trucks, not Buses (YV3)
		if brand == vinv1.Brand_VOLVO_TRUCKS {
			seriesCode := vin[3]
			switch seriesCode {
			case 'T':
				// Modern: FL (Low Tilt).
				// Historical (Pre-1998): Could be FH12 Tractor, but we prioritize Modern definition.
				model = vinv1.Model_FL
			case 'V':
				// FE Series
				model = vinv1.Model_FE
			case 'R':
				// FH Series (High Tilt)
				model = vinv1.Model_FH
			case 'A':
				// FM / FMX Series
				// "Code A is often associated with the FM/FMX series"
				model = vinv1.Model_FM
			default:
				// Unknown or specific specialized models
				model = vinv1.Model_MODEL_UNSPECIFIED
			}
		} else {
			// For Buses (YV3) or others, we do not have specific VDS logic yet.
			model = vinv1.Model_MODEL_UNSPECIFIED
		}
	}

	// Vehicle Type
	// Most Volvos here are Heavy Goods Vehicles.
	// VNR Electric is also HGV.
	// 4V5 (Incomplete) -> Incomplete? Or HGV?
	// Doc says "4V5 ... Incomplete Vehicles or chassis-cabs... now frequently associated with VNR Electric".
	// VNR Electric is a Truck.
	// If Model is VNR_ELECTRIC, it's HGV (or TRUCK).
	// If WMI is 4V5 but model is not determined (or is?), it might be Incomplete.
	// But "Incomplete" is a regulatory state. "Heavy Goods Vehicle" is the type.
	// The proto has `INCOMPLETE_VEHICLE`.
	// If WMI == 4V5 and Model != VNR_ELECTRIC (which has 'W' at pos 4), maybe Incomplete?
	// Note: VNR Electric uses 4V4 or 4V5?
	// "4V5 ... is now frequently associated with VNR Electric".
	// "4V4 ... is the preeminent identifier for complete Volvo commercial trucks".
	// So VNR Electric can be 4V5.
	// If I identify it as VNR_ELECTRIC, I should probably set type to HGV because that's what it *is* physically, even if legally incomplete at some stage?
	// Actually, `DecodeVehicle` usually returns what the vehicle is.
	// I'll default to HGV for Volvo Trucks brand.
	// Unless it's specifically a Bus WMI (YV3).

	if wmi == "YV3" {
		vehicleType = vinv1.VehicleType_BUS
	} else {
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	}

	// Refine for Incomplete if needed?
	// If 4V5 and NOT VNR Electric (W), maybe Incomplete?
	// "4V5 ... Incomplete Vehicles".
	// I'll add logic: if 4V5 and Pos 4 != 'W', set to INCOMPLETE?
	// Or just stick to HGV. HGV is broader and usually correct for the physical object.
	// I'll stick to HGV for now to be safe, as "Incomplete" is a legal status that changes.

	var output vinv1.Vehicle
	output.SetBrand(brand)
	if model != vinv1.Model_MODEL_UNSPECIFIED {
		output.SetModel(model)
	}
	output.SetType(vehicleType)
	output.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})

	return &output, true
}
