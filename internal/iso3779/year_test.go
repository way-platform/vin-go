package iso3779

import (
	"testing"
)

func TestYear(t *testing.T) {
	tests := []struct {
		character byte
		expectedYear int
		expectError bool
	}{
		// Codes directly from Table 1
		{character: 'M', expectedYear: 2021, expectError: false},
		{character: 'N', expectedYear: 2022, expectError: false},
		{character: 'P', expectedYear: 2023, expectError: false},
		{character: 'R', expectedYear: 2024, expectError: false},
		{character: 'S', expectedYear: 2025, expectError: false},
		{character: 'T', expectedYear: 2026, expectError: false},
		{character: 'V', expectedYear: 2027, expectError: false},
		{character: 'W', expectedYear: 2028, expectError: false},
		{character: 'X', expectedYear: 2029, expectError: false},
		{character: 'Y', expectedYear: 2030, expectError: false},
		{character: '1', expectedYear: 2001, expectError: false},
		{character: '2', expectedYear: 2002, expectError: false},
		{character: '3', expectedYear: 2003, expectError: false},
		{character: '4', expectedYear: 2004, expectError: false},
		{character: '5', expectedYear: 2005, expectError: false},
		{character: '6', expectedYear: 2006, expectError: false},
		{character: '7', expectedYear: 2007, expectError: false},
		{character: '8', expectedYear: 2008, expectError: false},
		{character: '9', expectedYear: 2009, expectError: false},
		{character: 'A', expectedYear: 2010, expectError: false},
		{character: 'B', expectedYear: 2011, expectError: false},
		{character: 'C', expectedYear: 2012, expectError: false},
		{character: 'D', expectedYear: 2013, expectError: false},
		{character: 'E', expectedYear: 2014, expectError: false},
		{character: 'F', expectedYear: 2015, expectError: false},
		{character: 'G', expectedYear: 2016, expectError: false},
		{character: 'H', expectedYear: 2017, expectError: false},
		{character: 'J', expectedYear: 2018, expectError: false},
		{character: 'K', expectedYear: 2019, expectError: false},
		{character: 'L', expectedYear: 2020, expectError: false},

		// Invalid characters
		{character: 'Z', expectedYear: 0, expectError: true},
		{character: '0', expectedYear: 0, expectError: true},
		{character: 'I', expectedYear: 0, expectError: true},
		{character: 'O', expectedYear: 0, expectError: true},
		{character: 'Q', expectedYear: 0, expectError: true},
	}

	for _, tt := range tests {
		t.Run(string(tt.character), func(t *testing.T) {
			year, ok := Year(tt.character)
			if ok == tt.expectError {
				t.Errorf("Year(%c) ok = %v, expectError %v", tt.character, ok, tt.expectError)
				return
			}
			if !tt.expectError && year != tt.expectedYear {
				t.Errorf("Year(%c) = %d, expected %d", tt.character, year, tt.expectedYear)
			}
		})
	}
}
