package cargobullvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers a vehicle's identity from its VIN.
// Returns false if the VIN is not recognized as a Schmitz Cargobull VIN.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]

	switch wmi {
	case "WSM", "WSK":
		v := &vinv1.Vehicle{}
		v.SetBrand(vinv1.Brand_SCHMITZ_CARGOBULL)
		v.SetType(vinv1.VehicleType_TRAILER)
		v.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})
		return v, true
	default:
		return nil, false
	}
}
