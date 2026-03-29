package volvotrucksvin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestDecodeVehicle(t *testing.T) {
	tests := []struct {
		name          string
		vin           string
		wantBrand     vinv1.Brand
		wantModel     vinv1.Model
		wantType      vinv1.VehicleType
		wantFuelTypes []vinv1.FuelType
		wantAxleCount int32
		wantYear      int32
		wantSuccess   bool
	}{
		{
			name:        "Invalid Length",
			vin:         "123",
			wantSuccess: false,
		},
		{
			name:        "Unknown WMI",
			vin:         "12345678900000000",
			wantSuccess: false,
		},
		// North American Tests
		{
			name:          "Volvo VNL (North America) - Diesel, 6x4, 2022",
			vin:           "4V4NC9EH5N0000000", // 4V4=Volvo, N=VNL, C=6x4, E=D13, H=425-474HP, N=2022
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNL,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2022,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNM (North America Legacy) - Diesel, 6x4",
			vin:           "4V4MC9EH5J0000000", // M=VNM, C=6x4, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNM,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR (North America) - Diesel, 6x4",
			vin:           "4V4RC9EH5J0000000", // R=VNR, C=6x4, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR - Diesel, 6x2",
			vin:           "4V4RB9EH5J0000000", // R=VNR, B=6x2, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR - Diesel, 4x2",
			vin:           "4V4R39EH5J0000000", // R=VNR, 3=4x2 Class 8, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VHD (North America) - Diesel, 6x4",
			vin:           "4V4KC9EH5J0000000", // K=VHD, C=6x4, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VHD,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNX (North America) - Diesel, 6x4",
			vin:           "4V4XC9EH5J0000000", // X=VNX, C=6x4, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNX,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR Electric (North America) - Electric, 6x2",
			vin:           "4V4WB9EH5J0000000", // W=VNR Electric, B=6x2, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR_ELECTRIC,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_ELECTRIC},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR Electric - Electric via Position 8",
			vin:           "4V4WB9EN5J0000000", // W=VNR Electric, B=6x2, N=Electric in Position 8, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR_ELECTRIC,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_ELECTRIC},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:      "Volvo VNR - CNG/LNG Engine",
			vin:       "4V4RC9VH5J0000000", // R=VNR, C=6x4, V=Cummins ISL G (CNG/LNG), J=2018
			wantBrand: vinv1.Brand_VOLVO_TRUCKS,
			wantModel: vinv1.Model_VNR,
			wantType:  vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{
				vinv1.FuelType_COMPRESSED_NATURAL_GAS,
				vinv1.FuelType_LIQUEFIED_NATURAL_GAS,
			},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR - Cummins Diesel (T)",
			vin:           "4V4RC9TH5J0000000", // R=VNR, C=6x4, T=Cummins X15, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR - Cummins L9 Diesel (S)",
			vin:           "4V4RC9SH5J0000000", // R=VNR, C=6x4, S=Cummins L9, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VHD - D11 Diesel (D)",
			vin:           "4V4KC9DH5J0000000", // K=VHD, C=6x4, D=D11, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VHD,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNL - D16 Diesel (K)",
			vin:           "4V4NC9KH5J0000000", // N=VNL, C=6x4, K=D16, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNL,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Incomplete Vehicle (4V5, not Electric)",
			vin:           "4V5KC9EH5J0000000", // 4V5=Incomplete, K=VHD, C=6x4, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VHD,
			wantType:      vinv1.VehicleType_INCOMPLETE_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR Electric with 4V5 WMI",
			vin:           "4V5WB9EH5J0000000", // 4V5=Incomplete, but W=VNR Electric (complete), J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR_ELECTRIC,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_ELECTRIC},
			wantAxleCount: 3,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR - Multi-axle (9)",
			vin:           "4V4R99EH5J0000000", // R=VNR, 9=Multi-axle, E=D13, J=2018
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 4,
			wantYear:      2018,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR - 2021 Model Year",
			vin:           "4V4RC9EH5M0000000", // M=2021
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2021,
			wantSuccess:   true,
		},
		{
			name:          "Volvo VNR - 2024 Model Year",
			vin:           "4V4RC9EH5R0000000", // R=2024
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_VNR,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2024,
			wantSuccess:   true,
		},
		// Global/European Tests
		{
			name:          "Volvo Trucks Global (Sweden) - FM Series (A)",
			vin:           "YV2AG30A550000000", // YV2=Volvo Sweden, A=FM, G30=D13K420(Diesel), Pos8=A=4x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FM,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global - FH Series (R)",
			vin:           "YV2RG30A550000000", // R=FH, G30=Diesel, Pos8=A=4x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FH,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global - FL Series (T)",
			vin:           "YV2TG30A550000000", // T=FL, G30=Diesel, Pos8=A=4x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FL,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global - FE Series (V)",
			vin:           "YV2VG30A550000000", // V=FE, G30=Diesel, Pos8=A=4x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FE,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global - FL 4x2 (A)",
			vin:           "YV2TG30A550000000", // T=FL, G30=Diesel, Pos8=A=4x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FL,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global - FH 6x4 (D)",
			vin:           "YV2RG30D550000000", // R=FH, G30=Diesel, Pos8=D=6x4, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FH,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global - FM 8x4 (G)",
			vin:           "YV2AG30G550000000", // A=FM, G30=Diesel, Pos8=G=8x4, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FM,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 4,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global - FL 6x2 (C)",
			vin:           "YV2TG30C550000000", // T=FL, G30=Diesel, Pos8=C=6x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FL,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 3,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global - Electric Engine Code",
			vin:           "YV2T0P0A550000000", // T=FL, 0P0=Electric (Pos5-7), Pos8=A=4x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FL,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_ELECTRIC},
			wantAxleCount: 2,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global (Belgium)",
			vin:           "YB3AG30A550000000", // YB3=Volvo Belgium, A=FM, G30=Diesel, Pos8=A=4x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FM,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:          "Volvo Trucks Global (Brazil)",
			vin:           "9BVAG30A550000000", // 9BV=Volvo Brazil, A=FM, G30=Diesel, Pos8=A=4x2, Pos10=5=2005
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FM,
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2005,
			wantSuccess:   true,
		},
		{
			name:        "Volvo Bus (Global)",
			vin:         "YV3A1234500000000", // YV3=Volvo Bus
			wantBrand:   vinv1.Brand_VOLVO_BUSES,
			wantModel:   vinv1.Model_MODEL_UNSPECIFIED,
			wantType:    vinv1.VehicleType_BUS,
			wantSuccess: true,
		},
		{
			name:          "Volvo Trucks Global - 2023 Model Year",
			vin:           "YV2TG30A5P0000000", // T=FL, G30=Diesel, Pos8=A=4x2, Pos10=P=2023
			wantBrand:     vinv1.Brand_VOLVO_TRUCKS,
			wantModel:     vinv1.Model_FL, // T = FL
			wantType:      vinv1.VehicleType_HEAVY_GOODS_VEHICLE,
			wantFuelTypes: []vinv1.FuelType{vinv1.FuelType_DIESEL},
			wantAxleCount: 2,
			wantYear:      2023,
			wantSuccess:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := DecodeVehicle(tt.vin)
			if tt.wantSuccess {
				if !ok {
					t.Errorf("DecodeVehicle() failed")
					return
				}

				// Construct expected object
				want := vinv1.Vehicle_builder{
					DataSources: []vinv1.DataSource{vinv1.DataSource_DEEP_RESEARCH},
				}
				if tt.wantBrand != vinv1.Brand_BRAND_UNSPECIFIED {
					want.Brand = new(tt.wantBrand)
				}
				if tt.wantModel != vinv1.Model_MODEL_UNSPECIFIED {
					want.Model = new(tt.wantModel)
				}
				if tt.wantType != vinv1.VehicleType_VEHICLE_TYPE_UNSPECIFIED {
					want.Type = new(tt.wantType)
				}
				if len(tt.wantFuelTypes) > 0 {
					want.FuelTypes = tt.wantFuelTypes
				}
				if tt.wantAxleCount > 0 {
					want.AxleCount = new(tt.wantAxleCount)
				}
				if tt.wantYear > 0 {
					want.Year = new(tt.wantYear)
				}

				if diff := cmp.Diff(want.Build(), got, protocmp.Transform()); diff != "" {
					t.Errorf("DecodeVehicle() mismatch (-want +got):\n%s", diff)
				}
			} else {
				if ok {
					t.Errorf("DecodeVehicle() succeeded unexpectedly")
				}
			}
		})
	}
}
