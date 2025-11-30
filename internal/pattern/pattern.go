package pattern

import "strings"

// alphabet defines the allowed characters in the pattern.
// 0-9, A-Z, *, [, ], |
const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ*[]|"

var reverseLookup [256]uint8

func init() {
	for i := range reverseLookup {
		reverseLookup[i] = 0 // 0 means invalid/unmapped
	}
	for i := 0; i < len(alphabet); i++ {
		// Map character to value 1..40 (Bijective Base-40)
		reverseLookup[alphabet[i]] = uint8(i + 1)
	}
}

// Pack converts a pattern string into a single uint64 using Bijective Base-40 encoding.
// Returns 0, false if the string contains invalid characters or exceeds 12 characters.
func Pack(s string) (uint64, bool) {
	// Max capacity is 12 characters.
	// 40^12 fits within uint64 max value.
	if len(s) > 12 {
		return 0, false
	}

	if len(s) > 0 && s[len(s)-1] == '|' {
		return 0, false // Pattern cannot end with a pipe
	}

	pipeCount := 0
	for _, r := range s {
		if r == '|' {
			pipeCount++
		}
	}

	if pipeCount > 1 {
		return 0, false // More than one pipe character is not allowed
	}

	var val uint64
	var power uint64 = 1

	// Process string from right to left (LSB to MSB)
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		digit := reverseLookup[char]
		if digit == 0 {
			return 0, false // Invalid character found
		}

		// val += digit * power
		val += uint64(digit) * power
		
		// Only multiply power if not the last iteration
		if i > 0 {
			power *= 40
		}
	}

	return val, true
}

// Unpack converts a uint64 back into the pattern string.
// It returns the unpacked string and a boolean indicating if the uint64
// corresponds to a pattern of valid length (<= 12 characters).
func Unpack(val uint64) (string, bool) {
	if val == 0 {
		return "", true // Empty string is valid and packed as 0
	}

	var sb strings.Builder
	var count int // To track unpacked length

	currentVal := val
	for currentVal > 0 {
		if count == 12 { // If we've already decoded 12 characters, and currentVal > 0, then the original string was > 12 chars
			return "", false
		}

		// Bijective decoding: digit = (currentVal-1) % 40 + 1
		idx := (currentVal - 1) % 40
		sb.WriteByte(alphabet[idx])
		currentVal = (currentVal - 1) / 40
		count++
	}

	// The builder constructed the string in reverse order (LSB first).
	// Reverse the string.
	runes := []rune(sb.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), true
}
