package wmi

const base36Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var base36LookupTable = makeBase36LookupTable()

func makeBase36LookupTable() [256]uint8 {
	var table [256]uint8
	for i := range table {
		table[i] = 0xFF
	}
	for i := '0'; i <= '9'; i++ {
		table[i] = uint8(i - '0')
	}
	// A-Z (all letters including I, O, Q)
	for i := 'A'; i <= 'Z'; i++ {
		table[i] = uint8(i - 'A' + 10)
	}
	return table
}

// ToBase36 converts a 3-character WMI string to a base36 index.
func ToBase36(wmi string) (uint16, bool) {
	if len(wmi) != 3 {
		return 0, false
	}
	// We use direct array indexing for O(1) conversion
	v0 := base36LookupTable[wmi[0]]
	v1 := base36LookupTable[wmi[1]]
	v2 := base36LookupTable[wmi[2]]
	// Check for 0xFF (invalid char marker)
	if v0 == 0xFF || v1 == 0xFF || v2 == 0xFF {
		return 0, false
	}
	// Formula: d1 * 36^2 + d2 * 36^1 + d3 * 36^0
	// 36^2 = 1296, 36^1 = 36
	return uint16(v0)*1296 + uint16(v1)*36 + uint16(v2), true
}

// FromBase36 converts a base36 index to a 3-character WMI string.
func FromBase36(base36 uint16) (string, bool) {
	// Safety check: Max index is 36^3 - 1 = 46655
	if base36 >= 36*36*36 {
		return "", false
	}
	// 1. Calculate the indices for each position
	// Math: id = (d1 * 36^2) + (d2 * 36^1) + d3
	// First char (Most Significant)
	idx0 := base36 / (36 * 36) // id / 1296
	remainder := base36 % (36 * 36)
	idx1 := remainder / 36
	idx2 := remainder % 36
	// 2. Look up the characters
	return string([]byte{
		base36Alphabet[idx0],
		base36Alphabet[idx1],
		base36Alphabet[idx2],
	}), true
}
