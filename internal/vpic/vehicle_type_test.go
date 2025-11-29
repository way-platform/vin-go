package vpic

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestResolveVehicleType(t *testing.T) {
	tests := []struct {
		vehicleTypeID int32
		vehicleType   vinv1.VehicleType
		ok            bool
	}{
		{vehicleTypeID: 2, vehicleType: vinv1.VehicleType_PASSENGER_CAR, ok: true},
		{vehicleTypeID: 6, vehicleType: vinv1.VehicleType_TRAILER, ok: true},
		{vehicleTypeID: 1, vehicleType: vinv1.VehicleType_MOTORCYCLE, ok: true},
		{vehicleTypeID: 3, vehicleType: vinv1.VehicleType_TRUCK, ok: true},
		{vehicleTypeID: 9, vehicleType: vinv1.VehicleType_LOW_SPEED_VEHICLE, ok: true},
		{vehicleTypeID: 0, vehicleType: vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED, ok: false},
		{vehicleTypeID: 999, vehicleType: vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED, ok: false},
	}
	for _, test := range tests {
		vehicleType, ok := ResolveVehicleType(test.vehicleTypeID)
		if vehicleType != test.vehicleType {
			t.Errorf("ResolveVehicleType(%d) = %v, want %v", test.vehicleTypeID, vehicleType, test.vehicleType)
		}
		if ok != test.ok {
			t.Errorf("ResolveVehicleType(%d) = %v, want %v", test.vehicleTypeID, ok, test.ok)
		}
	}
}
