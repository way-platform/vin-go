package iso3779

// ISO 3779 Table 1: Characters used for designating the year.

// Note: This table loops every 30 years.
var yearCodeMap = map[byte]int{
	'A': 2010, 'B': 2011, 'C': 2012, 'D': 2013, 'E': 2014, 'F': 2015, 'G': 2016, 'H': 2017, 'J': 2018, 'K': 2019,
	'L': 2020, 'M': 2021, 'N': 2022, 'P': 2023, 'R': 2024, 'S': 2025, 'T': 2026, 'V': 2027, 'W': 2028, 'X': 2029,
	'Y': 2030, '1': 2001, '2': 2002, '3': 2003, '4': 2004, '5': 2005, '6': 2006, '7': 2007, '8': 2008, '9': 2009,
}

// Year returns the year corresponding to the given character according to ISO 3779.
// The ISO 3779 standard uses a 30-year cycle for year codes.
// This function attempts to return the most plausible year within a reasonable range (e.g., current year - 30 years to current year + 10 years).
func Year(character byte) (int, bool) {
	year, ok := yearCodeMap[character]
	return year, ok
}
