package mercedesvin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestDecode(t *testing.T) {
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
			name:        "Mercedes Sprinter LCV (VS30)",
			vin:         "W1V90700000000000", // W1V = MB Van, 907 = Sprinter VS30
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Vito (447)",
			vin:         "WDF44700000000000", // WDF = MB Van, 447 = Vito
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_VITO,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Metris (US 447)",
			vin:         "W1W44700000000000", // W1W = MB MPV (US), 447 = Metris context
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_METRIS,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE, // Model preference LCV
			wantSuccess: true,
		},
		{
			name:        "Mercedes Actros HGV",
			vin:         "W1T96300000000000", // W1T = MB Truck, 963 = Actros MP4
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_ACTROS,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Freightliner Econic",
			vin:         "W1H00000000000000", // W1H = Freightliner Econic
			wantBrand:   vinv1.Brand_FREIGHTLINER,
			wantModel:   vinv1.Model_E_ECONIC,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Dodge Sprinter (T1N)",
			vin:         "WD090300000000000", // WD0 = Dodge Sprinter Truck, 903 = T1N
			wantBrand:   vinv1.Brand_DODGE,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Sprinter Bus (Explicit WMI override)",
			vin:         "W1Z90700000000000", // W1Z = MB Bus
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_BUS,
			wantSuccess: true,
		},
		{
			name:        "Ambiguous Dodge/Freightliner (WD1)",
			vin:         "WD190400000000000", // WD1 = Incomplete, 904 = Sprinter
			wantBrand:   vinv1.Brand_DODGE,   // Defaulted to Dodge
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE, // Sprinter model -> LCV
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gott, ok := DecodeVehicle(tt.vin)
			if tt.wantSuccess {
				if !ok {
					t.Fatalf("Decode(%q) failed: got !ok, want ok", tt.vin)
				}
				if gott == nil {
					t.Fatalf("Decode(%q) returned nil, want non-nil", tt.vin)
				}
				want := &vinv1.Vehicle{}
				want.SetBrand(tt.wantBrand)
				want.SetModel(tt.wantModel)
				want.SetVehicleType(tt.wantType)
				if diff := cmp.Diff(want, gott, protocmp.Transform()); diff != "" {
					t.Errorf("Decode(%q) mismatch (-want +got):\n%s", tt.vin, diff)
				}
			} else {
				if ok {
					t.Errorf("Decode(%q) succeeded, want failure", tt.vin)
				}
				if gott != nil {
					t.Errorf("Decode(%q) returned %v, want nil", tt.vin, gott)
				}
			}
		})
	}
}
