package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// decodeBaumuster decodes the model from a Baumuster series and subtype.
// series is the 3-digit Baumuster series (VIN positions 3-6).
// subtype is the 3-digit subtype (VIN positions 6-9).
// wmi is used for context-specific disambiguation (e.g., 447 series).
func decodeBaumuster(series string, subtype string, wmi string) vinv1.Model {
	switch series {
	case "447":
		// 447 is Vito/V-Class or Metris (US/Canada)
		// US Market
		if wmi == "W1W" || wmi == "WD4" || wmi == "WDA" || wmi == "WD3" {
			return vinv1.Model_METRIS
		}
		// W1K is typical for US/Global Passenger Car V-Class
		if wmi == "W1K" || wmi == "WMX" {
			return vinv1.Model_V_CLASS
		}
		// WDF, W1V, VSA are typically Vito (Commercial) or V-Class (Passenger)
		// Defaulting to Vito is safer for a commercial decoder.
		return vinv1.Model_VITO

	case "638", "639":
		return vinv1.Model_VITO

	case "901", "902", "903", "904", "905", "906":
		return vinv1.Model_SPRINTER

	case "907":
		return vinv1.Model_SPRINTER

	case "910":
		// VS30 FWD
		// Check for eSprinter (Gen 1) using subtype
		if isElectricSubtype(series, subtype) {
			return vinv1.Model_E_SPRINTER
		}
		return vinv1.Model_SPRINTER

	case "415":
		return vinv1.Model_CITAN

	case "420":
		// Citan / T-Class
		return vinv1.Model_CITAN

	// Heavy Trucks
	case "930", "932", "933", "934": // Actros MP1-3
		return vinv1.Model_ACTROS
	case "963": // Actros MP4/5 / Antos
		return vinv1.Model_ACTROS

	case "983": // eActros
		return vinv1.Model_E_ACTROS

	case "964":
		return vinv1.Model_AROCS

	case "213": // E-Class
		return vinv1.Model_E_CLASS

	case "950", "952", "953", "954":
		return vinv1.Model_AXOR

	case "956", "957", "958":
		return vinv1.Model_E_ECONIC

	case "967", "970", "972", "974", "975", "976":
		return vinv1.Model_ATEGO

	case "949", "959":
		return vinv1.Model_ZETROS

	// Unimog (400 series, exclude 447)
	case "404", "405", "406", "416", "424", "425", "435", "437":
		return vinv1.Model_UNIMOG
	}

	return vinv1.Model_MODEL_UNSPECIFIED
}

// isElectricSubtype checks if a Baumuster subtype indicates an electric variant.
// Research indicates: "eSprinter typically occupies the 910.6xx range"
// VIN: W1V 910 6 33... -> subtype starts with '6' at position 6 (0-indexed).
func isElectricSubtype(series string, subtype string) bool {
	if len(subtype) == 0 {
		return false
	}

	switch series {
	case "910":
		// eSprinter Gen 1: subtype starts with '6' (910.6xx)
		// Subtype is at VIN positions 6-9 (0-indexed: 6,7,8)
		// First character of subtype (position 6) is '6'
		return subtype[0] == '6'
	case "447":
		// eVito/EQV subtypes: Research mentions 447.605, 447.705 for electric variants
		// These don't overlap with OM654 diesel variants (447.601)
		// Check if subtype starts with '6' or '7' (indicating electric range)
		if len(subtype) >= 1 {
			firstDigit := subtype[0]
			return firstDigit == '6' || firstDigit == '7'
		}
	}

	return false
}

