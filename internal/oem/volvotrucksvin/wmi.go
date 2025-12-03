package volvotrucksvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// North American WMIs (4V1-4V5)
// Reference: docs/deep-research/volvo-trucks.md Section 2.1
var northAmericanWMIs = map[string]bool{
	"4V1": true, // Historical: Volvo GM Heavy Truck Corporation (WC, WI, early VN)
	"4V2": true, // Historical: Volvo GM Heavy Truck Corporation
	"4V3": true, // Historical: Volvo GM Heavy Truck Corporation
	"4V4": true, // Complete Volvo commercial trucks (VNL, VNR, VNX)
	"4V5": true, // Incomplete vehicles / chassis-cabs, also VNR Electric
}

// Global WMIs (YV2, YB3, 9BV, YV3)
// Reference: docs/deep-research/volvo-trucks.md Section 2.2
var globalWMIs = map[string]bool{
	"YV2": true, // Volvo Trucks Sweden (FH, FM, FMX, FL, FE)
	"YB3": true, // Volvo Europa Truck NV Belgium
	"9BV": true, // Volvo do Brasil
}

// Bus WMIs
var busWMIs = map[string]bool{
	"YV3": true, // Volvo Buses
}

// getBrandFromWMI determines the brand from the WMI.
// Returns the brand and true if the WMI is recognized, false otherwise.
func getBrandFromWMI(wmi string) (vinv1.Brand, bool) {
	if northAmericanWMIs[wmi] || globalWMIs[wmi] {
		return vinv1.Brand_VOLVO_TRUCKS, true
	}
	if busWMIs[wmi] {
		return vinv1.Brand_VOLVO_BUSES, true
	}
	return vinv1.Brand_BRAND_UNSPECIFIED, false
}

// isNorthAmericanWMI checks if the WMI belongs to North American Volvo Trucks.
func isNorthAmericanWMI(wmi string) bool {
	return northAmericanWMIs[wmi]
}

// isGlobalWMI checks if the WMI belongs to Global/European Volvo Trucks.
func isGlobalWMI(wmi string) bool {
	return globalWMIs[wmi]
}

// isBusWMI checks if the WMI indicates a Volvo Bus.
func isBusWMI(wmi string) bool {
	return busWMIs[wmi]
}

