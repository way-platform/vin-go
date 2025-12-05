package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// Mercedes-Benz WMIs
var mercedesWMIs = map[string]bool{
	"VSA": true, "WDB": true, "WDF": true, "WDD": true, "WDC": true,
	"W1T": true, "W1V": true, "W1W": true, "W1X": true, "W1Y": true, "W1Z": true,
	"W1K": true, "W1N": true, "WMX": true, "4JG": true,
	"WD3": true, "WD4": true, "WDA": true, "WDZ": true, "1MB": true,
	"8AB": true, "8AC": true, "8BN": true, "8BR": true, "8BT": true, "8BU": true,
	"9BM": true, "NMB": true,
}

// Freightliner WMIs
var freightlinerWMIs = map[string]bool{
	"W1H": true, "WD6": true, "WD7": true, "WDR": true, "WDP": true, "WDY": true,
	"WCD": true, "W2W": true, "W2X": true, "W2Y": true, "W2Z": true,
}

// Dodge WMIs
var dodgeWMIs = map[string]bool{
	"WD0": true, "WD8": true, "WDW": true, "WDX": true,
	// Ambiguous (Dodge or Freightliner) - defaulted to Dodge
	"WD1": true, "WD2": true, "WD5": true,
}

// US Commercial WMIs (for 2-digit code decoding)
var usCommercialWMIs = map[string]bool{
	"W1V": true, "W1W": true, "W1Y": true, "W1X": true,
	"WD3": true, "WD4": true, "WDA": true, "WDZ": true,
	"WD0": true, "WD8": true, "WDW": true, "WDX": true,
	"WD1": true, "WD2": true, "WD5": true,
	"W2W": true, "W2X": true,
}

// Bus WMIs
var busWMIs = map[string]bool{
	"W1Z": true, "WDZ": true, "WDW": true, "WCD": true, "W2Z": true, "8BR": true,
}

// isMercedesWMI checks if the WMI belongs to Mercedes-Benz brand.
func isMercedesWMI(wmi string) bool {
	return mercedesWMIs[wmi]
}

// isFreightlinerWMI checks if the WMI belongs to Freightliner brand.
func isFreightlinerWMI(wmi string) bool {
	return freightlinerWMIs[wmi]
}

// isDodgeWMI checks if the WMI belongs to Dodge brand.
func isDodgeWMI(wmi string) bool {
	return dodgeWMIs[wmi]
}

// getBrandFromWMI determines the brand from the WMI.
// Returns the brand and true if the WMI is recognized, false otherwise.
func getBrandFromWMI(wmi string) (vinv1.Brand, bool) {
	if isMercedesWMI(wmi) {
		return vinv1.Brand_MERCEDES_BENZ, true
	}
	if isFreightlinerWMI(wmi) {
		return vinv1.Brand_FREIGHTLINER, true
	}
	if isDodgeWMI(wmi) {
		return vinv1.Brand_DODGE, true
	}
	return vinv1.Brand_BRAND_UNSPECIFIED, false
}

// isUSCommercialWMI checks if the WMI indicates a US commercial vehicle
// that uses 2-digit model codes.
func isUSCommercialWMI(wmi string) bool {
	return usCommercialWMIs[wmi]
}

// isBusWMI checks if the WMI explicitly indicates a bus.
func isBusWMI(wmi string) bool {
	return busWMIs[wmi]
}

// isUSVehicleWMI checks if the WMI indicates a US-market vehicle
// (commercial or passenger) that may follow US VIN decoding rules for year/axle.
func isUSVehicleWMI(wmi string) bool {
	return isUSCommercialWMI(wmi) || wmi == "W1K" || wmi == "W1N" || wmi == "WDD"
}
