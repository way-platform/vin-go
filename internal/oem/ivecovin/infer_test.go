package ivecovin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestDecodeVehicle(t *testing.T) {
	tests := []struct {
		name        string
		vin         string
		wantBrand   vinv1.Brand
		wantModel   vinv1.Model
		wantType    vinv1.VehicleType
		wantFuel    []vinv1.FuelType
		wantAxles   int32
		wantSuccess bool
	}{
		{
			name:        "Iveco Eurocargo 120E28",
			vin:         "ZCFA71EF800000000",
			wantBrand:   vinv1.Brand_IVECO,
			wantModel:   vinv1.Model_EUROCARGO,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuel:    []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxles:   2,
			wantSuccess: true,
		},
		{
			name:        "Iveco Daily 35S",
			vin:         "ZCFCR35A700000000",
			wantBrand:   vinv1.Brand_IVECO,
			wantModel:   vinv1.Model_DAILY,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantFuel:    []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxles:   2,
			wantSuccess: true,
		},
		{
			name:        "Iveco Daily 72C",
			vin:         "ZCFCS72A300000000",
			wantBrand:   vinv1.Brand_IVECO,
			wantModel:   vinv1.Model_DAILY,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantFuel:    []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxles:   2,
			wantSuccess: true,
		},
		{
			name:        "Iveco Daily Natural Power (CNG)",
			vin:         "ZCFCN70A800000000",
			wantBrand:   vinv1.Brand_IVECO,
			wantModel:   vinv1.Model_DAILY,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantFuel:    []vinv1.FuelType{vinv1.FuelType_COMPRESSED_NATURAL_GAS},
			wantAxles:   2,
			wantSuccess: true,
		},
		{
			name:        "Iveco Daily Diesel (Reported Bug)",
			vin:         "ZCFCE52C900000000",
			wantBrand:   vinv1.Brand_IVECO,
			wantModel:   vinv1.Model_DAILY,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantFuel:    []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxles:   2,
			wantSuccess: true,
		},
		{
			name:        "Non-Iveco",
			vin:         "1M8GDM9A_KP042788",
			wantSuccess: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := DecodeVehicle(tt.vin)
			if tt.wantSuccess {
				if !ok {
					t.Fatalf("DecodeVehicle(%q) failed: got !ok, want ok", tt.vin)
				}
				if got == nil {
					t.Fatalf("DecodeVehicle(%q) returned nil, want non-nil", tt.vin)
				}
				want := vinv1.Vehicle_builder{
					DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
				}
				if tt.wantBrand != vinv1.Brand_BRAND_UNSPECIFIED {
					want.Brand = new(tt.wantBrand)
				}
				if tt.wantModel != vinv1.Model_MODEL_UNSPECIFIED {
					want.Model = new(tt.wantModel)
				}
				if tt.wantType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
					want.Type = new(tt.wantType)
				}
				if len(tt.wantFuel) > 0 {
					want.FuelTypes = tt.wantFuel
				}
				if tt.wantAxles > 0 {
					want.AxleCount = new(tt.wantAxles)
				}
				if diff := cmp.Diff(want.Build(), got, protocmp.Transform()); diff != "" {
					t.Errorf("DecodeVehicle(%q) mismatch (-want +got):\n%s", tt.vin, diff)
				}
			} else {
				if ok {
					t.Errorf("DecodeVehicle(%q) succeeded, want failure", tt.vin)
				}
			}
		})
	}
}
