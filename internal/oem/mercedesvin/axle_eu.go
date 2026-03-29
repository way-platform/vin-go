package mercedesvin

// decodeEUAxleFromDriveCode extracts the axle count from the 4th digit of a
// Mercedes-Benz EU commercial vehicle VIN. When the 4th digit is alphabetic,
// it directly encodes the drive configuration per the manufacturer's standard.
//
// Reference: Mercedes-Benz Trucks VIN Structure (service-info.mercedes-benz-trucks.com)
func decodeEUAxleFromDriveCode(digit byte) (int32, bool) {
	switch digit {
	case 'Y', 'D': // 4x2, 4x4
		return 2, true
	case 'T', 'K', 'L': // 6x2, 6x4, 6x6
		return 3, true
	case 'N', 'P', 'S': // 8x4, 8x6, 8x8
		return 4, true
	case '2': // 8x2
		return 4, true
	case '7': // 10x4
		return 5, true
	case 'E': // 10x6
		return 5, true
	default:
		return 0, false
	}
}

// decodeEUAxleFromBaumuster extracts the axle count from the full 6-digit
// Baumuster (VIN positions 4-9) when the 4th digit is numeric. The Baumuster
// suffix encodes the chassis configuration for each series.
//
// This is a Tier 2 fallback when the alphabetic drive code is not present.
func decodeEUAxleFromBaumuster(baumuster string) (int32, bool) {
	if len(baumuster) != 6 {
		return 0, false
	}
	series := baumuster[0:3]
	suffix := baumuster[3:6]
	axles, ok := baumusterAxleMap[series+suffix]
	if ok {
		return axles, true
	}
	// Try suffix-only lookup for well-known series.
	axles, ok = seriesSuffixAxleMap[series][suffix]
	return axles, ok
}

// baumusterAxleMap maps specific full Baumuster codes to axle counts.
var baumusterAxleMap = map[string]int32{}

// seriesSuffixAxleMap maps Baumuster series to suffix → axle count.
// Documented examples from deep research on Mercedes-Benz Actros, Arocs, Atego.
var seriesSuffixAxleMap = map[string]map[string]int32{
	// Actros MP4/MP5 (963 series)
	"963": {
		"025": 3, // 6x2 tag axle
		"403": 2, // 4x2 tractor
	},
}
