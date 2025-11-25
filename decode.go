package vin

import (
	"fmt"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// Decode validates and decodes a Vehicle Identification Number (VIN).
func Decode(vin string) (*vinv1.Vin, error) {
	if len(vin) != 17 {
		return nil, fmt.Errorf("invalid VIN length: expected 17 characters, got %d", len(vin))
	}
	if err := validateVINCharacters(vin); err != nil {
		return nil, err
	}

	// Calculate check digit
	calculatedCheckDigit, err := calculateCheckDigit(vin)
	if err != nil {
		return nil, fmt.Errorf("check digit calculation error: %w", err)
	}

	// Validate check digit
	checkDigitValid, err := validateCheckDigit(vin)
	if err != nil {
		return nil, fmt.Errorf("check digit validation error: %w", err)
	}

	// Decode WMI (World Manufacturer Identifier) - positions 1-3
	wmiCode := vin[0:3]
	wmi2 := ""
	// Check if this is an LVM (Low Volume Manufacturer) - third char is '9'
	if len(wmiCode) >= 3 && wmiCode[2] == '9' && len(vin) >= 14 {
		// Extract WMI2 from positions 12-14 (0-indexed: 11-13)
		wmi2 = vin[11:14]
	}

	wmiRecord := Lookup(wmiCode, wmi2)

	// Decode VDS (Vehicle Descriptor Section) - positions 4-9
	vds := vin[3:9]

	// Decode VIS (Vehicle Identifier Section) - positions 10-17
	vis := vin[9:17]

	// Create proto message
	result := &vinv1.Vin{}
	result.SetValue(vin)
	result.SetWmi(wmiCode)
	result.SetVds(vds)
	result.SetVis(vis)

	// Set check digit fields
	checkDigit := string(vin[8])
	result.SetCheckDigit(checkDigit)
	result.SetCalculatedCheckDigit(calculatedCheckDigit)
	result.SetCheckDigitValid(checkDigitValid)

	// Set WMI lookup fields
	if wmiRecord != nil {
		result.SetManufacturer(wmiRecord.M)
		if wmiRecord.C != vinv1.Country_COUNTRY_UNSPECIFIED {
			result.SetCountry(wmiRecord.C)
		}
		if wmiRecord.R != vinv1.Region_REGION_UNSPECIFIED {
			result.SetRegion(wmiRecord.R)
		}
		if wmiRecord.B != vinv1.Brand_BRAND_UNSPECIFIED {
			result.SetBrand(wmiRecord.B)
		}
	}

	return result, nil
}

// validateVINCharacters checks each character of the VIN.
// Valid characters are digits 0-9 and letters A-H, J-N, P, R-Z.
func validateVINCharacters(vin string) error {
	for i, r := range vin {
		if !isValidVINChar(r) {
			return fmt.Errorf("invalid VIN: invalid character '%c' at position %d (I, O, Q are not allowed)", r, i+1)
		}
	}
	return nil
}

func isValidVINChar(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		if r == 'I' || r == 'O' || r == 'Q' {
			return false
		}
		return true
	}
	return false
}
