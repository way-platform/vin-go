package vin

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestLookupTraficom(t *testing.T) {
	tests := []struct {
		name      string
		brand     vinv1.Brand
		model     vinv1.Model
		fuelTypes []vinv1.FuelType
		want      traficomEnergySpec
		wantOK    bool
	}{
		{
			name:   "VW Transporter single entry, no fuel type → picks it",
			brand:  vinv1.Brand_VOLKSWAGEN,
			model:  vinv1.Model_TRANSPORTER,
			want:   traficomEnergySpec{FuelType: vinv1.FuelType_DIESEL, FuelTankCapacityL: 70},
			wantOK: true,
		},
		{
			name:   "Ford Transit single entry, no fuel type → picks it",
			brand:  vinv1.Brand_FORD,
			model:  vinv1.Model_TRANSIT,
			want:   traficomEnergySpec{FuelType: vinv1.FuelType_DIESEL, FuelTankCapacityL: 70},
			wantOK: true,
		},
		{
			name:      "Mercedes Sprinter multiple entries, fuel type selects electric",
			brand:     vinv1.Brand_MERCEDES_BENZ,
			model:     vinv1.Model_SPRINTER,
			fuelTypes: []vinv1.FuelType{vinv1.FuelType_ELECTRIC},
			want:      traficomEnergySpec{FuelType: vinv1.FuelType_ELECTRIC, BatteryCapacityKwh: 113},
			wantOK:    true,
		},
		{
			name:   "Mercedes Vito no fuel type → falls back to sole non-electric entry",
			brand:  vinv1.Brand_MERCEDES_BENZ,
			model:  vinv1.Model_VITO,
			want:   traficomEnergySpec{FuelType: vinv1.FuelType_DIESEL, FuelTankCapacityL: 57},
			wantOK: true,
		},
		{
			name:   "Ford Transit Custom multiple non-electric entries, no fuel type → ambiguous",
			brand:  vinv1.Brand_FORD,
			model:  vinv1.Model_TRANSIT_CUSTOM,
			wantOK: false,
		},
		{
			name:   "unknown model → false",
			brand:  vinv1.Brand_VOLKSWAGEN,
			model:  vinv1.Model(9999),
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := lookupTraficomEnergySpec(tt.brand, tt.model, tt.fuelTypes)
			if ok != tt.wantOK {
				t.Fatalf("lookupTraficomEnergySpec() ok = %v, want %v", ok, tt.wantOK)
			}
			if ok && *got != tt.want {
				t.Errorf("lookupTraficomEnergySpec() = %+v, want %+v", *got, tt.want)
			}
		})
	}
}

func TestDecodeEnrichesCapacity(t *testing.T) {
	// W1V9106331P383288 is a Mercedes E-Sprinter — the mercedesvin decoder sets
	// fuelTypes: [ELECTRIC], which allows the Traficom selection to pick electric.
	decoded, err := Decode("W1V9106331P383288")
	if err != nil {
		t.Fatalf("Decode() error: %v", err)
	}
	v := decoded.GetVehicle()
	if v == nil {
		t.Fatal("expected vehicle, got nil")
	}
	if v.GetBatteryCapacityKwh() != 113 {
		t.Errorf("BatteryCapacityKwh = %v, want 113", v.GetBatteryCapacityKwh())
	}

	dataSources := v.GetDataSources()
	var hasTraficom bool
	for _, ds := range dataSources {
		if ds == vinv1.DataSource_TRAFICOM {
			hasTraficom = true
			break
		}
	}
	if !hasTraficom {
		t.Errorf("expected TRAFICOM in dataSources, got %v", dataSources)
	}
}

func TestDecodeEnrichesWithoutFuelType(t *testing.T) {
	// NM0GEXTTS0A000000 is a Ford Transit — the fordvin decoder does NOT set fuelTypes,
	// but Transit has only one Traficom entry (diesel), so it gets selected.
	decoded, err := Decode("NM0GEXTTS0A000000")
	if err != nil {
		t.Fatalf("Decode() error: %v", err)
	}
	v := decoded.GetVehicle()
	if v == nil {
		t.Fatal("expected vehicle, got nil")
	}
	if v.GetBrand() != vinv1.Brand_FORD {
		t.Fatalf("expected FORD, got %v", v.GetBrand())
	}
	if v.GetModel() != vinv1.Model_TRANSIT {
		t.Fatalf("expected TRANSIT, got %v", v.GetModel())
	}
	if v.GetFuelTankCapacityL() != 70 {
		t.Errorf("FuelTankCapacityL = %v, want 70", v.GetFuelTankCapacityL())
	}
}
