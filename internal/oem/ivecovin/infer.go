package ivecovin

import (
	"strconv"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}
	wmi := vin[0:3]

	// WMI Check
	// ZCF: Iveco Italy (Main)
	// ZGA: Iveco Bus (Main)
	// VF5: Iveco France (Unic)
	// WJM: Iveco Magirus (Germany)
	// 8AT: Iveco Argentina
	// 93Z: Iveco Brazil
	// 6T9: Iveco Australia
	// AAV: Iveco South Africa
	// XUE: Iveco Russia (AMT)
	// VNE, VFE: Iveco Bus France

	isIveco := false
	isBus := false

	switch wmi {
	case "ZCF", "VF5", "WJM", "8AT", "93Z", "6T9", "AAV", "XUE":
		isIveco = true
	case "ZGA", "VNE", "VFE": // Bus
		isIveco = true
		isBus = true
	}

	if !isIveco {
		return nil, false
	}

	output := &vinv1.Vehicle{}
	output.SetBrand(vinv1.Brand_IVECO)
	output.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})

	if isBus {
		output.SetType(vinv1.VehicleType_BUS)
		return output, true
	}

	// Default assumption for Iveco commercial vehicles is Diesel and 2 axles
	output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_DIESEL})
	output.SetAxleCount(2)

	// VDS (Positions 4-9) - 0-indexed 3-8
	pos4 := vin[3] // Family

	// Model Logic
	var model vinv1.Model
	var vehicleType vinv1.VehicleType

	switch pos4 {
	case 'C', 'D': // Daily
		model = vinv1.Model_DAILY
		vehicleType = vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE
	case 'E', 'A': // Eurocargo (A is usually older or lighter, E is standard)
		model = vinv1.Model_EUROCARGO
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	case 'S': // Stralis / S-Way
		// Defaulting to Stralis as it covers a longer period, S-Way is relatively new (2019+).
		// Without year or specific generation indicators, Stralis is the safer broad category,
		// though S-Way is the current model.
		model = vinv1.Model_STRALIS
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	case 'T': // Trakker
		model = vinv1.Model_TRAKKER
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	case 'M': // Heavy (Trakker/Stralis/S-Way)
		vehicleType = vinv1.VehicleType_HEAVY_GOODS_VEHICLE
	}

	// Daily Specific Logic
	if model == vinv1.Model_DAILY {
		// Fuel / Chassis (Pos 5)
		pos5 := vin[4]
		switch pos5 {
		case 'N':
			output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_COMPRESSED_NATURAL_GAS})
		case 'E':
			// Only confirm Electric if Engine Code (Pos 8) is 'S' (Electric Motor).
			// Otherwise, 'E' at Pos 5 might denote a specific Diesel chassis variant (e.g. legacy or specific cab).
			if vin[7] == 'S' {
				output.SetModel(vinv1.Model_E_DAILY)
				output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_ELECTRIC})
			}
		}

		// Pos 8 check for CNG/Electric (Engine Code)
		pos8 := vin[7]
		switch pos8 {
		case 'G': // Natural Power (CNG)
			output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_COMPRESSED_NATURAL_GAS})
		case 'S': // Electric
			output.SetModel(vinv1.Model_E_DAILY)
			output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_ELECTRIC})
		}

		// Weight (Pos 6-7)
		// digits correspond to tons * 10 (e.g. 35 -> 3.5t, 70 -> 7.0t)
		// We could potentially set VehicleType to HEAVY_GOODS_VEHICLE if > 35,
		// but Daily is canonically an LCV (N1/N2 mix). Proto defines Daily as LCV.
		// We verify the digits are numeric just in case.
		if _, err := strconv.Atoi(vin[5:7]); err == nil {
			_ = err // TODO: Consider re-classifying as HGV if heavier than 3.5T
			// Logic valid
		}
	} else if vehicleType == vinv1.VehicleType_HEAVY_GOODS_VEHICLE {
		// Heavy Logic (Eurocargo, Stralis, Trakker)
		// Pos 8: Engine/Propulsion
		pos8 := vin[7]
		switch pos8 {
		case 'G': // Natural Power
			output.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_COMPRESSED_NATURAL_GAS})
		}
	}

	if model != vinv1.Model_MODEL_UNSPECIFIED {
		output.SetModel(model)
	}
	if vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
		output.SetType(vehicleType)
	}

	return output, true
}
