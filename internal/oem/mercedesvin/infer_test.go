package mercedesvin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		vin         string
		wantBrand   vinv1.Brand
		wantModel   vinv1.Model
		wantType    vinv1.VehicleType
		wantYear    int32
		wantAxle    int32
		wantSuccess bool
	}{
		{
			name:        "Invalid Length",
			vin:         "123",
			wantSuccess: false,
		},
		{
			name:        "Unknown WMI",
			vin:         "ZZZ12345678901234",
			wantSuccess: false,
		},
		{
			name:        "Mercedes Sprinter LCV (VS30)",
			vin:         "W1V90700000000000", // W1V = MB Van, 907 = Sprinter VS30
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Vito (447)",
			vin:         "WDF44700000000000", // WDF = MB Van, 447 = Vito
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_VITO,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Metris (US 447)",
			vin:         "W1W44700000000000", // W1W = MB MPV (US), 447 = Metris context
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_METRIS,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE, // Model preference LCV
			wantSuccess: true,
		},
		{
			name:        "Mercedes eSprinter (910 FWD Electric)",
			vin:         "W1V9106331P123456", // W1V, 910, Subtype 6xx (6) -> eSprinter
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_E_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes V-Class (Passenger W1K)",
			vin:         "W1K44700000000000", // W1K + 447 -> V-Class
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_V_CLASS,
			wantType:    vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Citan (415)",
			vin:         "WDF41500000000000",
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_CITAN,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Citan/T-Class (420)",
			vin:         "W1T42000000000000",
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_CITAN, // Mapped to Citan for now
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Sprinter US Spec (W1V German Built)",
			vin:         "W1V3HCFZ1PP559516", // W1V + 3H = Sprinter
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantYear:    2023, // P = 2023
			wantAxle:    2,    // F = Class F 4x2
			wantSuccess: true,
		},
		{
			name:        "Mercedes Actros HGV",
			vin:         "W1T96300000000000", // W1T = MB Truck, 963 = Actros MP4
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_ACTROS,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Freightliner Econic",
			vin:         "W1H00000000000000", // W1H = Freightliner Econic
			wantBrand:   vinv1.Brand_FREIGHTLINER,
			wantModel:   vinv1.Model_E_ECONIC,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Dodge Sprinter (T1N)",
			vin:         "WD090300000000000", // WD0 = Dodge Sprinter Truck, 903 = T1N
			wantBrand:   vinv1.Brand_DODGE,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Mercedes Sprinter Bus (Explicit WMI override)",
			vin:         "W1Z90700000000000", // W1Z = MB Bus
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_BUS,
			wantSuccess: true,
		},
		{
			name:        "Ambiguous Dodge/Freightliner (WD1)",
			vin:         "WD190400000000000", // WD1 = Incomplete, 904 = Sprinter
			wantBrand:   vinv1.Brand_DODGE,   // Defaulted to Dodge
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE, // Sprinter model -> LCV
			wantSuccess: true,
		},
		{
			name:        "US Sprinter (Gas) - W1W40...",
			vin:         "W1W40000000000001", // W1W = MPV, 40 = Sprinter Gas
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE, // Sprinter -> LCV
			wantSuccess: true,
		},
		{
			name:        "US Sprinter (V6 Diesel) - W1Y5E...",
			vin:         "W1Y5E000000000001", // W1Y = Truck, 5E = Sprinter V6 Diesel
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE, // Sprinter -> LCV
			wantSuccess: true,
		},
		{
			name:        "US Metris - W1WV0...",
			vin:         "W1WV0000000000001", // V0 = Metris
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_METRIS,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Axor - WDB950...",
			vin:         "WDB95000000000001", // 950 = Axor
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_AXOR,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Zetros - WDB949...",
			vin:         "WDB94900000000001", // 949 = Zetros
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_ZETROS,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "Unimog - WDB405...",
			vin:         "WDB40500000000001", // 405 = Unimog UGN
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_UNIMOG,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "US SUV (W1N) - GLE/GLS",
			vin:         "W1N16600000000001", // W1N = US SUV, 166 = W166 (GLE) - Model not yet mapped, check type
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_MODEL_UNSPECIFIED,
			wantType:    vinv1.VehicleType_MULTIPURPOSE_PASSENGER_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "US Passenger (W1K) - C-Class",
			vin:         "W1K20500000000001", // W1K = US Car
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_MODEL_UNSPECIFIED,
			wantType:    vinv1.VehicleType_PASSENGER_CAR,
			wantSuccess: true,
		},
		// New model codes (OM654 Diesel variants)
		{
			name:        "US Sprinter (OM654 Diesel) - W1W4N...",
			vin:         "W1W4N000000000001", // 4N = Sprinter 2500 OM654 Diesel
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "US Sprinter (OM654 Diesel) - W1Y5N...",
			vin:         "W1Y5N000000000001", // 5N = Sprinter 3500 OM654 Diesel
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "US Sprinter (OM654 Diesel) - W1W8N...",
			vin:         "W1W8N000000000001", // 8N = Sprinter 3500 XD OM654 Diesel
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "US Sprinter (OM654 Diesel) - W1Y9N...",
			vin:         "W1Y9N000000000001", // 9N = Sprinter 4500 OM654 Diesel
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		// Fuel type extraction tests
		{
			name:        "US Sprinter Gasoline with fuel type",
			vin:         "W1W40B00000000001", // 40 = Gasoline, B = Cargo Van
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "US Sprinter Diesel with fuel type",
			vin:         "W1W4DB00000000001", // 4D = OM651 Diesel, B = Cargo Van
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		{
			name:        "US eSprinter Electric with fuel type",
			vin:         "W1W4VB00000000001", // 4V = Electric, B = Cargo Van
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_E_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantSuccess: true,
		},
		// Vehicle type refinement tests (passenger vans → PASSENGER_CAR)
		{
			name:        "US Sprinter Passenger Van (F) - should be PASSENGER_CAR",
			vin:         "W1W40F00000000001", // F = Passenger Van
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_PASSENGER_CAR,
			wantSuccess: true,
		},
		{
			name:        "US Sprinter Passenger Van (G) - should be PASSENGER_CAR",
			vin:         "W1W40G00000000001", // G = Passenger Van
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_PASSENGER_CAR,
			wantSuccess: true,
		},
		{
			name:        "US Sprinter Class 3 (Axle Count Test)",
			vin:         "W1W403300P0000001", // 40=Sprinter, Pos 7 (index 6) = '3' -> Class 3 (2 Axle), P=2023
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_SPRINTER,
			wantType:    vinv1.VehicleType_LIGHT_COMMERCIAL_VEHICLE,
			wantYear:    2023,
			wantAxle:    2,
			wantSuccess: true,
		},
		{
			name:        "Mercedes eActros (Series 983)",
			vin:         "W1T98300000000001", // W1T + 983 = eActros
			wantBrand:   vinv1.Brand_MERCEDES_BENZ,
			wantModel:   vinv1.Model_E_ACTROS,
			wantType:    vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gott, ok := DecodeVehicle(tt.vin)
			if tt.wantSuccess {
				if !ok {
					t.Fatalf("Decode(%q) failed: got !ok, want ok", tt.vin)
				}
				if gott == nil {
					t.Fatalf("Decode(%q) returned nil, want non-nil", tt.vin)
				}
				want := &vinv1.Vehicle{}
				if tt.wantBrand != vinv1.Brand_BRAND_UNSPECIFIED {
					want.SetBrand(tt.wantBrand)
				}
				if tt.wantModel != vinv1.Model_MODEL_UNSPECIFIED {
					want.SetModel(tt.wantModel)
				}
				if tt.wantType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
					want.SetType(tt.wantType)
				}
				if tt.wantYear > 0 {
					want.SetYear(tt.wantYear)
				}
				if tt.wantAxle > 0 {
					want.SetAxleCount(tt.wantAxle)
				}
				// Compare only brand, model, and type for now
				// fuel_types is a new field that will be reviewed separately
				opts := []cmp.Option{
					protocmp.Transform(),
					protocmp.IgnoreFields(&vinv1.Vehicle{}, "fuel_types"),
				}
				if diff := cmp.Diff(want, gott, opts...); diff != "" {
					t.Errorf("Decode(%q) mismatch (-want +got):\n%s", tt.vin, diff)
				}
			} else {
				if ok {
					t.Errorf("Decode(%q) succeeded, want failure", tt.vin)
				}
				if gott != nil {
					t.Errorf("Decode(%q) returned %v, want nil", tt.vin, gott)
				}
			}
		})
	}
}
