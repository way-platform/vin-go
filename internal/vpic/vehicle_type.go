package vpic

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// ResolveVehicleType resolves the vehicle type of a vPIC vehicle type ID.
func ResolveVehicleType(vehicleTypeID int32) (vinv1.VehicleType, bool) {
	values := vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED.Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		ext := proto.GetExtension(value.Options(), vinv1.E_VpicVehicleTypeId)
		if ids, ok := ext.([]int32); ok {
			for _, id := range ids {
				if id == vehicleTypeID {
					return vinv1.VehicleType(value.Number()), true
				}
			}
		} else if id, ok := ext.(int32); ok {
			if id == vehicleTypeID {
				return vinv1.VehicleType(value.Number()), true
			}
		}
	}
	return vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED, false
}
