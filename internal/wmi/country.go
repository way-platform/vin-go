package wmi

import (
	"strings"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// ResolveCountry resolves the country of a WMI.
func ResolveCountry(wmi string) (vinv1.Country, bool) {
	values := vinv1.Country_COUNTRY_UNSPECIFIED.Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		wmiPrefixes, ok := proto.GetExtension(value.Options(), vinv1.E_WmiPrefix).([]string)
		if !ok {
			continue
		}
		for _, wmiPrefix := range wmiPrefixes {
			if strings.HasPrefix(wmi, wmiPrefix) {
				return vinv1.Country(value.Number()), true
			}
		}
	}
	return vinv1.Country_COUNTRY_UNSPECIFIED, false
}
