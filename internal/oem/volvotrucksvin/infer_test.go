package volvotrucksvin

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
			name:        "Invalid Length",
			vin:         "123",
			wantSuccess: false,
		},
		{
			name:        "Unknown WMI",
			vin:         "12345678901234567",
			wantSuccess: false,
		},
		{
			name:        "Volvo VNL (North America)",
			vin:         "4V4NC9EH5JN123456", // 4V4=Volvo, N=VNL
			wantBrand:   vinv1.Brand_VOLVO_TRUCKS,
			wantModel:   vinv1.Model_VNL,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Volvo VNM (North America Legacy)",
			vin:         "4V4MC9EH5JN123456", // M=VNM
			wantBrand:   vinv1.Brand_VOLVO_TRUCKS,
			wantModel:   vinv1.Model_VNM,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Volvo VNR (North America)",
			vin:         "4V4RC9EH5JN123456", // R=VNR
			wantBrand:   vinv1.Brand_VOLVO_TRUCKS,
			wantModel:   vinv1.Model_VNR,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Volvo VHD (North America)",
			vin:         "4V4KC9EH5JN123456", // K=VHD
			wantBrand:   vinv1.Brand_VOLVO_TRUCKS,
			wantModel:   vinv1.Model_VHD,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Volvo VNX (North America)",
			vin:         "4V4XC9EH5JN123456", // X=VNX
			wantBrand:   vinv1.Brand_VOLVO_TRUCKS,
			wantModel:   vinv1.Model_VNX,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Volvo VNR Electric (North America)",
			vin:         "4V4WB9EH5JN123456", // W=VNR Electric
			wantBrand:   vinv1.Brand_VOLVO_TRUCKS,
			wantModel:   vinv1.Model_VNR_ELECTRIC,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Volvo Trucks Global (Sweden)",
			vin:         "YV2A1234567890123", // YV2=Volvo Sweden
			wantBrand:   vinv1.Brand_VOLVO_TRUCKS,
			wantModel:   vinv1.Model_MODEL_UNSPECIFIED, // Not determining global model yet
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Volvo Trucks Global (Belgium)",
			vin:         "YB3A1234567890123", // YB3=Volvo Belgium
			wantBrand:   vinv1.Brand_VOLVO_TRUCKS,
			wantModel:   vinv1.Model_MODEL_UNSPECIFIED,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Volvo Bus (Global)",
			vin:         "YV3A1234567890123", // YV3=Volvo Bus
			wantBrand:   vinv1.Brand_VOLVO_BUSES,
			wantModel:   vinv1.Model_MODEL_UNSPECIFIED,
			wantType:    vinv1.VehicleType_BUS,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gott, ok := DecodeVehicle(tt.vin)
			if tt.wantSuccess {
				if !ok {
					t.Errorf("DecodeVehicle() failed")
					return
				}

				// Construct expected object
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

				if diff := cmp.Diff(want, gott, protocmp.Transform()); diff != "" {
					t.Errorf("DecodeVehicle() mismatch (-want +got):\n%s", diff)
				}
			} else {
				if ok {
					t.Errorf("DecodeVehicle() succeeded unexpectedly")
				}
			}
		})
	}
}
