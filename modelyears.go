package vin

func lookupModelYear(code string) (int, bool) {
	switch code {
	case "1":
		return 2001, true
	case "2":
		return 2002, true
	case "3":
		return 2003, true
	case "4":
		return 2004, true
	case "5":
		return 2005, true
	case "6":
		return 2006, true
	case "7":
		return 2007, true
	case "8":
		return 2008, true
	case "9":
		return 2009, true
	case "A":
		return 2010, true
	case "B":
		return 2011, true
	case "C":
		return 2012, true
	case "D":
		return 2013, true
	case "E":
		return 2014, true
	case "F":
		return 2015, true
	case "G":
		return 2016, true
	case "H":
		return 2017, true
	case "J":
		return 2018, true
	case "K":
		return 2019, true
	case "L":
		return 2020, true
	case "M":
		return 2021, true
	case "N":
		return 2022, true
	case "P":
		return 2023, true
	case "R":
		return 2024, true
	case "S":
		return 2025, true
	case "T":
		return 2026, true
	case "V":
		return 2027, true
	case "W":
		return 2028, true
	case "X":
		return 2029, true
	case "Y":
		return 2030, true
	default:
		return 0, false
	}
}
