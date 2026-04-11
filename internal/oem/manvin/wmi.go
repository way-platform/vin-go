package manvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// MAN Truck & Bus WMIs.
// Reference: docs/deep-research/man.md Section "Geopolitical and Corporate Identification"
var knownWMIs = map[string]bool{
	"WMA": true, // MAN Truck & Bus (Germany) - primary WMI
}

// getBrandFromWMI determines the brand from the WMI.
// Returns the brand and true if the WMI is recognized, false otherwise.
func getBrandFromWMI(wmi string) (vinv1.Brand, bool) {
	if knownWMIs[wmi] {
		return vinv1.Brand_MAN, true
	}
	return vinv1.Brand_BRAND_UNSPECIFIED, false
}
