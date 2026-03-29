package toyotavin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported Toyota VIN.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]
	
	// Check if it is a Toyota WMI.
	// List based on research and common Toyota WMIs.
	isToyota := false
	isStellantisPlatform := false // For ProAce family (YAR)

	switch wmi {
	case "YAR": // Toyota Motor Europe (Stellantis Platform)
		isToyota = true
		isStellantisPlatform = true
	case "JT1", "JT2", "JT3", "JT4", "JT5", "JT6", "JT7", "JT8", "JTA", "JTB", "JTC", "JTD", "JTE", "JTF", "JTG", "JTH", "JTJ", "JTK", "JTL", "JTM", "JTN", "JTP", "JTR", "JTS", "JTT", "JTU", "JTV", "JTW", "JTX", "JTY", "JTZ":
		// Japan (simplified check, usually JT is enough but let's be safe or just check prefix JT)
		isToyota = true
	case "SB1": // UK
		isToyota = true
	case "VNK": // France
		isToyota = true
	case "NMT": // Turkey
		isToyota = true
	case "AHT", "LTV", "MRO", "MR0": // Others mentioned in some contexts, but sticking to confirmed ones.
		// Add more if needed.
	}

	// Broad check for JT if not caught above
	if !isToyota && wmi[0:2] == "JT" {
		isToyota = true
	}
	// Check for South Africa (AHT), USA (4T), etc?
	// For now, focusing on the European scope of the request + generic Toyota.
	// 4T1, 4T3, etc are Toyota USA.
	if !isToyota && (wmi[0:2] == "4T" || wmi[0:2] == "5T" || wmi[0:2] == "2T") {
		isToyota = true
	}
	// South Africa
	if !isToyota && wmi == "AHT" {
		isToyota = true
	}

	if !isToyota {
		return nil, false
	}

	var output vinv1.Vehicle
	output.SetBrand(vinv1.Brand_TOYOTA)
	output.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})

	if isStellantisPlatform {
		decodeStellantisPlatform(vin, &output)
	} else {
		decodeNativePlatform(vin, &output)
	}

	return &output, true
}

func decodeStellantisPlatform(vin string, output *vinv1.Vehicle) {
	// WMI: YAR
	// VDS Pos 4 (Index 3): Platform
	// VDS Pos 6-8 (Index 5-7): Engine

	platformCode := vin[3]
	engineCode := vin[5:8]

	switch platformCode {
	case 'K': // EMP2-K9 -> ProAce City
		output.SetModel(vinv1.Model_PROACE_CITY)
		output.SetType(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE)
	case 'V': // EMP2-K0 -> ProAce / ProAce Verso
		output.SetModel(vinv1.Model_PROACE)
		// Default to LCV, though Verso is MPV. Can't easily distinguish without more data or specific body codes (Pos 5).
		output.SetType(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE)
	case 'E': // K0 platform → ProAce (same as Citroën Dispatch/Peugeot Expert)
		output.SetModel(vinv1.Model_PROACE)
		output.SetType(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE)
	case 'M': // X250 -> ProAce Max (Fiat Ducato based)
		// Model_PROACE_MAX not in enum yet, fall back to PROACE or leave unspecified?
		// Leaving model unspecified but brand Toyota.
		output.SetType(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE)
	}

	// Powertrain/Fuel Logic
	// Z in Pos 6 (Index 5) usually means Electric in PSA logic.
	if engineCode[0] == 'Z' {
		output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_ELECTRIC})
	} else {
		// Default to Diesel for these vans if not electric
		// AC3 = 1.5L Diesel
		// HVM = 2.0L Diesel
		// Check for specific codes
		if engineCode == "AC3" || engineCode == "HVM" || engineCode == "YHVM" || (len(engineCode) >= 2 && engineCode[0:2] == "YH") {
			output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_DIESEL})
		}
	}
}

func decodeNativePlatform(vin string, output *vinv1.Vehicle) {
	// Native Toyota Logic
	// Focus on models identified in research: Yaris, Yaris Cross, Corolla.

	vds := vin[3:9]
	// Pos 4 (Index 3): Body/Drive
	// Pos 5 (Index 4): Engine
	// Pos 7 (Index 6): Model/Platform differentiator
	
	// Yaris Cross: VNK KD 3 D 3 ...
	// KD3D3
	// Pos 4='K' (Body), Pos 5='D' (Engine M15A-FXE), Pos 7='D' (Yaris Cross)

	// Corolla Commercial: SB1 Z 9 3 BE ...
	// Z93BE
	// Pos 4='Z' (Wagon), Pos 5='9' (Series), Pos 6-7='BE' (1.8L Hybrid)

	// Heuristic Matching

	// Yaris / Yaris Cross
	// XP210 Platform
	if (vds[0] == 'K' || vds[0] == 'P' || vds[0] == 'M') && (vds[1] == 'D' || vds[1] == 'A' || vds[1] == 'H' || vds[1] == 'B') {
		// Likely Yaris family.
		// Check for Yaris Cross specific 'D' in Pos 7 (Index 6 of VIN? No, VDS is 6 chars. Pos 7 is Index 6 of VIN, so Index 3 of VDS).
		// Wait, VDS is vin[3:9] (indices 3,4,5,6,7,8).
		// Pos 4 is vin[3] (VDS[0])
		// Pos 5 is vin[4] (VDS[1])
		// Pos 6 is vin[5] (VDS[2])
		// Pos 7 is vin[6] (VDS[3])
		// Pos 8 is vin[7] (VDS[4])
		
		if vds[3] == 'D' {
			output.SetModel(vinv1.Model_YARIS_CROSS)
			output.SetType(vinv1.VehicleType_PASSENGER_CAR)
		} else {
			// Default to Yaris for XP210 platform variants.
			output.SetModel(vinv1.Model_YARIS)
			output.SetType(vinv1.VehicleType_PASSENGER_CAR)
		}
	} else if vds[0] == 'K' && vds[1] == 'D' {
         // Specific match for VNK KD...
         if vds[3] == 'D' {
             output.SetModel(vinv1.Model_YARIS_CROSS)
         } else {
             output.SetModel(vinv1.Model_YARIS)
         }
         output.SetType(vinv1.VehicleType_PASSENGER_CAR)
    }

	// Corolla
	// E210 Platform
	if vds[0] == 'Z' && (vds[1] == '9' || vds[1] == 'E') {
		output.SetModel(vinv1.Model_COROLLA)
		// Check for Commercial
		// There isn't a robust way to distinguish Commercial from Touring Sports just by VIN VDS for SB1... 
		// except that the user research says "SB1Z93BE1..." is commercial.
		// We'll set it as Passenger Car by default unless we have a specific override, 
		// or if we want to infer LCV for all wagons from UK? No, that's risky.
		// User doc says: "The definitive check is the V5C... however based on the provided list context, we classify this as the Toyota Corolla Commercial."
		// I will leave Type as Passenger Car for Corolla unless I'm sure.
		output.SetType(vinv1.VehicleType_PASSENGER_CAR)
	}

	// Engine / Fuel
	// D in Pos 5 (VDS[1]) for Yaris/Yaris Cross = M15A-FXE (Hybrid)
	if vds[1] == 'D' || (len(vds) >= 5 && vds[3:5] == "BE") {
		// BE in Pos 7-8 (VDS[3-4]) for Corolla = Hybrid.
		// Actually doc says Pos 6-7 'BE'.
		// VDS: Z(0) 9(1) 3(2) B(3) E(4)
		// So vin[6:8] -> VDS[3:5] is 'BE'.
		output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_GASOLINE})
		// Implicitly Hybrid.
	}
}
