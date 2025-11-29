package vpic

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// ResolveCountry resolves the country of a vPIC country ID.
func ResolveCountry(countryID int32) (vinv1.Country, bool) {
	values := vinv1.Country_COUNTRY_UNSPECIFIED.Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		ext := proto.GetExtension(value.Options(), vinv1.E_VpicCountryId)
		if ids, ok := ext.([]int32); ok {
			for _, id := range ids {
				if id == countryID {
					return vinv1.Country(value.Number()), true
				}
			}
		} else if id, ok := ext.(int32); ok {
			if id == countryID {
				return vinv1.Country(value.Number()), true
			}
		}
	}
	return vinv1.Country_COUNTRY_UNSPECIFIED, false
}
