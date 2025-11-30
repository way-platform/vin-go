package scaniavin

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
			vin:         "ZZZ12345678901234",
			wantSuccess: false,
		},
		{
			name:        "Scania R-Series (Sweden)",
			vin:         "YS2R4x20000000001", // YS2=Sweden, R=R-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_R_SERIES,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Scania S-Series (NTG Sweden)",
			vin:         "YS2S5x20000000001", // S=S-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_S_SERIES,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Scania P-Series (Sweden)",
			vin:         "YS2P6x20000000001", // P=P-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_P_SERIES,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Scania G-Series (Sweden)",
			vin:         "YS2G4x20000000001", // G=G-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_G_SERIES,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Scania L-Series (Sweden)",
			vin:         "YS2L4x20000000001", // L=L-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_L_SERIES,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Scania T-Series (Legacy)",
			vin:         "YS2T4x20000000001", // T=T-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_T_SERIES,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Scania Bus K-Series (Sweden Bus WMI)",
			vin:         "YS4K4x20000000001", // YS4=Bus, K=K-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_K_SERIES,
			wantType:    vinv1.VehicleType_BUS,
			wantSuccess: true,
		},
		{
			name:        "Scania Bus N-Series (Mexico Bus WMI)",			
			vin:         "3BEN4x20000000001", // 3BE=Mexico Bus, N=N-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_N_SERIES,
			wantType:    vinv1.VehicleType_BUS,
			wantSuccess: true,
		},
		{
			name:        "Scania Bus F-Series (Brazil)",
			vin:         "9BSF4x20000000001", // 9BS=Brazil, F=F-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_F_SERIES,
			wantType:    vinv1.VehicleType_BUS,
			wantSuccess: true,
		},
		{
			name:        "Scania Netherlands Truck",
			vin:         "XLER4x20000000001", // XLE=Netherlands, R=R-Series. Using XLE for WMI.
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_R_SERIES,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Scania Poland Bus (SZA WMI, K-Series)",
			vin:         "SZAK4x20000000001", // SZA=Poland (Bus factory), K=K-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_K_SERIES,
			wantType:    vinv1.VehicleType_BUS,
			wantSuccess: true,
		},
		{
			name:        "Scania Truck (VLU WMI, R-Series)",
			vin:         "VLUR4x20000000001", // VLU=France (Truck factory), R=R-Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantModel:   vinv1.Model_R_SERIES,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Known WMI, Unknown Series Code",
			vin:         "YS2Z4x20000000001", // YS2=Sweden, Z=Unknown Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE, // Default to HGV if WMI is truck-centric
			wantSuccess: true,
		},
		{
			name:        "Known Bus WMI, Unknown Series Code",
			vin:         "YS4Z4x20000000001", // YS4=Bus WMI, Z=Unknown Series
			wantBrand:   vinv1.Brand_SCANIA,
			wantType:    vinv1.VehicleType_BUS, // Default to BUS if WMI is bus-centric
			wantSuccess: true,
		},
		{
			name:        "No Data Identified",
			vin:         "AAA12345678901234", // Neither brand, model, nor type are set based on mapping
			wantSuccess: false,
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
													want.SetVehicleType(tt.wantType)
												}
												if diff := cmp.Diff(want, gott, protocmp.Transform()); diff != "" {
													t.Errorf("DecodeVehicle(%q) mismatch (-want +got):\n%s", tt.vin, diff)
												}
				
			} else {
				if ok {
					t.Errorf("DecodeVehicle(%q) succeeded, want failure", tt.vin)
				}
				if gott != nil {
					t.Errorf("DecodeVehicle(%q) returned %v, want nil", tt.vin, gott)
				}
			}
		})
	}
}
