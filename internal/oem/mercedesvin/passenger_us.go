package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// decodeUSPassengerCarModel infers the model for US-market Mercedes-Benz passenger cars
// based on VIN Position 4 (Series Indicator) and Model Year.
// This logic is derived from Table 3 in baumuster.md.
func decodeUSPassengerCarModel(pos4 byte, year int32) vinv1.Model {
	switch pos4 {
	case 'A':
		if year >= 1976 && year <= 1985 { // W123 E-Class
			return vinv1.Model_E_CLASS
		}
		// Ambiguity: A also maps to W206 C-Class (2021-Present) and C199 SLR McLaren (2003-2010)
		// We don't have C_CLASS or SLR_MCLAREN enums yet, so cannot map exhaustively.

	case 'E':
		if year >= 1985 && year <= 1995 { // W124 E-Class
			return vinv1.Model_E_CLASS
		}
		// Ambiguity: E also maps to V295 EQE Sedan (2022-Present)
		// We don't have EQE_SEDAN enum yet.

	case 'H':
		if year >= 2009 && year <= 2016 { // W212 E-Class
			return vinv1.Model_E_CLASS
		}

	case 'J':
		if year >= 1995 && year <= 2002 { // W210 E-Class
			return vinv1.Model_E_CLASS
		}

	case 'L':
		if year >= 2023 { // W214 E-Class
			return vinv1.Model_E_CLASS
		}

	case 'U':
		if year >= 2002 && year <= 2009 { // W211 E-Class
			return vinv1.Model_E_CLASS
		}
	}

	return vinv1.Model_MODEL_UNSPECIFIED
}
