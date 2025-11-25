package kba

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestResolveBrand(t *testing.T) {
	tests := []struct {
		manufacturerID int32
		brand          vinv1.Brand
		ok             bool
	}{
		{manufacturerID: 600, brand: vinv1.Brand_VOLKSWAGEN, ok: true},
		{manufacturerID: 603, brand: vinv1.Brand_VOLKSWAGEN, ok: true},
	}
	for _, test := range tests {
		brand, ok := ResolveBrand(test.manufacturerID)
		if brand != test.brand {
			t.Errorf("ResolveBrand(%d) = %v, want %v", test.manufacturerID, brand, test.brand)
		}
		if ok != test.ok {
			t.Errorf("ResolveBrand(%d) = %v, want %v", test.manufacturerID, ok, test.ok)
		}
	}
}
