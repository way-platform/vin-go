package volkswagenvin

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
			name:        "Non-VW VIN",
			vin:         "WDB12345600000000",
			wantSuccess: false,
		},
		{
			name:        "Case A: ID. Buzz Passenger (WV2...EB)",
			vin:         "WV2ZZZEB800000000",
			wantBrand:   vinv1.Brand_VOLKSWAGEN,
			wantModel:   vinv1.Model_ID_BUZZ,
			wantType:    vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Case B: Crafter Gen 2 Panel Van (WV1...SY)",
			vin:         "WV1ZZZSYZ00000000",
			wantBrand:   vinv1.Brand_VOLKSWAGEN,
			wantModel:   vinv1.Model_CRAFTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Case C: T7 Multivan (WV2...ST)",
			vin:         "WV2ZZZST000000000",
			wantBrand:   vinv1.Brand_VOLKSWAGEN,
			wantModel:   vinv1.Model_MULTIVAN,
			wantType:    vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Transporter T6 (WV1...7H)",
			vin:         "WV1ZZZ7HZ00000000",
			wantBrand:   vinv1.Brand_VOLKSWAGEN,
			wantModel:   vinv1.Model_TRANSPORTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Caddy Mk 3 (WV1...2K)",
			vin:         "WV1ZZZ2KZ00000000",
			wantBrand:   vinv1.Brand_VOLKSWAGEN,
			wantModel:   vinv1.Model_CADDY,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Amarok (WV1...2H)",
			vin:         "WV1ZZZ2HZ00000000",
			wantBrand:   vinv1.Brand_VOLKSWAGEN,
			wantModel:   vinv1.Model_AMAROK,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Crafter Gen 1 (WV1...2E)",
			vin:         "WV1ZZZ2EZ00000000",
			wantBrand:   vinv1.Brand_VOLKSWAGEN,
			wantModel:   vinv1.Model_CRAFTER,
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
				if diff := cmp.Diff(want.Build(), gott, protocmp.Transform()); diff != "" {
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
