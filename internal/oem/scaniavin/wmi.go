package scaniavin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// Swedish WMIs
// Reference: docs/deep-research/scania.md Section 2.2
var swedishWMIs = map[string]bool{
	"YS2": true, // Scania CV AB (Södertälje) - Primary code for trucks and commercial vehicles
	"YS4": true, // Scania Bus (Katrineholm) - Historically linked to bus production (pre-2002)
}

// Brazilian WMIs
// Reference: docs/deep-research/scania.md Section 2.2
var brazilianWMIs = map[string]bool{
	"9BS": true, // Scania Latin America Ltda. (São Bernardo do Campo)
}

// Mexican WMIs
// Reference: docs/deep-research/scania.md Section 2.2
var mexicanWMIs = map[string]bool{
	"3AX": true, // Scania Mexico - Trucks
	"3BE": true, // Scania Mexico - Buses
}

// Netherlands WMIs
// Reference: docs/deep-research/scania.md Section 2.2
var netherlandsWMIs = map[string]bool{
	"XLE": true, // Scania Production (Zwolle/Meppel) - Note: XLE is a prefix, full codes may be XLER, XLEP
}

// Polish WMIs
var polishWMIs = map[string]bool{
	"SZA": true, // Scania Production Slupsk - Bus factory
}

// French WMIs
var frenchWMIs = map[string]bool{
	"VLU": true, // Scania Production France - Truck factory
}

// Bus-specific WMIs
// Reference: docs/deep-research/scania.md Section 2.2
var busWMIs = map[string]bool{
	"YS4": true, // Scania Bus (Katrineholm)
	"3BE": true, // Scania Mexico - Buses
	"SZA": true, // Scania Production Slupsk - Bus factory
}

// getBrandFromWMI determines the brand from the WMI.
// Returns the brand and true if the WMI is recognized, false otherwise.
// Reference: docs/deep-research/scania.md Section 2
func getBrandFromWMI(wmi string) (vinv1.Brand, bool) {
	if swedishWMIs[wmi] ||
		brazilianWMIs[wmi] ||
		mexicanWMIs[wmi] ||
		netherlandsWMIs[wmi] ||
		polishWMIs[wmi] ||
		frenchWMIs[wmi] {
		return vinv1.Brand_SCANIA, true
	}

	// Handle XLE prefix codes (XLER, XLEP)
	if len(wmi) >= 3 && wmi[:3] == "XLE" {
		return vinv1.Brand_SCANIA, true
	}

	return vinv1.Brand_BRAND_UNSPECIFIED, false
}

// isBusWMI checks if the WMI indicates a Scania Bus.
func isBusWMI(wmi string) bool {
	return busWMIs[wmi]
}
