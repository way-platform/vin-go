package vpic

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestResolveBrand(t *testing.T) {
	tests := []struct {
		makeID int32
		brand  vinv1.Brand
		ok     bool
	}{
		{makeID: 482, brand: vinv1.Brand_VOLKSWAGEN, ok: true},
		{makeID: 13647, brand: vinv1.Brand_RENAULT, ok: true},
		{makeID: 0, brand: vinv1.Brand_BRAND_UNSPECIFIED, ok: false},
	}
	for _, test := range tests {
		brand, ok := ResolveBrand(test.makeID)
		if brand != test.brand {
			t.Errorf("ResolveBrand(%d) = %v, want %v", test.makeID, brand, test.brand)
		}
		if ok != test.ok {
			t.Errorf("ResolveBrand(%d) = %v, want %v", test.makeID, ok, test.ok)
		}
	}
}
