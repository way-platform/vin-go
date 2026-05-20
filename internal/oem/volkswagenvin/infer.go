package volkswagenvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a supported Volkswagen VIN.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]
	var brand = vinv1.Brand_VOLKSWAGEN
	var model = vinv1.Model_MODEL_UNSPECIFIED
	var vehicleType = vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED

	isVW := false

	// 1. Identify Volkswagen WMI
	switch wmi {
	// Germany
	case "WV1": // Commercial (Goods)
		isVW = true
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
	case "WV2": // Commercial (Passenger/MPV)
		isVW = true
		vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
	case "WV3": // Incomplete
		isVW = true
		vehicleType = vinv1.VehicleType_INCOMPLETE_VEHICLE
	case "WVG": // SUVs / MPVs
		isVW = true
		vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
	case "WVW": // Passenger Cars
		isVW = true
		vehicleType = vinv1.VehicleType_PASSENGER_CAR
	case "WV4": // Commercial (Ford-VW alliance special bodies)
		isVW = true
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE

	// Czech Republic
	case "TMB": // Skoda
		isVW = true
		brand = vinv1.Brand_SKODA
		vehicleType = vinv1.VehicleType_PASSENGER_CAR

	// Spain
	case "VWV":
		isVW = true

	// Poland
	case "SBS":
		isVW = true
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE

	// Mexico
	case "3VW":
		isVW = true

	// Brazil
	case "9BW":
		isVW = true
	}

	// Check for generic Polish VW if not caught above (S.. range can be tricky, sticking to explicit or known prefixes if strictly needed)
	// But for now, the list above covers the primary European commercial ones mentioned in docs.

	if !isVW {
		return nil, false
	}

	// 2. Decode Model (VDS Positions 7-8)
	// Note: Positions are 0-indexed, so 7-8 corresponds to indices 6 and 7.
	// However, the docs often refer to "Positions 7 and 8".
	// In a 17-char VIN, indices 6 and 7.
	// Let's verify: WMI is 0-2. VDS is 3-8.
	// "Positions 7 and 8" usually means the 7th and 8th characters *of the VIN*.
	// Index 6 and 7.
	// Example T5: WV1 ZZZ 7H Z...
	// WV1 (0,1,2)
	// ZZZ (3,4,5)
	// 7H (6,7) -> Correct.

	modelCode := vin[6:8]

	switch modelCode {
	// Transporter Lineage
	case "70", "7D": // T4
		model = vinv1.Model_TRANSPORTER
	case "7H", "7J": // T5, T6, T6.1
		model = vinv1.Model_TRANSPORTER

	// Crafter Lineage
	case "2E", "2F": // Crafter Gen 1
		model = vinv1.Model_CRAFTER
	case "SY", "SZ": // Crafter Gen 2
		model = vinv1.Model_CRAFTER

	// Caddy Lineage
	case "2K": // Caddy Mk 3 / Mk 4
		model = vinv1.Model_CADDY
	case "SA": // Caddy Mk 4 (Internal code, sometimes appears)
		model = vinv1.Model_CADDY
	case "SB": // Caddy Mk 5
		model = vinv1.Model_CADDY
	case "SK": // Caddy 4th gen (MQB)
		model = vinv1.Model_CADDY

	// T7 Multivan
	case "ST", "TV": // ST = internal code, TV = Ford-VW alliance code
		model = vinv1.Model_MULTIVAN
	// ID. Buzz
	case "EB":
		model = vinv1.Model_ID_BUZZ
	// Amarok
	case "2H":
		model = vinv1.Model_AMAROK

	// Skoda Octavia (MQB Evo)
	case "NX":
		model = vinv1.Model_OCTAVIA

	// T-Roc
	case "A1":
		model = vinv1.Model_T_ROC

	// Passat
	case "3C":
		model = vinv1.Model_PASSAT
	}

	// 3. Refine Vehicle Type based on Model if not set by WMI
	if vehicleType == vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
		switch model {
		case vinv1.Model_TRANSPORTER, vinv1.Model_CRAFTER, vinv1.Model_CADDY:
			vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
		}
	}

	// T7 Multivan and ID. Buzz are always MPV regardless of WMI.
	if model == vinv1.Model_MULTIVAN || model == vinv1.Model_ID_BUZZ {
		vehicleType = vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE
	}
	// Amarok is always LCV.
	if model == vinv1.Model_AMAROK && vehicleType == vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
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

	return builder.Build(), true
}
