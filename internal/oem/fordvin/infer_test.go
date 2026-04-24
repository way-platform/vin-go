package fordvin

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
			name:        "Non-Ford VIN",
			vin:         "WDB21200200000000", // Mercedes VIN
			wantVehicle: nil,
			wantOk:      false,
		},
		{
			name: "Ford NM0 Transit (Kocaeli, Pos 9 X)",
			vin:  "NM0TXXTTX00000000", // NM0, Body T, Source T, Plant T, Model X
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford NM0 Transit Custom (Kocaeli, Pos 9 F)",
			vin:  "NM0TXXTTF00000000", // NM0, Body T, Source T, Plant T, Model F
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CUSTOM),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Connect (Valencia, Body W, Pos 9 undefined)",
			vin:  "WF0WXXGPU00000000", // WF0, Body W, Plant P, Model U (undefined/connect)
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CONNECT),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Passenger Car (Valencia, Body F, Pos 9 B)",
			vin:  "WF0FXXGPB00000000", // WF0, Body F, Plant P, Model B (Mondeo)
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Type:        new(vinv1.VehicleType_PASSENGER_CAR),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Courier (Craiova, Body W)",
			vin:  "WF0WXXGRW00000000", // WF0, Body W, Plant R, Model W?
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_COURIER),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Courier (Plant A, Model C)",
			vin:  "WF0WXXTAC00000000", // WF0, Body W, Plant A, Model C
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_COURIER),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford Ranger (WMI AFA)",
			vin:  "AFAAAAAAA00000000", // AFA, 17 chars
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_RANGER),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Custom TA override",
			vin:  "WF0RXXTA200000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CUSTOM),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Custom check digit 0",
			vin:  "WF0RXXTA000000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CUSTOM),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Custom check digit 1",
			vin:  "WF0RXXTA100000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CUSTOM),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Custom check digit 5",
			vin:  "WF0RXXTA500000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CUSTOM),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Custom BEV (Pos 8 Z)",
			vin:  "WF0RXXTZ000000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CUSTOM),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Tourneo Custom M1 (Pos 4 A)",
			vin:  "WF0AXXTA000000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CUSTOM),
				Type:        new(vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Transit Custom WPG override",
			vin:  "WF0RXXWPG00000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT_CUSTOM),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 full-size Transit (Kocaeli, legacy model code F)",
			vin:  "WF0FXXTTFF8000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Model:       new(vinv1.Model_TRANSIT),
				Type:        new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
		{
			name: "Ford WF0 Puma (Craiova, non-commercial body)",
			vin:  "WF02XXERK2N000000",
			wantVehicle: vinv1.Vehicle_builder{
				Brand:       new(vinv1.Brand_FORD),
				Type:        new(vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE),
				DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
			}.Build(),
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVehicle, gotOk := DecodeVehicle(tt.vin)
			if gotOk != tt.wantOk {
				t.Errorf("DecodeVehicle() gotOk = %v, wantOk %v", gotOk, tt.wantOk)
				return
			}
			if tt.wantOk {
				if diff := cmp.Diff(tt.wantVehicle, gotVehicle, protocmp.Transform()); diff != "" {
					t.Errorf("DecodeVehicle() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
