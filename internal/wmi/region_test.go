package wmi

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestResolveRegion(t *testing.T) {
	tests := []struct {
		wmi    string
		region vinv1.Region
		ok     bool
	}{
		{wmi: "1A9", region: vinv1.Region_NORTH_AMERICA, ok: true},
		{wmi: "W1V", region: vinv1.Region_EUROPE, ok: true},
		{wmi: "E11", region: vinv1.Region_REGION_UNSPECIFIED, ok: false},
	}
	for _, test := range tests {
		region, ok := ResolveRegion(test.wmi)
		if region != test.region {
			t.Errorf("ResolveRegion(%s) = %v, want %v", test.wmi, region, test.region)
		}
		if ok != test.ok {
			t.Errorf("ResolveRegion(%s) = %v, want %v", test.wmi, ok, test.ok)
		}
	}
}
