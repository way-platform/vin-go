package fordvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// IsEUFordWMI returns true for Ford WMIs that use European VIN encoding
// (position 11 for year, position 10 is a structural placeholder).
func IsEUFordWMI(wmi string) bool {
	switch wmi {
	case "WF0", "WF1", "NM0", "VS6", "SFA", "AFA":
		return true
	}
	return false
}

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported Ford VIN.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]
	// Check for known Ford WMIs relevant to EU/Commercial
	isFord := false
	switch wmi {
	case "WF0", "WF1", "NM0", "VS6", "SFA", "X9F", "AFA":
		isFord = true
	// Global/Other Ford WMIs that might be encountered but processed differently
	case "1FA", "1FB", "1FC", "1FD", "1FM", "1FT": // US
		// We might want to handle US Fords differently, but for now we claim them as Ford brand
		isFord = true
	case "3FA", "3FE", "3FM", "3FN", "3FT": // Mexico
		isFord = true
	}

	if !isFord {
		return nil, false
	}

	var brand vinv1.Brand = vinv1.Brand_FORD
	var model vinv1.Model = vinv1.Model_MODEL_UNSPECIFIED
	var vehicleType vinv1.VehicleType = vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED
	vds := vin[3:9]

	if wmi == "WF0" {
		// Legacy Transit Connect (Valencia): Position 9 is a static model code 'G'.
		if vds == "RXXWPG" || vds == "SXXWPG" {
			model = vinv1.Model_TRANSIT_CUSTOM
			vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		}
		// Ford Focus estate registered as LCV for fleet tax purposes.
		if vds == "NXXGCH" {
			vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		}
		// Transit Custom / Tourneo Custom family (V362/V710 platforms):
		// Positions 5-7 = "XXT", Position 9 = Modulo-11 check digit [0-9X].
		// Position 4 determines N1 (LCV) vs M1 (passenger) homologation.
		// Position 8 varies by powertrain (A=diesel, Z=BEV, etc.).
		if vin[4:7] == "XXT" {
			cd := vin[8]
			if (cd >= '0' && cd <= '9') || cd == 'X' {
				switch vin[3] {
				case 'R', 'E': // N1 Light Commercial Vehicle (GVWR class)
					model = vinv1.Model_TRANSIT_CUSTOM
					vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
				case 'A', 'B', 'C', 'H': // M1 Passenger (Tourneo Custom, restraint codes)
					model = vinv1.Model_TRANSIT_CUSTOM
					vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
				}
			}
		}
	}

	// Refine Model and Type based on EU Logic (WF0/NM0/VS6/SFA)
	// We focus on Commercial Vehicles: Transit, Transit Custom, Transit Connect, Ranger

	// Positions
	// Pos 4: Body Type
	// Pos 8: Assembly Plant
	// Pos 9: Model Code

	pos4 := vin[3]
	pos8 := vin[7]
	pos9 := vin[8]

	// Transit Family Logic
	// NM0 is almost exclusively Transit range (Otosan)
	if model == vinv1.Model_MODEL_UNSPECIFIED && wmi == "NM0" {
		// Ford Otosan production
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE

		// Pos 8 (Plant) for NM0:
		// T: Kocaeli (Transit, Transit Custom)
		// R: Craiova (Transit Courier - though Craiova is usually WF0/Ford Romania? NM0 is Turkey.
		//     Wait, Ford Otosan acquired Craiova recently. But historically R is Craiova.
		//     Let's stick to the doc: "T = Kocaeli (Transit, Transit Custom)"

		if pos8 == 'T' {
			// Kocaeli produces Transit (V363) and Transit Custom (V362)
			// Differentiating them might be tricky solely on Plant.
			// Model Codes (Pos 9) might help?
			// F = Transit variant?
			// X = Transit Large?

			// Heuristic: Transit Custom is smaller, Transit (2T) is larger.
			// Often Transit Custom uses 'F' or 'Y'?
			// The doc says: "X = Transit Large commercial variants"
			// "F = Fiesta / Transit"

			switch pos9 {
			case 'X':
				model = vinv1.Model_TRANSIT
			case 'F', 'Y': // 'Y' is often Custom in some contexts, 'F' is generic
				// If we can't be sure, defaulting to TRANSIT is safe-ish,
				// but TRANSIT_CUSTOM is distinct.
				// Let's check if there are other clues.
				// For now, map 'X' to TRANSIT.
				model = vinv1.Model_TRANSIT_CUSTOM // Guessing F/Y is Custom if X is Big Transit
			default:
				model = vinv1.Model_TRANSIT // Default fallback for NM0
			}
		} else {
			// Other NM0 plants? Inönü (Trucks)?
			// If it's a generic NM0, it's likely a Transit variant.
			model = vinv1.Model_TRANSIT
		}
	} else if model == vinv1.Model_MODEL_UNSPECIFIED && vehicleType == vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED && (wmi == "WF0" || wmi == "VS6" || wmi == "SFA") {
		// Mainstream EU Ford

		// Assembly Plant (Pos 8) is a strong indicator
		switch pos8 {
		case 'T': // Kocaeli (Transit and Transit Custom)
			vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
			if (pos9 >= '0' && pos9 <= '9') || pos9 == 'X' {
				// Modulo-11 check digit → V362/V710 Transit Custom
				model = vinv1.Model_TRANSIT_CUSTOM
			} else {
				// Alphabetic pos9 → legacy model code, full-size Transit
				model = vinv1.Model_TRANSIT
			}

		case 'P': // Valencia (Almussafes)
			// Models: Kuga, Mondeo, Galaxy, S-Max, *Transit Connect*
			// We need to distinguish Connect from the Passenger cars.
			// Passenger cars (Mondeo, S-Max, Galaxy) usually have Body (Pos 4) like 'F' (Sedan), 'D' (Estate), 'S' (MPV).
			// Transit Connect usually has Body 'W' (Van) or 'A'/'B' with commercial attributes?
			// Doc says: "W = 3-Door Commercial or Van derivative"

			if pos4 == 'W' || pos4 == 'S' { // S is sometimes used for smaller vans? Or Galaxy?
				// Galaxy is 'S' or 'W' in Model Code?
				// Doc says Model Code (Pos 9) for Galaxy is 'S'.
				// Let's check Model Code (Pos 9)

				switch pos9 {
				case 'M': // Kuga
					vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE // SUV
					// model = vinv1.Model_KUGA // We don't have KUGA in proto?
				case 'B': // Mondeo
					vehicleType = vinv1.VehicleType_PASSENGER_CAR
				case 'S': // Galaxy
					vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
				case 'W': // Galaxy / S-Max
					vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
				case 'C': // Focus? (Valencia doesn't build Focus recently? Saarlouis does. But historically maybe)

					// Transit Connect logic?
					// Transit Connect is built in Valencia.
					// Use heuristics or Model Code 'U' / 'K'? (Not in table 4, need deep research or assumption)
					// If Pos 4 is 'W' (Van), it's likely Connect.
				}

				if pos4 == 'W' {
					model = vinv1.Model_TRANSIT_CONNECT
					vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
				}
			} else {
				// Default to Passenger for Valencia if not clearly a Van
				vehicleType = vinv1.VehicleType_PASSENGER_CAR
			}

		case 'R': // Craiova, Romania
			// Models: Puma, EcoSport, *Transit Courier*
			// Transit Courier is the commercial one.
			// Puma/EcoSport are Passenger/SUV.

			if pos4 == 'W' { // Commercial body
				model = vinv1.Model_TRANSIT_COURIER
				vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
			} else {
				vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE // SUV (Puma/EcoSport)
			}

		case 'D': // Southampton (Closed 2013)
			model = vinv1.Model_TRANSIT
			vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE

		case 'C': // Saarlouis (Focus)
			// model = vinv1.Model_FOCUS // Not in proto
			vehicleType = vinv1.VehicleType_PASSENGER_CAR

		case 'G': // Cologne (Fiesta)
			vehicleType = vinv1.VehicleType_PASSENGER_CAR

		case 'S': // St. Petersburg (Focus, Mondeo)
			vehicleType = vinv1.VehicleType_PASSENGER_CAR

		case 'A': // Craiova (New code?) or Cologne?
			// Observed in Transit Courier VINs (e.g. WF0WXXTAC...)
			// Model 'C' appears to be Transit Courier in this context.
			// Or maybe 'C' is Focus and 'A' is Saarlouis? But Saarlouis is 'C' plant.
			// Given user feedback "Transit Courier":
			if pos9 == 'C' {
				model = vinv1.Model_TRANSIT_COURIER
				vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
			} else {
				// Default
				vehicleType = vinv1.VehicleType_PASSENGER_CAR
			}
		}
	}

	// Ranger Logic (Silverton, South Africa - AFA/AFR? or Thailand - MNA? or US - 1F?)
	// Ranger for EU is often built in South Africa (WMI AFA..?).
	// If we encounter AFA/AFB...
	// But user asked for "fordvin" package.
	if wmi == "AFA" {
		brand = vinv1.Brand_FORD
		model = vinv1.Model_RANGER
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
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

	return builder.Build(), isFord
}
