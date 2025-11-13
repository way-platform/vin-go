package vin

import (
	"testing"
)

func TestShouldSkipCheckDigitValidation(t *testing.T) {
	tests := []struct {
		name       string
		vin        string
		shouldSkip bool
		reason     string
	}{
		// North America (1, 4, 5) - MUST validate
		{
			name:       "North America - position 1 = 1",
			vin:        "1HGBH41JXMN109186",
			shouldSkip: false,
			reason:     "North American VINs (position 1 = 1, 4, 5) require check digit validation",
		},
		{
			name:       "North America - position 1 = 4",
			vin:        "4T1BF1FK4CU123456",
			shouldSkip: false,
			reason:     "North American VINs (position 1 = 1, 4, 5) require check digit validation",
		},
		{
			name:       "North America - position 1 = 5",
			vin:        "5YJSA1E14HF123456",
			shouldSkip: false,
			reason:     "North American VINs (position 1 = 1, 4, 5) require check digit validation",
		},

		// Africa (A-C) - SKIP validation
		{
			name:       "Africa - position 1 = A",
			vin:        "AAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "African VINs (position 1 = A-C) do not require check digit validation",
		},
		{
			name:       "Africa - position 1 = B",
			vin:        "BAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "African VINs (position 1 = A-C) do not require check digit validation",
		},
		{
			name:       "Africa - position 1 = C",
			vin:        "CAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "African VINs (position 1 = A-C) do not require check digit validation",
		},

		// Central/South America (D-G) - SKIP validation
		{
			name:       "Central/South America - position 1 = D",
			vin:        "DAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Central/South American VINs (position 1 = D-G) do not require check digit validation",
		},
		{
			name:       "Central/South America - position 1 = E",
			vin:        "EAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Central/South American VINs (position 1 = D-G) do not require check digit validation",
		},
		{
			name:       "Central/South America - position 1 = F",
			vin:        "FAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Central/South American VINs (position 1 = D-G) do not require check digit validation",
		},
		{
			name:       "Central/South America - position 1 = G",
			vin:        "GAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Central/South American VINs (position 1 = D-G) do not require check digit validation",
		},

		// Asia (H-R) - SKIP validation
		{
			name:       "Asia - position 1 = H",
			vin:        "HAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = J",
			vin:        "JAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = K",
			vin:        "KAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = L",
			vin:        "LAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = M",
			vin:        "MAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = N",
			vin:        "NAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = P",
			vin:        "PAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = R",
			vin:        "RAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},

		// Europe (S-Z) - SKIP validation
		{
			name:       "Europe - position 1 = S",
			vin:        "SAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = T",
			vin:        "TAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = U",
			vin:        "UAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = V",
			vin:        "VAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = W (Ford Germany)",
			vin:        "WF0WXXTACWLD32086",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = X",
			vin:        "XAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = Y (Scania)",
			vin:        "YS2R4X20005690181",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = Z",
			vin:        "ZAVAA1A11AA123456",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},

		// Low Volume Manufacturers (position 3 = '9') - SKIP validation
		{
			name:       "Low Volume - position 3 = 9, position 1 = 1",
			vin:        "1A912345678901234",
			shouldSkip: true,
			reason:     "Low volume manufacturers (position 3 = 9) do not require check digit validation",
		},
		{
			name:       "Low Volume - position 3 = 9, position 1 = W",
			vin:        "WA912345678901234",
			shouldSkip: true,
			reason:     "Low volume manufacturers (position 3 = 9) do not require check digit validation",
		},

		// Australia Government Issued (6ZZ) - SKIP validation
		{
			name:       "Australia Government - 6ZZ",
			vin:        "6ZZ12345678901234",
			shouldSkip: true,
			reason:     "Australia government-issued VINs (6ZZ) do not require check digit validation",
		},

		// Edge cases
		{
			name:       "Short VIN",
			vin:        "12",
			shouldSkip: false,
			reason:     "Short VINs should default to requiring validation",
		},
		{
			name:       "Empty VIN",
			vin:        "",
			shouldSkip: false,
			reason:     "Empty VINs should default to requiring validation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shouldSkipCheckDigitValidation(tt.vin)
			if result != tt.shouldSkip {
				t.Errorf("shouldSkipCheckDigitValidation(%q) = %v, want %v. Reason: %s",
					tt.vin, result, tt.shouldSkip, tt.reason)
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
			vin:       "1HGBH41JXMN109186",
			wantErr:   false,
			wantValid: true,
		},
		{
			name:      "Invalid check digit",
			vin:       "1HGBH41JXMN109187", // Changed last digit
			wantErr:   false,
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := validateCheckDigit(tt.vin)
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

func TestValidateCheckDigit_NonNorthAmerica(t *testing.T) {
	tests := []struct {
		name      string
		vin       string
		wantErr   bool
		wantValid bool
	}{
		{
			name:      "European Ford (should skip)",
			vin:       "WF0WXXTACWLD32086",
			wantErr:   false,
			wantValid: true, // Should pass because validation is skipped
		},
		{
			name:      "Scania (should skip)",
			vin:       "YS2R4X20005690181",
			wantErr:   false,
			wantValid: true, // Should pass because validation is skipped
		},
		{
			name:      "Low volume manufacturer (should skip)",
			vin:       "1A912345678901234",
			wantErr:   false,
			wantValid: true, // Should pass because validation is skipped
		},
		{
			name:      "Australia government (should skip)",
			vin:       "6ZZ12345678901234",
			wantErr:   false,
			wantValid: true, // Should pass because validation is skipped
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := validateCheckDigit(tt.vin)
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
