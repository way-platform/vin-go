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
		vpicCountryId, ok := proto.GetExtension(value.Options(), vinv1.E_VpicCountryId).(int32)
		if !ok || vpicCountryId == 0 {
			continue
		}
		if vpicCountryId == countryID {
			return vinv1.Country(value.Number()), true
		}
	}
	return vinv1.Country_COUNTRY_UNSPECIFIED, false
}
