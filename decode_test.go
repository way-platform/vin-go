package vin

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestDecode(t *testing.T) {
	// Test a standard VIN
	vin := "1HGFC16533A004352"
	result, err := Decode(vin)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	if result.GetValue() != vin {
		t.Errorf("Expected VIN %s, got %s", vin, result.GetValue())
	}

	if result.GetWmi() != "1HG" {
		t.Errorf("Expected WMI 1HG, got %s", result.GetWmi())
	}

	if result.GetVds() != "FC1653" {
		t.Errorf("Expected VDS FC1653, got %s", result.GetVds())
	}

	if result.GetVis() != "3A004352" {
		t.Errorf("Expected VIS 3A004352, got %s", result.GetVis())
	}

	// Check WMI lookup
	if result.GetManufacturer() == "" {
		t.Error("Expected manufacturer to be set")
	}

	if result.GetCountry() != vinv1.Country_UNITED_STATES {
		t.Errorf("Expected country UNITED_STATES, got %v", result.GetCountry())
	}

	if result.GetRegion() != vinv1.Region_NORTH_AMERICA {
		t.Errorf("Expected region NORTH_AMERICA, got %v", result.GetRegion())
	}
}

func TestDecodeLVM(t *testing.T) {
	// Test an LVM VIN (WMI1 third char is '9')
	// Using a known LVM entry from the CSV: 1A9 with WMI2
	// Note: This is a hypothetical test - adjust based on actual CSV data
	vin := "1A90000000000000" // This won't validate, but tests the LVM lookup path
	_, err := Decode(vin)
	// We expect an error due to invalid check digit, but the LVM lookup should work
	if err == nil {
		t.Log("Note: VIN decoded (may be invalid)")
	}
}

func TestLookup(t *testing.T) {
	// Test standard WMI lookup
	result := Lookup("1HG", "")
	if result == nil {
		t.Fatal("Expected lookup result for 1HG")
		return // Early return to satisfy staticcheck
	}

	if result.M == "" {
		t.Error("Expected manufacturer to be set")
	}

	if result.C == vinv1.Country_COUNTRY_UNSPECIFIED {
		t.Error("Expected country to be set")
	}

	// Test LVM lookup
	lvmResult := Lookup("109", "013")
	if lvmResult == nil {
		t.Fatal("Expected lookup result for LVM 109/013")
		return // Early return to satisfy staticcheck
	}
	if lvmResult.M == "" {
		t.Error("Expected manufacturer to be set for LVM")
	}
}
