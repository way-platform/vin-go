package checkdigit

import (
	"fmt"
	"strings"
)

// Calculate calculates the expected check digit for a VIN
func Calculate(vin string) (string, error) {
	if len(vin) != 17 {
		return "", fmt.Errorf("VIN must be 17 characters")
	}
	vin = strings.ToUpper(vin)
	sum := 0
	for i, char := range vin {
		value, ok := vinCharValues[char]
		if !ok {
			return "", fmt.Errorf("invalid character '%c' at position %d", char, i+1)
		}
		// Position 9 (index 8) has weight 0, so it doesn't contribute to the sum
		sum += value * vinWeights[i]
	}
	remainder := sum % 11
	if remainder == 10 {
		return "X", nil
	}
	return fmt.Sprintf("%d", remainder), nil
}

// vinCharValues maps each valid VIN character to its numeric value for check digit calculation
var vinCharValues = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8,
	'J': 1, 'K': 2, 'L': 3, 'M': 4, 'N': 5, 'P': 7, 'R': 9,
	'S': 2, 'T': 3, 'U': 4, 'V': 5, 'W': 6, 'X': 7, 'Y': 8, 'Z': 9,
}

// vinWeights are the position weights for check digit calculation (positions 1-17)
var vinWeights = []int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}
