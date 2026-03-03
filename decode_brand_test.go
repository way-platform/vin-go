package vin

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestDecode_InferBrandFromWMIMapping(t *testing.T) {
	tests := []struct {
		name  string
		vin   string
		brand vinv1.Brand
	}{
		{
			name:  "Piako trailer",
			vin:   "YKB0442XX00000000",
			brand: vinv1.Brand_PIAKO,
		},
		{
			name:  "Ekeri trailer",
			vin:   "YF2T32SBG00000000",
			brand: vinv1.Brand_EKERI,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := Decode(tt.vin)
			if err != nil {
				t.Fatalf("Decode(%q) error: %v", tt.vin, err)
			}

			if len(decoded.GetManufacturer().GetBrands()) != 1 {
				t.Fatalf("expected exactly one manufacturer brand, got %v", decoded.GetManufacturer().GetBrands())
			}
			if got := decoded.GetManufacturer().GetBrands()[0]; got != tt.brand {
				t.Fatalf("manufacturer brand = %v, want %v", got, tt.brand)
			}

			if got := decoded.GetVehicle().GetBrand(); got != tt.brand {
				t.Fatalf("vehicle brand = %v, want %v", got, tt.brand)
			}
			if got := decoded.GetVehicle().GetType(); got != vinv1.VehicleType_TRAILER {
				t.Fatalf("vehicle type = %v, want %v", got, vinv1.VehicleType_TRAILER)
			}
		})
	}
}
