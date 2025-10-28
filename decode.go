package vin

import (
	_ "embed"
	"fmt"
)

// Specification represents a fully decoded Vehicle Identification Number
type Specification struct {
	// Wmi is the World Manufacturer Identifier (positions 1-3)
	WMI WMI `json:"wmi"`
	// Vds is the Vehicle Descriptor Section (positions 4-9)
	VDS VDS `json:"vds"`
	// Vis is the Vehicle Identifier Section (positions 10-17)
	VIS VIS `json:"vis"`
}

// WMI represents the World Manufacturer Identifier
type WMI struct {
	// Code is the 3-character WMI code (positions 1-3)
	Code string `json:"code"`
	// Manufacturer is the vehicle manufacturer name
	Manufacturer string `json:"manufacturer"`
	// Country is the country of manufacture
	Country string `json:"country"`
	// GeographicArea is the geographic region
	GeographicArea GeographicArea `json:"geographicArea"`
}

// VDS represents the Vehicle Descriptor Section
type VDS struct {
	// ManufacturerSpecific contains positions 4-8 (manufacturer-defined codes)
	ManufacturerSpecific string `json:"manufacturerSpecific"`
	// CheckDigit is the check digit value (position 9)
	CheckDigit string `json:"checkDigit"`
}

// VIS represents the Vehicle Identifier Section
type VIS struct {
	// ModelYear is the vehicle's model year decoded from position 10
	ModelYear int `json:"modelYear"`
	// PlantCode is the manufacturing plant code (position 11)
	PlantCode string `json:"plantCode"`
	// SerialNumber is the vehicle's serial number (positions 12-17)
	SerialNumber string `json:"serialNumber"`
}

// Decode validates and decodes a Vehicle Identification Number (VIN).
// It returns a DecodedVIN struct containing all decoded information,
// or an error if the VIN is invalid.
func Decode(vin string) (*Specification, error) {
	if len(vin) != 17 {
		return nil, fmt.Errorf("invalid VIN length: expected 17 characters, got %d", len(vin))
	}
	if err := validateVINCharacters(vin); err != nil {
		return nil, err
	}
	checkDigitValid, err := validateCheckDigit(vin)
	if err != nil {
		return nil, fmt.Errorf("check digit validation error: %w", err)
	}
	if !checkDigitValid {
		return nil, fmt.Errorf("invalid VIN: check digit validation failed")
	}
	// Decode WMI (World Manufacturer Identifier) - positions 1-3
	wmi := decodeWMI(vin)
	// Decode VDS (Vehicle Descriptor Section) - positions 4-9
	vds := decodeVDS(vin)
	// Decode VIS (Vehicle Identifier Section) - positions 10-17
	vis := decodeVIS(vin)
	return &Specification{
		WMI: wmi,
		VDS: vds,
		VIS: vis,
	}, nil
}

// decodeWMI decodes the World Manufacturer Identifier section (positions 1-3)
func decodeWMI(vin string) WMI {
	wmiCode := vin[0:3]
	wmiEntry := lookupWMI(wmiCode)
	geoArea := lookupGeographicArea(rune(vin[0]))
	return WMI{
		Code:           wmiCode,
		Manufacturer:   wmiEntry.Manufacturer,
		Country:        wmiEntry.Country,
		GeographicArea: geoArea,
	}
}

// decodeVDS decodes the Vehicle Descriptor Section (positions 4-9)
func decodeVDS(vin string) VDS {
	manufacturerSpecific := vin[3:8]
	checkDigit := string(vin[8])
	return VDS{
		ManufacturerSpecific: manufacturerSpecific,
		CheckDigit:           checkDigit,
	}
}

// decodeVIS decodes the Vehicle Identifier Section (positions 10-17)
func decodeVIS(vin string) VIS {
	modelYearCode := string(vin[9])
	var modelYear int
	if yr, ok := lookupModelYear(modelYearCode); ok {
		modelYear = yr
	}
	plantCode := string(vin[10])
	serialNumber := vin[11:17]
	return VIS{
		ModelYear:    modelYear,
		PlantCode:    plantCode,
		SerialNumber: serialNumber,
	}
}

// validateVINCharacters checks each character of the VIN.
// // Valid characters are digits 0-9 and letters A-H, J-N, P, R-Z.
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
