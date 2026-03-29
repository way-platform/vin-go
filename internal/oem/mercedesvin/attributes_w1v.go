package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// decodeAttributesW1V decodes the model and fuel type from W1V-series (Mercedes-Benz Vans) VINs
// using the Attribute-based logic (Positions 4-9).
// This is used when US 2-digit decoding and EU Baumuster decoding fail.
func decodeAttributesW1V(vin string) (vinv1.Model, []vinv1.FuelType) {
	if len(vin) < 6 {
		return vinv1.Model_MODEL_UNSPECIFIED, nil
	}

	var model vinv1.Model
	var fuelTypes []vinv1.FuelType

	// Position 4 (Index 3) determines the platform/model family
	pos4 := vin[3]

	switch pos4 {
	case '3': // 3-Series (e.g., 315, 317, 319) -> Sprinter VS30
		model = vinv1.Model_SPRINTER
		// Sprinter Fuel Logic (Position 5)
		// K = OM654 Diesel (2.0L)
		// D = OM651 Diesel (2.1L)
		// E = OM642 Diesel (3.0L V6)
		// 4 = Gasoline (M274)
		switch vin[4] {
		case 'K', 'D', 'E':
			fuelTypes = []vinv1.FuelType{vinv1.FuelType_DIESEL}
		case '4':
			fuelTypes = []vinv1.FuelType{vinv1.FuelType_GASOLINE}
		}

	case 'G': // EQV (electric V-Class W447)
		model = vinv1.Model_V_CLASS
		fuelTypes = []vinv1.FuelType{vinv1.FuelType_ELECTRIC}

	case 'T': // T-Class (consumer Citan derivative)
		model = vinv1.Model_T_CLASS

	case 'V': // V-Series -> Vito / V-Class / Metris Platform (W447)
		// Defaulting to VITO for Commercial Vans division
		model = vinv1.Model_VITO
		// W447 Fuel Logic (Pos 6-8)
		// Research indicates BEZ and CEZ differentiate output classes of the diesel engine.
		if len(vin) >= 9 {
			suffix := vin[5:8]
			switch suffix {
			case "BEZ", "CEZ":
				fuelTypes = []vinv1.FuelType{vinv1.FuelType_DIESEL}
			}
		}
	}

	return model, fuelTypes
}
