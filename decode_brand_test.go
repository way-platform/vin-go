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

func TestDecode_DeepResearchOverrides(t *testing.T) {
	tests := []struct {
		name      string
		vin       string
		wantBrand vinv1.Brand
		wantModel vinv1.Model
		wantType  vinv1.VehicleType
	}{
		{
			name:      "Ford Transit Custom override",
			vin:       "WF0RXXTA200000000",
			wantBrand: vinv1.Brand_FORD,
			wantModel: vinv1.Model_TRANSIT_CUSTOM,
			wantType:  vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
		},
		{
			name:      "Opel Vivaro override",
			vin:       "W0VF7D60000000000",
			wantBrand: vinv1.Brand_OPEL,
			wantModel: vinv1.Model_VIVARO,
			wantType:  vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
		},
		{
			name:      "Generic Opel leaves type unset",
			vin:       "W0LABCDEF00000000",
			wantBrand: vinv1.Brand_OPEL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := Decode(tt.vin)
			if err != nil {
				t.Fatalf("Decode(%q) error: %v", tt.vin, err)
			}

			if got := decoded.GetVehicle().GetBrand(); got != tt.wantBrand {
				t.Fatalf("vehicle brand = %v, want %v", got, tt.wantBrand)
			}
			if got := decoded.GetVehicle().GetModel(); got != tt.wantModel {
				t.Fatalf("vehicle model = %v, want %v", got, tt.wantModel)
			}
			if got := decoded.GetVehicle().GetType(); got != tt.wantType {
				t.Fatalf("vehicle type = %v, want %v", got, tt.wantType)
			}
		})
	}
}
