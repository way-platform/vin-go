package fordvin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func newVehicle(brand vinv1.Brand, model vinv1.Model, vehicleType vinv1.VehicleType) *vinv1.Vehicle {
	v := &vinv1.Vehicle{}
	v.SetBrand(brand)
	if model != vinv1.Model_MODEL_UNSPECIFIED {
		v.SetModel(model)
	}
	if vehicleType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
		v.SetType(vehicleType)
	}
	return v
}

func TestDecodeVehicle(t *testing.T) {
	tests := []struct {
		name        string
		vin         string
		wantVehicle *vinv1.Vehicle
		wantOk      bool
	}{
		{
			name:        "Non-Ford VIN",
			vin:         "WDB2120021A123456", // Mercedes VIN
			wantVehicle: nil,
			wantOk:      false,
		},
		{
			name:        "Ford NM0 Transit (Kocaeli, Pos 9 X)",
			vin:         "NM0TXXTTXTNM12345", // NM0, Body T, Source T, Plant T, Model X
			wantVehicle: newVehicle(vinv1.Brand_FORD, vinv1.Model_TRANSIT, vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
			wantOk:      true,
		},
		{
			name:        "Ford NM0 Transit Custom (Kocaeli, Pos 9 F)",
			vin:         "NM0TXXTTFTNM12345", // NM0, Body T, Source T, Plant T, Model F
			wantVehicle: newVehicle(vinv1.Brand_FORD, vinv1.Model_TRANSIT_CUSTOM, vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
			wantOk:      true,
		},
		{
			name:        "Ford WF0 Transit Connect (Valencia, Body W, Pos 9 undefined)",
			vin:         "WF0WXXGPUWNM12345", // WF0, Body W, Plant P, Model U (undefined/connect)
			wantVehicle: newVehicle(vinv1.Brand_FORD, vinv1.Model_TRANSIT_CONNECT, vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
			wantOk:      true,
		},
		{
			name:        "Ford WF0 Passenger Car (Valencia, Body F, Pos 9 B)",
			vin:         "WF0FXXGPBFNM12345", // WF0, Body F, Plant P, Model B (Mondeo)
			wantVehicle: newVehicle(vinv1.Brand_FORD, vinv1.Model_MODEL_UNSPECIFIED, vinv1.VehicleType_PASSENGER_CAR),
			wantOk:      true,
		},
		{
			name: "Ford WF0 Transit Courier (Craiova, Body W)",
			// Note: Transit Courier is not a distinct model in proto. Mapping to Transit Connect as closest.
			vin:         "WF0WXXGRWWNM12345", // WF0, Body W, Plant R, Model W?
			wantVehicle: newVehicle(vinv1.Brand_FORD, vinv1.Model_TRANSIT_COURIER, vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
			wantOk:      true,
		},
		{
			name:        "Ford WF0 Transit Courier (Plant A, Model C)",
			vin:         "WF0WXXTACLLD12345", // WF0, Body W, Plant A, Model C, Year L (2020), Month L
			wantVehicle: newVehicle(vinv1.Brand_FORD, vinv1.Model_TRANSIT_COURIER, vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
			wantOk:      true,
		},
		{
			name:        "Ford Ranger (WMI AFA)",
			vin:         "AFAXXXXXXXXXX0001", // AFA, 17 chars
			wantVehicle: newVehicle(vinv1.Brand_FORD, vinv1.Model_RANGER, vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE),
			wantOk:      true,
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
