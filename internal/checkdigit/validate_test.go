package checkdigit

import "testing"

func TestValidateCheckDigit_NonNorthAmerica(t *testing.T) {
	tests := []struct {
		name      string
		vin       string
		wantErr   bool
		wantValid bool
	}{
		{
			name:      "European Ford (should skip)",
			vin:       "WF0WXXTAC00000000",
			wantErr:   false,
			wantValid: true, // Should pass because validation is skipped
		},
		{
			name:      "Scania (should skip)",
			vin:       "YS2R4X20000000000",
			wantErr:   false,
			wantValid: true, // Should pass because validation is skipped
		},
		{
			name:      "Low volume manufacturer (should skip)",
			vin:       "1A912345600000000",
			wantErr:   false,
			wantValid: true, // Should pass because validation is skipped
		},
		{
			name:      "Australia government (should skip)",
			vin:       "6ZZ12345600000000",
			wantErr:   false,
			wantValid: true, // Should pass because validation is skipped
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := Validate(tt.vin)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCheckDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if valid != tt.wantValid {
				t.Errorf("validateCheckDigit() = %v, want %v", valid, tt.wantValid)
			}
		})
	}
}

func TestValidateCheckDigit_NorthAmerica(t *testing.T) {
	tests := []struct {
		name      string
		vin       string
		wantErr   bool
		wantValid bool
	}{
		{
			name:      "Valid North American VIN",
			vin:       "1HGBH41J700000000",
			wantErr:   false,
			wantValid: true,
		},
		{
			name:      "Invalid check digit",
			vin:       "1HGBH41J800000000",
			wantErr:   false,
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := Validate(tt.vin)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCheckDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if valid != tt.wantValid {
				t.Errorf("validateCheckDigit() = %v, want %v", valid, tt.wantValid)
			}
		})
	}
}
