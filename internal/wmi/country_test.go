package wmi

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestResolveCountry(t *testing.T) {
	tests := []struct {
		wmi     string
		country vinv1.Country
		ok      bool
	}{
		{wmi: "1A9", country: vinv1.Country_UNITED_STATES, ok: true},
		{wmi: "W1V", country: vinv1.Country_GERMANY, ok: true},
		{wmi: "YF9", country: vinv1.Country_FINLAND, ok: true},
		{wmi: "2AA", country: vinv1.Country_CANADA, ok: true},
		{wmi: "7CB", country: vinv1.Country_NEW_ZEALAND, ok: true},
		{wmi: "foo", country: vinv1.Country_COUNTRY_UNSPECIFIED, ok: false},
		{wmi: "6XX", country: vinv1.Country_COUNTRY_UNSPECIFIED, ok: false},
	}
	for _, test := range tests {
		country, ok := ResolveCountry(test.wmi)
		if country != test.country {
			t.Errorf("ResolveCountry(%s) = %v, want %v", test.wmi, country, test.country)
		}
		if ok != test.ok {
			t.Errorf("ResolveCountry(%s) = %v, want %v", test.wmi, ok, test.ok)
		}
	}
}
