package manvin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestDecodeVehicle(t *testing.T) {
	tests := []struct {
		name          string
		vin           string
		wantBrand     vinv1.Brand
		wantModel     vinv1.Model
		wantType      vinv1.VehicleType
		wantFuelTypes []vinv1.FuelType
		wantAxleCount int32
		wantYear      int32
		wantSuccess   bool
	}{
		{
			name:        "Invalid Length",
			vin:         "123",
			wantSuccess: false,
		},
		{
			name:        "Unknown WMI",
			vin:         "ZZZ12345600000000",
			wantSuccess: false,
		},
		{
			name:          "MAN TGE (06K, 2020)",
			vin:           "WMA06KZZ0LP000001", // WMA=MAN, 06K=TGE, ZZ=filler, L=2020, P=Poland
			wantBrand:     vinv1.Brand_MAN,
			wantModel:     vinv1.Model_TGE,
			wantType:      vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2020,
			wantSuccess:   true,
		},
		{
			name:          "MAN Heavy Truck (28S, 2022)",
			vin:           "WMA28SZZ0NP000001", // WMA=MAN, 28S=heavy truck, ZZ=filler, N=2022, P=Poland
			wantBrand:     vinv1.Brand_MAN,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantYear:      2022,
			wantSuccess:   true,
		},
		{
			name:        "MAN Unknown VDS",
			vin:         "WMAXXXZZ0NP000001", // WMA=MAN, XXX=unknown model code
			wantBrand:   vinv1.Brand_MAN,
			wantYear:    2022,
			wantSuccess: true,
		},
		{
			name:          "MAN TGE (06K, 2019)",
			vin:           "WMA06KZZ0KP000001", // WMA=MAN, 06K=TGE, ZZ=filler, K=2019, P=Poland
			wantBrand:     vinv1.Brand_MAN,
			wantModel:     vinv1.Model_TGE,
			wantType:      vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2019,
			wantSuccess:   true,
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
				if len(tt.wantFuelTypes) > 0 {
					want.FuelTypes = tt.wantFuelTypes
				}
				if tt.wantAxleCount > 0 {
					want.AxleCount = new(tt.wantAxleCount)
				}
				if tt.wantYear > 0 {
					want.ModelYear = new(tt.wantYear)
				}
				if diff := cmp.Diff(want.Build(), got, protocmp.Transform()); diff != "" {
					t.Errorf("DecodeVehicle(%q) mismatch (-want +got):\n%s", tt.vin, diff)
				}
			} else {
				if ok {
					t.Errorf("DecodeVehicle(%q) succeeded, want failure", tt.vin)
				}
				if got != nil {
					t.Errorf("DecodeVehicle(%q) returned %v, want nil", tt.vin, got)
				}
			}
		})
	}
}
