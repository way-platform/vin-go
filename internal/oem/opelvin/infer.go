package opelvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// DecodeVehicle infers Opel vehicle identity from VINs that cannot be trusted to
// resolve cleanly from the merged WMI manufacturer data alone.
func DecodeVehicle(vin string) (*vinv1.Vehicle, bool) {
	if len(vin) != 17 {
		return nil, false
	}

	wmi := vin[0:3]
	if wmi != "W0L" && wmi != "W0V" {
		return nil, false
	}

	builder := vinv1.Vehicle_builder{
		Brand:       new(vinv1.Brand_OPEL),
		DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
	}

	if wmi == "W0V" && vin[3:5] == "F7" {
		builder.Model = new(vinv1.Model_VIVARO)
		builder.Type = new(vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE)
	}

	return builder.Build(), true
}
