package cargobullvin

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
		wantType    vinv1.VehicleType
		wantSuccess bool
	}{
		{
			name:        "WSM Schmitz Cargobull trailer",
			vin:         "WSM00000000000000",
			wantBrand:   vinv1.Brand_SCHMITZ_CARGOBULL,
			wantType:    vinv1.VehicleType_TRAILER,
			wantSuccess: true,
		},
		{
			name:        "WSK Schmitz Cargobull Gotha trailer",
			vin:         "WSK00000000000000",
			wantBrand:   vinv1.Brand_SCHMITZ_CARGOBULL,
			wantType:    vinv1.VehicleType_TRAILER,
			wantSuccess: true,
		},
		{
			name:        "Unknown WMI",
			vin:         "ZZZ00000000000000",
			wantSuccess: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := DecodeVehicle(tt.vin)
			if tt.wantSuccess {
				if !ok {
					t.Fatalf("DecodeVehicle(%q) failed: got !ok, want ok", tt.vin)
				}
				want := &vinv1.Vehicle{}
				want.SetBrand(tt.wantBrand)
				want.SetType(tt.wantType)
				want.SetDataSources([]vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH})
				if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
					t.Errorf("DecodeVehicle(%q) mismatch (-want +got):\n%s", tt.vin, diff)
				}
			} else {
				if ok {
					t.Errorf("DecodeVehicle(%q) succeeded, want failure", tt.vin)
				}
			}
		})
	}
}
