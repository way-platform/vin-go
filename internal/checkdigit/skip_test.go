package checkdigit

import (
	"testing"
)

func TestShouldSkipValidation(t *testing.T) {
	tests := []struct {
		name       string
		vin        string
		shouldSkip bool
		reason     string
	}{
		// North America (1, 4, 5) - MUST validate
		{
			name:       "North America - position 1 = 1",
			vin:        "1HGBH41JX00000000",
			shouldSkip: false,
			reason:     "North American VINs (position 1 = 1, 4, 5) require check digit validation",
		},
		{
			name:       "North America - position 1 = 4",
			vin:        "4T1BF1FK400000000",
			shouldSkip: false,
			reason:     "North American VINs (position 1 = 1, 4, 5) require check digit validation",
		},
		{
			name:       "North America - position 1 = 5",
			vin:        "5YJSA1E1400000000",
			shouldSkip: false,
			reason:     "North American VINs (position 1 = 1, 4, 5) require check digit validation",
		},

		// Africa (A-C) - SKIP validation
		{
			name:       "Africa - position 1 = A",
			vin:        "AAVAA1A1100000000",
			shouldSkip: true,
			reason:     "African VINs (position 1 = A-C) do not require check digit validation",
		},
		{
			name:       "Africa - position 1 = B",
			vin:        "BAVAA1A1100000000",
			shouldSkip: true,
			reason:     "African VINs (position 1 = A-C) do not require check digit validation",
		},
		{
			name:       "Africa - position 1 = C",
			vin:        "CAVAA1A1100000000",
			shouldSkip: true,
			reason:     "African VINs (position 1 = A-C) do not require check digit validation",
		},

		// Central/South America (D-G) - SKIP validation
		{
			name:       "Central/South America - position 1 = D",
			vin:        "DAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Central/South American VINs (position 1 = D-G) do not require check digit validation",
		},
		{
			name:       "Central/South America - position 1 = E",
			vin:        "EAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Central/South American VINs (position 1 = D-G) do not require check digit validation",
		},
		{
			name:       "Central/South America - position 1 = F",
			vin:        "FAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Central/South American VINs (position 1 = D-G) do not require check digit validation",
		},
		{
			name:       "Central/South America - position 1 = G",
			vin:        "GAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Central/South American VINs (position 1 = D-G) do not require check digit validation",
		},

		// Asia (H-R) - SKIP validation
		{
			name:       "Asia - position 1 = H",
			vin:        "HAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = J",
			vin:        "JAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = K",
			vin:        "KAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = L",
			vin:        "LAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = M",
			vin:        "MAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = N",
			vin:        "NAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = P",
			vin:        "PAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},
		{
			name:       "Asia - position 1 = R",
			vin:        "RAVAA1A1100000000",
			shouldSkip: true,
			reason:     "Asian VINs (position 1 = H-R) do not require check digit validation",
		},

		// Europe (S-Z) - SKIP validation
		{
			name:       "Europe - position 1 = S",
			vin:        "SAVAA1A1100000000",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = T",
			vin:        "TAVAA1A1100000000",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = U",
			vin:        "UAVAA1A1100000000",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = V",
			vin:        "VAVAA1A1100000000",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = W (Ford Germany)",
			vin:        "WF0WXXTAC00000000",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = X",
			vin:        "XAVAA1A1100000000",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = Y (Scania)",
			vin:        "YS2R4X20000000000",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},
		{
			name:       "Europe - position 1 = Z",
			vin:        "ZAVAA1A1100000000",
			shouldSkip: true,
			reason:     "European VINs (position 1 = S-Z) do not require check digit validation",
		},

		// Low Volume Manufacturers (position 3 = '9') - SKIP validation
		{
			name:       "Low Volume - position 3 = 9, position 1 = 1",
			vin:        "1A912345600000000",
			shouldSkip: true,
			reason:     "Low volume manufacturers (position 3 = 9) do not require check digit validation",
		},
		{
			name:       "Low Volume - position 3 = 9, position 1 = W",
			vin:        "WA912345600000000",
			shouldSkip: true,
			reason:     "Low volume manufacturers (position 3 = 9) do not require check digit validation",
		},

		// Australia Government Issued (6ZZ) - SKIP validation
		{
			name:       "Australia Government - 6ZZ",
			vin:        "6ZZ12345600000000",
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
			result := ShouldSkipValidation(tt.vin)
			if result != tt.shouldSkip {
				t.Errorf("shouldSkipCheckDigitValidation(%q) = %v, want %v. Reason: %s",
					tt.vin, result, tt.shouldSkip, tt.reason)
			}
		})
	}
}
