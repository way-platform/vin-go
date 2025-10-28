package vin

// GeographicArea represents the geographic region for VIN manufacturing
type GeographicArea string

const (
	GeographicArea_GEOGRAPHIC_AREA_UNSPECIFIED   GeographicArea = "UNSPECIFIED"
	GeographicArea_GEOGRAPHIC_AREA_AFRICA        GeographicArea = "AFRICA"
	GeographicArea_GEOGRAPHIC_AREA_ASIA          GeographicArea = "ASIA"
	GeographicArea_GEOGRAPHIC_AREA_EUROPE        GeographicArea = "EUROPE"
	GeographicArea_GEOGRAPHIC_AREA_NORTH_AMERICA GeographicArea = "NORTH_AMERICA"
	GeographicArea_GEOGRAPHIC_AREA_OCEANIA       GeographicArea = "OCEANIA"
	GeographicArea_GEOGRAPHIC_AREA_SOUTH_AMERICA GeographicArea = "SOUTH_AMERICA"
)

func lookupGeographicArea(code rune) GeographicArea {
	switch code {
	case '1':
		return GeographicArea_GEOGRAPHIC_AREA_NORTH_AMERICA
	case '2':
		return GeographicArea_GEOGRAPHIC_AREA_NORTH_AMERICA
	case '3':
		return GeographicArea_GEOGRAPHIC_AREA_NORTH_AMERICA
	case '4':
		return GeographicArea_GEOGRAPHIC_AREA_NORTH_AMERICA
	case '5':
		return GeographicArea_GEOGRAPHIC_AREA_NORTH_AMERICA
	case '6':
		return GeographicArea_GEOGRAPHIC_AREA_OCEANIA
	case '7':
		return GeographicArea_GEOGRAPHIC_AREA_OCEANIA
	case '8':
		return GeographicArea_GEOGRAPHIC_AREA_SOUTH_AMERICA
	case '9':
		return GeographicArea_GEOGRAPHIC_AREA_SOUTH_AMERICA
	case 'A':
		return GeographicArea_GEOGRAPHIC_AREA_AFRICA
	case 'B':
		return GeographicArea_GEOGRAPHIC_AREA_AFRICA
	case 'C':
		return GeographicArea_GEOGRAPHIC_AREA_AFRICA
	case 'D':
		return GeographicArea_GEOGRAPHIC_AREA_AFRICA
	case 'E':
		return GeographicArea_GEOGRAPHIC_AREA_AFRICA
	case 'F':
		return GeographicArea_GEOGRAPHIC_AREA_AFRICA
	case 'G':
		return GeographicArea_GEOGRAPHIC_AREA_AFRICA
	case 'H':
		return GeographicArea_GEOGRAPHIC_AREA_AFRICA
	case 'J':
		return GeographicArea_GEOGRAPHIC_AREA_ASIA
	case 'K':
		return GeographicArea_GEOGRAPHIC_AREA_ASIA
	case 'L':
		return GeographicArea_GEOGRAPHIC_AREA_ASIA
	case 'M':
		return GeographicArea_GEOGRAPHIC_AREA_ASIA
	case 'N':
		return GeographicArea_GEOGRAPHIC_AREA_ASIA
	case 'P':
		return GeographicArea_GEOGRAPHIC_AREA_ASIA
	case 'Q':
		return GeographicArea_GEOGRAPHIC_AREA_ASIA
	case 'R':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	case 'S':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	case 'T':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	case 'U':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	case 'V':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	case 'W':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	case 'X':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	case 'Y':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	case 'Z':
		return GeographicArea_GEOGRAPHIC_AREA_EUROPE
	default:
		return GeographicArea_GEOGRAPHIC_AREA_UNSPECIFIED
	}
}
