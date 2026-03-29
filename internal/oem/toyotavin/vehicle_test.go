package toyotavin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestDecodeVehicle(t *testing.T) {
	tests := []struct {
		name string
		vin  string
		want *vinv1.Vehicle
		ok   bool
	}{
		{
			name: "Yaris Cross Hybrid (France)",
			vin:  "VNKKD3D3800000000",
			want: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_TOYOTA),
				Model:       new(vinv1.Model_YARIS_CROSS),
				Type:        new(vinv1.VehicleType_PASSENGER_CAR),
				FuelTypes:   []vinv1.FuelType{vinv1.FuelType_GASOLINE},
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			ok: true,
		},
		{
			name: "Corolla Commercial Hybrid (UK)",
			vin:  "SB1Z93BE100000000",
			want: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_TOYOTA),
				Model:       new(vinv1.Model_COROLLA),
				Type:        new(vinv1.VehicleType_PASSENGER_CAR),
				FuelTypes:   []vinv1.FuelType{vinv1.FuelType_GASOLINE},
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			ok: true,
		},
		{
			name: "ProAce City (Stellantis Platform)",
			vin:  "YARKBAC3000000000",
			want: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_TOYOTA),
				Model:       new(vinv1.Model_PROACE_CITY),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				FuelTypes:   []vinv1.FuelType{vinv1.FuelType_DIESEL},
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			ok: true,
		},
		{
			name: "ProAce Electric (Stellantis Platform)",
			vin:  "YARV1ZKXZ00000000",
			want: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_TOYOTA),
				Model:       new(vinv1.Model_PROACE),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				FuelTypes:   []vinv1.FuelType{vinv1.FuelType_ELECTRIC},
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			ok: true,
		},
		{
			name: "Unknown VIN",
			vin:  "12345678900000000",
			want: nil,
			ok:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := DecodeVehicle(tt.vin)
			if ok != tt.ok {
				t.Errorf("DecodeVehicle() ok = %v, want %v", ok, tt.ok)
				return
			}
			if !ok {
				return
			}
			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("DecodeVehicle() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
