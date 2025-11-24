package vpic

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestResolveCountry(t *testing.T) {
	tests := []struct {
		countryID int32
		country   vinv1.Country
		ok        bool
	}{
		{countryID: 6, country: vinv1.Country_UNITED_STATES, ok: true},
		{countryID: 1, country: vinv1.Country_CANADA, ok: true},
		{countryID: 49, country: vinv1.Country_COSTA_RICA, ok: true},
		{countryID: 23, country: vinv1.Country_SWEDEN, ok: true},
		{countryID: 25, country: vinv1.Country_FINLAND, ok: true},
		{countryID: 0, country: vinv1.Country_COUNTRY_UNSPECIFIED, ok: false},
		{countryID: 999, country: vinv1.Country_COUNTRY_UNSPECIFIED, ok: false},
	}
	for _, test := range tests {
		country, ok := ResolveCountry(test.countryID)
		if country != test.country {
			t.Errorf("ResolveCountry(%d) = %v, want %v", test.countryID, country, test.country)
		}
		if ok != test.ok {
			t.Errorf("ResolveCountry(%d) = %v, want %v", test.countryID, ok, test.ok)
		}
	}
}
