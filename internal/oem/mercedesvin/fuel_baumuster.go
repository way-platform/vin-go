package mercedesvin

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

// decodeFuelTypeBaumuster infers fuel type based on the Baumuster series.
// This applies to Heavy Trucks and other vehicles identified by their 3-digit series.
func decodeFuelTypeBaumuster(series string, subtype string) []vinv1.FuelType {
	switch series {
	// Electric Series
	case "983": // eActros
		return []vinv1.FuelType{vinv1.FuelType_ELECTRIC}

	// Diesel Heavy Truck Series
	case "930", "932", "933", "934", // Actros MP1-3
		"963",                      // Actros MP4/5 / Antos
		"964",                      // Arocs
		"967",                      // Atego (New)
		"950", "952", "953", "954", // Axor
		"949", "959", // Zetros
		"970", "972", "974", "975", "976": // Atego (Classic)
		return []vinv1.FuelType{vinv1.FuelType_DIESEL}

	// Unimog (Diesel)
	case "405", "406", "416", "424", "425", "435", "437":
		return []vinv1.FuelType{vinv1.FuelType_DIESEL}

	// Econic (Diesel vs Electric vs Gas)
	case "956", "957":
		// 956/957 are standard Econic (Diesel) but eEconic also shares 956.
		// Without specific subtype for eEconic, we default to Diesel as it's the vast majority.
		// eEconic detection should ideally happen via subtype if known.
		return []vinv1.FuelType{vinv1.FuelType_DIESEL}
	case "958": // Econic NGT
		return []vinv1.FuelType{vinv1.FuelType_NATURAL_GAS} // CNG/LNG

	// Sprinter 910 (FWD) - Check for Electric Subtype
	case "910":
		if isElectricSubtype(series, subtype) {
			return []vinv1.FuelType{vinv1.FuelType_ELECTRIC}
		}
		// Default 910 is Diesel
		return []vinv1.FuelType{vinv1.FuelType_DIESEL}

	// Sprinter 901-906 (Legacy)
	case "901", "902", "903", "904", "905", "906", "907":
		// Primarily Diesel, but Gas exists. Defaulting to Diesel is 99% accurate for commercial EU.
		return []vinv1.FuelType{vinv1.FuelType_DIESEL}
	}

	return nil
}
