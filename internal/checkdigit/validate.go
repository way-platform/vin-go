package checkdigit

import (
	"fmt"
	"strings"
)

// Validate validates the check digit in position 9 of the VIN.
// Returns true (valid) if the VIN is from a manufacturer that doesn't use
// check digit validation, or if the check digit matches the calculated value.
func Validate(vin string) (bool, error) {
	if len(vin) != 17 {
		return false, fmt.Errorf("VIN must be 17 characters")
	}
	// Skip check digit validation for manufacturers that don't use it
	if ShouldSkipValidation(vin) {
		return true, nil
	}
	vin = strings.ToUpper(vin)
	actualCheckDigit := string(vin[8]) // Position 9 (0-indexed as 8)
	expectedCheckDigit, err := Calculate(vin)
	if err != nil {
		return false, err
	}
	return actualCheckDigit == expectedCheckDigit, nil
}
