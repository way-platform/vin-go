package vin

import (
	"sort"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// traficomEnergySpec holds the energy capacity for a specific fuel type configuration.
type traficomEnergySpec struct {
	FuelType           vinv1.FuelType
	FuelTankCapacityL  float64 // Fuel tank capacity in liters (0 = unknown)
	BatteryCapacityKwh float64 // Battery capacity in kWh (0 = unknown)
}

// traficomConfig holds a keyed energy spec for sorted lookup.
// Key layout: brand<<32 | model<<16 | fuel_type
type traficomConfig struct {
	Key uint64 // See key layout above
	traficomEnergySpec
}

// brandModelMask zeroes the fuel_type bits (lower 16) to match on brand+model only.
const brandModelMask = 0xFFFFFFFF0000

// lookupTraficomEnergySpec returns the best-effort energy capacity for a brand+model pair.
// Selection logic:
//  1. If fuelTypes contains a match, use it (first match wins).
//  2. If fuelTypes is empty and only one entry exists, use it.
//  3. If fuelTypes is empty and exactly one non-electric entry exists, use it.
//     Electric vehicles are always explicitly identified (separate model or annotation),
//     so an unknown fuel type is never electric.
//  4. Otherwise, return false (ambiguous).
func lookupTraficomEnergySpec(brand vinv1.Brand, model vinv1.Model, fuelTypes []vinv1.FuelType) (*traficomEnergySpec, bool) {
	prefix := uint64(brand)<<32 | uint64(model)<<16
	i := sort.Search(len(traficomData), func(i int) bool {
		return traficomData[i].Key >= prefix
	})
	var results []traficomEnergySpec
	for i < len(traficomData) && traficomData[i].Key&brandModelMask == prefix {
		results = append(results, traficomData[i].traficomEnergySpec)
		i++
	}
	for _, ft := range fuelTypes {
		for _, r := range results {
			if r.FuelType == ft {
				return &r, true
			}
		}
	}
	if len(fuelTypes) == 0 && len(results) == 1 {
		return &results[0], true
	}
	if len(fuelTypes) == 0 {
		var nonElectric []traficomEnergySpec
		for _, r := range results {
			if r.FuelType != vinv1.FuelType_ELECTRIC {
				nonElectric = append(nonElectric, r)
			}
		}
		if len(nonElectric) == 1 {
			return &nonElectric[0], true
		}
	}
	return nil, false
}
