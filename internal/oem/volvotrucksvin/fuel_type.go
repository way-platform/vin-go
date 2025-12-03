package volvotrucksvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// inferFuelTypeFromNAEngineCode infers fuel types from North American VIN positions.
// Reference: docs/deep-research/volvo-trucks.md Section 3.4, 3.5
// Position 7: Engine manufacturer and displacement
// Position 8: Engine performance class (horsepower)
// Position 4: Model series (W = VNR Electric)
func inferFuelTypeFromNAEngineCode(pos7 byte, pos8 byte, pos4 byte) []vinv1.FuelType {
	var fuelTypes []vinv1.FuelType

	// Check Position 4 first (VNR Electric model)
	if pos4 == 'W' {
		fuelTypes = append(fuelTypes, vinv1.FuelType_ELECTRIC)
		return fuelTypes
	}

	// Check Position 8 for Electric (VNR Electric)
	if pos8 == 'N' {
		fuelTypes = append(fuelTypes, vinv1.FuelType_ELECTRIC)
		return fuelTypes
	}

	// Check Position 7 for engine type
	switch pos7 {
	case 'D', 'E', 'K', 'T', 'S', 'J':
		// D = D11, E = D13, K = D16 (Volvo Diesel)
		// T = Cummins X15, S = Cummins L9, J = Cummins N14 (Cummins Diesel)
		fuelTypes = append(fuelTypes, vinv1.FuelType_DIESEL)
	case 'V', 'U':
		// V = Cummins ISL G (CNG/LNG)
		// U = Cummins ISX12 G (CNG/LNG)
		// Natural gas engines can use both CNG and LNG
		fuelTypes = append(fuelTypes, vinv1.FuelType_COMPRESSED_NATURAL_GAS)
		fuelTypes = append(fuelTypes, vinv1.FuelType_LIQUEFIED_NATURAL_GAS)
	}

	return fuelTypes
}

// inferFuelTypeFromGlobalEngineCode infers fuel types from Global/European VIN positions 5-7.
// Reference: docs/deep-research/volvo-trucks-yv2.md Section 3.2
// Position 5-7: Composite engine codes (e.g., "2B6", "BZ0", "G30", "0P0")
func inferFuelTypeFromGlobalEngineCode(pos5to7 string) []vinv1.FuelType {
	var fuelTypes []vinv1.FuelType

	// Check for electric codes
	// Reference: docs/deep-research/volvo-trucks-yv2.md Section 3.2
	// "Electric Drivetrains: Emerging codes like 0P0 designate 'UENGINE' (Electric Motor)."
	if pos5to7 == "0P0" {
		fuelTypes = append(fuelTypes, vinv1.FuelType_ELECTRIC)
		return fuelTypes
	}

	// Known Diesel Codes from Documentation
	// Reference: docs/deep-research/volvo-trucks-yv2.md Section 3.2
	// Reference: docs/gemini-search/volvo-trucks.md
	switch pos5to7 {
	case "2B6": // D16A520
		fuelTypes = append(fuelTypes, vinv1.FuelType_DIESEL)
	case "BZ0": // D13C500
		fuelTypes = append(fuelTypes, vinv1.FuelType_DIESEL)
	case "G30": // D13K420
		fuelTypes = append(fuelTypes, vinv1.FuelType_DIESEL)
	case "P90": // D16K750
		fuelTypes = append(fuelTypes, vinv1.FuelType_DIESEL)
	case "TY0": // D13K500
		fuelTypes = append(fuelTypes, vinv1.FuelType_DIESEL)
	case "B40": // D8
		fuelTypes = append(fuelTypes, vinv1.FuelType_DIESEL)
	// Inferred Diesel Codes (Contextual: Pre-Electric FL/FH models)
	case "0X1", "0Y1", "0U1", "T40", "T60", "9J0":
		fuelTypes = append(fuelTypes, vinv1.FuelType_DIESEL)
	}

	return fuelTypes
}

