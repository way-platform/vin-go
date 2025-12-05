package renaulttrucksvin

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
		wantSuccess bool
	}{
		{
			name:        "Renault Trucks T (VF611)",
			vin:         "VF611G12345678901", // VF6 + 11 = T, G = Tractor
			wantBrand:   vinv1.Brand_RENAULT_TRUCKS,
			wantModel:   vinv1.Model_T,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Trucks K (VF634)",
			vin:         "VF634K12345678901", // VF6 + 34 = K
			wantBrand:   vinv1.Brand_RENAULT_TRUCKS,
			wantModel:   vinv1.Model_K,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Trucks C (VF624)",
			vin:         "VF624G12345678901", // VF6 + 24 = C
			wantBrand:   vinv1.Brand_RENAULT_TRUCKS,
			wantModel:   vinv1.Model_C,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Trucks D (VF621)",
			vin:         "VF621A12345678901", // VF6 + 21 = D
			wantBrand:   vinv1.Brand_RENAULT_TRUCKS,
			wantModel:   vinv1.Model_D,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Trucks Magnum (VF617)",
			vin:         "VF617G12345678901", // VF6 + 17 = Magnum
			wantBrand:   vinv1.Brand_RENAULT_TRUCKS,
			wantModel:   vinv1.Model_MAGNUM,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Master (VN1 MA)",
			vin:         "VN1MA123456789012", // VN1 + MA = Master
			wantBrand:   vinv1.Brand_RENAULT,
			wantModel:   vinv1.Model_MASTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Trafic (VF1 FL)",
			vin:         "VF1FL123456789012", // VF1 + FL = Trafic II
			wantBrand:   vinv1.Brand_RENAULT,
			wantModel:   vinv1.Model_TRAFIC,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Master HD (VF6 VF)",
			vin:         "VF6VFA12345678901", // VF6 + VF = Master HD
			wantBrand:   vinv1.Brand_RENAULT_TRUCKS,
			wantModel:   vinv1.Model_MASTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Master HD (VF6 MF)",
			vin:         "VF6MF000364077973", // VF6 + MF = Master
			wantBrand:   vinv1.Brand_RENAULT_TRUCKS,
			wantModel:   vinv1.Model_MASTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Renault Maxity (VS...)",
			vin:         "VSWA1234567890123", // VS...
			wantBrand:   vinv1.Brand_RENAULT_TRUCKS,
			wantModel:   vinv1.Model_MAXITY,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gott, ok := DecodeVehicle(tt.vin)
			if tt.wantSuccess {
				if !ok {
					t.Fatalf("DecodeVehicle(%q) failed: got !ok, want ok", tt.vin)
				}
				if gott == nil {
					t.Fatalf("DecodeVehicle(%q) returned nil, want non-nil", tt.vin)
				}
				want := &vinv1.Vehicle{}
				if tt.wantBrand != vinv1.Brand_BRAND_UNSPECIFIED {
					want.SetBrand(tt.wantBrand)
				}
				if tt.wantModel != vinv1.Model_MODEL_UNSPECIFIED {
					want.SetModel(tt.wantModel)
				}
				if tt.wantType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
					want.SetType(tt.wantType)
				}
				want.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})
				if diff := cmp.Diff(want, gott, protocmp.Transform()); diff != "" {
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
