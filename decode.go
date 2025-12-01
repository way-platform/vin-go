package vin

import (
	"fmt"

	"github.com/way-platform/vin-go/internal/checkdigit"
	"github.com/way-platform/vin-go/internal/iso3779"
	"github.com/way-platform/vin-go/internal/oem/fordvin"
	"github.com/way-platform/vin-go/internal/oem/mercedesvin"
	"github.com/way-platform/vin-go/internal/oem/scaniavin"
	"github.com/way-platform/vin-go/internal/oem/volkswagenvin"
	"github.com/way-platform/vin-go/internal/oem/volvotrucksvin"
	"github.com/way-platform/vin-go/internal/wmi"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// Decode validates and decodes a Vehicle Identification Number (VIN).
func Decode(vin string) (*vinv1.Vin, error) {
	if len(vin) != 17 {
		return nil, fmt.Errorf("invalid VIN length: expected 17 characters, got %d", len(vin))
	}
	for i, r := range vin {
		// I, O, Q are not allowed in VINs but they appear anyway in some VINs.
		if '0' <= r && r <= '9' || 'A' <= r && r <= 'Z' {
			continue
		}
		return nil, fmt.Errorf("invalid VIN: invalid character '%c' at position %d", r, i+1)
	}
	var output vinv1.Vin
	output.SetValue(vin)
	output.SetWmi(vin[0:3])
	output.SetVds(vin[3:9])
	output.SetVis(vin[9:17])
	if region, ok := wmi.ResolveRegion(output.GetWmi()); ok {
		output.SetRegion(region)
	}
	if country, ok := wmi.ResolveCountry(output.GetWmi()); ok {
		output.SetCountry(country)
	}
	if year, ok := iso3779.Year(vin[9]); ok {
		output.SetYear(int32(year))
	}
	calculatedCheckDigit, err := checkdigit.Calculate(vin)
	if err != nil {
		return nil, fmt.Errorf("check digit calculation error: %w", err)
	}
	output.SetCalculatedCheckDigit(calculatedCheckDigit)
	checkDigitValid, err := checkdigit.Validate(vin)
	if err != nil {
		output.SetCheckDigitValid(false)
	}
	output.SetCheckDigitValid(checkDigitValid)
	if output.GetWmi()[2] == '9' {
		// Low Volume Manufacturer - extract WMI2 from positions 12-14 (0-indexed: 11-13)
		wmi2 := vin[11:14]
		if m, found := LookupLowVolumeManufacturer(output.GetWmi(), wmi2); found {
			output.SetManufacturer(m)
		}
	} else {
		if m, found := LookupManufacturer(output.GetWmi()); found {
			output.SetManufacturer(m)
		}
	}
	vehicleDecoders := []func(string) (*vinv1.Vehicle, bool){
		mercedesvin.DecodeVehicle,
		scaniavin.DecodeVehicle,
		volkswagenvin.DecodeVehicle,
		volvotrucksvin.DecodeVehicle,
		fordvin.DecodeVehicle,
	}
	for _, vehicleDecoder := range vehicleDecoders {
		if vehicle, ok := vehicleDecoder(vin); ok {
			output.SetVehicle(vehicle)
			break
		}
	}
	if len(output.GetManufacturer().GetBrands()) == 1 && !output.GetVehicle().HasBrand() {
		if !output.HasVehicle() {
			output.SetVehicle(&vinv1.Vehicle{})
		}
		output.GetVehicle().SetBrand(output.GetManufacturer().GetBrands()[0])
	}
	if len(output.GetManufacturer().GetVehicleTypes()) == 1 && !output.GetVehicle().HasType() {
		if !output.HasVehicle() {
			output.SetVehicle(&vinv1.Vehicle{})
		}
		output.GetVehicle().SetType(output.GetManufacturer().GetVehicleTypes()[0])
	}
	return &output, nil
}
