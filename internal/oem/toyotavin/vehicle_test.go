package toyotavin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
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
			want: func() *vinv1.Vehicle {
				v := &vinv1.Vehicle{}
				v.SetBrand(vinv1.Brand_TOYOTA)
				v.SetModel(vinv1.Model_YARIS_CROSS)
				v.SetType(vinv1.VehicleType_PASSENGER_CAR)
				v.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_GASOLINE})
				v.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})
				return v
			}(),
			ok: true,
		},
		{
			name: "Corolla Commercial Hybrid (UK)",
			vin:  "SB1Z93BE100000000",
			want: func() *vinv1.Vehicle {
				v := &vinv1.Vehicle{}
				v.SetBrand(vinv1.Brand_TOYOTA)
				v.SetModel(vinv1.Model_COROLLA)
				v.SetType(vinv1.VehicleType_PASSENGER_CAR)
				v.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_GASOLINE})
				v.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})
				return v
			}(),
			ok: true,
		},
		{
			name: "ProAce City (Stellantis Platform)",
			vin:  "YARKBAC3000000000",
			want: func() *vinv1.Vehicle {
				v := &vinv1.Vehicle{}
				v.SetBrand(vinv1.Brand_TOYOTA)
				v.SetModel(vinv1.Model_PROACE_CITY)
				v.SetType(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE)
				v.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_DIESEL})
				v.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})
				return v
			}(),
			ok: true,
		},
		{
			name: "ProAce Electric (Stellantis Platform)",
			vin:  "YARV1ZKXZ00000000",
			want: func() *vinv1.Vehicle {
				v := &vinv1.Vehicle{}
				v.SetBrand(vinv1.Brand_TOYOTA)
				v.SetModel(vinv1.Model_PROACE)
				v.SetType(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE)
				v.SetFuelTypes([]vinv1.FuelType{vinv1.FuelType_ELECTRIC})
				v.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})
				return v
			}(),
			ok: true,
		},
		{
			name: "Unknown VIN",
			vin:  "12345678901234567",
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
