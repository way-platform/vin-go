package vin

import (
	"fmt"

	"github.com/way-platform/vin-go/internal/checkdigit"
	"github.com/way-platform/vin-go/internal/iso3779"
	"github.com/way-platform/vin-go/internal/oem/cargobullvin"
	"github.com/way-platform/vin-go/internal/oem/fordvin"
	"github.com/way-platform/vin-go/internal/oem/ivecovin"
	"github.com/way-platform/vin-go/internal/oem/mercedesvin"
	"github.com/way-platform/vin-go/internal/oem/opelvin"
	"github.com/way-platform/vin-go/internal/oem/renaulttrucksvin"
	"github.com/way-platform/vin-go/internal/oem/scaniavin"
	"github.com/way-platform/vin-go/internal/oem/toyotavin"
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
	output := vinv1.Vin_builder{
		Value: new(vin),
		Wmi:   new(vin[0:3]),
		Vds:   new(vin[3:9]),
		Vis:   new(vin[9:17]),
	}
	if region, ok := wmi.ResolveRegion(*output.Wmi); ok {
		output.Region = &region
	}
	if country, ok := wmi.ResolveCountry(*output.Wmi); ok {
		output.Country = &country
	}
	if year, ok := iso3779.Year(vin[9]); ok {
		yearValue := int32(year)
		output.Year = &yearValue
	}
	calculatedCheckDigit, err := checkdigit.Calculate(vin)
	if err != nil {
		return nil, fmt.Errorf("check digit calculation error: %w", err)
	}
	output.CalculatedCheckDigit = &calculatedCheckDigit
	checkDigitValid, err := checkdigit.Validate(vin)
	if err != nil {
		checkDigitValid = false
	}
	output.CheckDigitValid = &checkDigitValid
	if (*output.Wmi)[2] == '9' {
		// Low Volume Manufacturer - extract WMI2 from positions 12-14 (0-indexed: 11-13)
		wmi2 := vin[11:14]
		if m, found := LookupLowVolumeManufacturer(*output.Wmi, wmi2); found {
			output.Manufacturer = m
		}
	} else {
		if m, found := LookupManufacturer(*output.Wmi); found {
			output.Manufacturer = m
		}
	}
	vehicleDecoders := []func(string) (*vinv1.Vehicle, bool){
		cargobullvin.DecodeVehicle,
		ivecovin.DecodeVehicle,
		mercedesvin.DecodeVehicle,
		opelvin.DecodeVehicle,
		renaulttrucksvin.DecodeVehicle,
		scaniavin.DecodeVehicle,
		toyotavin.DecodeVehicle,
		volkswagenvin.DecodeVehicle,
		volvotrucksvin.DecodeVehicle,
		fordvin.DecodeVehicle,
		inferVehicleFromManufacturer(output.Manufacturer), // fallback
	}
	for _, vehicleDecoder := range vehicleDecoders {
		if vehicle, ok := vehicleDecoder(vin); ok {
			output.Vehicle = vehicle
			break
		}
	}
	// Suppress incorrect year for EU Mercedes VINs.
	// EU Mercedes uses VIN position 10 for steering orientation (1=LHD, 2=RHD),
	// not model year. The Mercedes OEM decoder only sets Vehicle.Year for US-spec
	// VINs, so an absent Vehicle.Year indicates the position-10 year is unreliable.
	if v := output.Vehicle; v != nil && v.HasBrand() &&
		v.GetBrand() == vinv1.Brand_MERCEDES_BENZ && !v.HasYear() {
		output.Year = nil
	}
	return output.Build(), nil
}

func inferVehicleFromManufacturer(input *vinv1.Manufacturer) func(string) (*vinv1.Vehicle, bool) {
	return func(vin string) (*vinv1.Vehicle, bool) {
		output := vinv1.Vehicle_builder{
			DataSources: input.GetDataSources(),
		}
		if len(input.GetBrands()) == 1 {
			output.Brand = new(input.GetBrands()[0])
		}
		if len(input.GetVehicleTypes()) == 1 {
			output.Type = new(input.GetVehicleTypes()[0])
		}
		built := output.Build()
		return built, built.HasBrand() || built.HasType()
	}
}
