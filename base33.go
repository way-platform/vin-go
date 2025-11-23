package vin

var base33LookupTable = makeBase33LookupTable()

func makeBase33LookupTable() [256]uint8 {
	var table [256]uint8
	for i := range table {
		table[i] = 0xFF
	}
	for i := '0'; i <= '9'; i++ {
		table[i] = uint8(i - '0')
	}
	// A-H
	for i := 'A'; i <= 'H'; i++ {
		table[i] = uint8(i - 'A' + 10)
	}
	// J-N
	for i := 'J'; i <= 'N'; i++ {
		table[i] = uint8(i - 'J' + 18)
	}
	// P
	table['P'] = 23
	// R-Z
	for i := 'R'; i <= 'Z'; i++ {
		table[i] = uint8(i - 'R' + 24)
	}
	return table
}

// packBase33 converts a 3-char string to a compact uint16 index.
// Returns 0, false if the input contains invalid characters or bad length.
func wmiToBase33(s string) (uint16, bool) {
	if len(s) != 3 {
		return 0, false
	}
	// We use direct array indexing for O(1) conversion
	v0 := base33LookupTable[s[0]]
	v1 := base33LookupTable[s[1]]
	v2 := base33LookupTable[s[2]]
	// Check for 0xFF (invalid char marker)
	if v0 == 0xFF || v1 == 0xFF || v2 == 0xFF {
		return 0, false
	}
	// Formula: d1 * 33^2 + d2 * 33^1 + d3 * 33^0
	// 33^2 = 1089, 33^1 = 33
	return uint16(v0)*1089 + uint16(v1)*33 + uint16(v2), true
}
