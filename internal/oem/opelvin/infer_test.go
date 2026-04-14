package opelvin

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
		wantVehicle *vinv1.Vehicle
		wantOk      bool
	}{
		{
			name:   "Non-Opel VIN",
			vin:    "WF0RXXTA200000000",
			wantOk: false,
		},
		{
			name: "Opel Vivaro W0V F7D600",
			vin:  "W0VF7D60000000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_OPEL),
				Model:       new(vinv1.Model_VIVARO),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Opel Vivaro W0V F7D601",
			vin:  "W0VF7D60100000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_OPEL),
				Model:       new(vinv1.Model_VIVARO),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Opel Vivaro W0V F7D607",
			vin:  "W0VF7D60700000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_OPEL),
				Model:       new(vinv1.Model_VIVARO),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Generic Opel W0L leaves model and type unset",
			vin:  "W0LABCDEF00000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_OPEL),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "VXE Vivaro (Hordain, Pos 4 V)",
			vin:  "VXEV1ZKXZ00000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_OPEL),
				Model:       new(vinv1.Model_VIVARO),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "VXE Combo (Hordain, Pos 4 E)",
			vin:  "VXEE1ABCD00000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_OPEL),
				Model:       new(vinv1.Model_COMBO),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "VXE generic LCV (Hordain, unknown platform)",
			vin:  "VXEX1ABCD00000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_OPEL),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVehicle, gotOk := DecodeVehicle(tt.vin)
			if gotOk != tt.wantOk {
				t.Fatalf("DecodeVehicle() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
			}
			if diff := cmp.Diff(tt.wantVehicle, gotVehicle, protocmp.Transform()); diff != "" {
				t.Fatalf("DecodeVehicle() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
