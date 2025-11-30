package pattern

import (
	"math"
	"testing"
)

func TestMatches(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		vin      string
		expected bool
	}{
		// Standard VDS/VIS Patterns (Offset 3)
		{"ExactMatch", "FW1", "1FTFW100000000000", true},
		{"ExactMatchTooLongInput", "FW1", "1FTFW1E0000000000", true},
		{"ExactMatchTooShortInput", "FW1E", "1FTFW100000000000", false},

		{"WildcardSingle", "F*1", "1FTFW100000000000", true},
		{"WildcardSingleNoMatch", "F*1", "1FTFX200000000000", false},

		{"WildcardEnd", "FW*", "1FTFW1E0000000000", true},
		{"WildcardEndExact", "FW1*", "1FTFW100000000000", true},

		{"CharClassListMatch", "F[WX]1", "1FTFW100000000000", true},
		{"CharClassListNoMatch", "F[WX]1", "1FTFY100000000000", false},

		{"CharClassRangeMatch", "F[0-9]1", "1FTF5100000000000", true},
		{"CharClassRangeNoMatch", "F[0-4]1", "1FTF5100000000000", false},

		{"CharClassCombined", "F[A-C1-3]1", "1FTFB100000000000", true},
		{"CharClassCombinedDigit", "F[A-C1-3]1", "1FTF2100000000000", true},
		{"CharClassCombinedNoMatch", "F[A-C1-3]1", "1FTFD100000000000", false},

		{"EmptyPatternInvalidVIN", "", "", false},
		{"EmptyPatternValidVIN", "", "1FTFW100000000000", false},

		// VIS / Plant Code Patterns (Offset 10)
		{"VISPatternMatchExact", "*****|*F", "1FTFW1E83LF000000", true},
		{"VISPatternNoMatch", "*****|*F", "1FTFW1E83LX000000", false},
		{"VISPatternWildcardMetadata", "*****|**", "1FTFW1E83LA000000", true},
		{"VISPatternWildcardMetadata2", "*****|**", "1FTFW1E83LZ000000", true},

		{"VISPatternInvalidVIN", "*****|*U", "TOO_SHORT", false},

		{"MixedPatternMatch", "1G*[A-C]", "1FT1GBA0000000000", true},
		{"MixedPatternMatch2", "1G*[A-C]", "1FT1GDA0000000000", true},

		{"ComplexWildcardMatch", "A*B*C", "1FTAXBYC000000000", true},
		{"ComplexWildcardNoMatch", "A*B*C", "1FTAXBC0000000000", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Matches(tt.pattern, tt.vin)
			if got != tt.expected {
				t.Errorf("Matches(%q, %q) = %v, want %v", tt.pattern, tt.vin, got, tt.expected)
			}
		})
	}
}

func TestScore(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		vin      string
		expected float64
	}{
		// Standard VDS/VIS Scoring
		{"ExactMatchFullScore", "FW1", "1FTFW100000000000", 1.0},
		{"ExactMatchMismatch", "FW1", "1FTFW200000000000", 0.0},

		// Pattern: F*1 vs Input: FW1... -> F(1) + *(0.5) + 1(1) = 2.5/3
		{"WildcardHalfScore", "F*1", "1FTFW100000000000", (1.0 + 0.5 + 1.0) / 3.0},

		// Pattern: FW* vs Input: FW1... -> F(1) + W(1) + *(0.5) = 2.5/3
		{"WildcardEndPartialScore", "FW*", "1FTFW100000000000", (1.0 + 1.0 + 0.5) / 3.0},

		// Pattern: F* vs Input: FW1... -> F(1) + *(0.5) = 1.5/2 = 0.75
		{"WildcardEndMultiChar", "F*", "1FTFW100000000000", (1.0 + 0.5) / 2.0},

		// Pattern: FW* vs Input: FW1234... -> F(1) + W(1) + *(0.5) = 2.5/3
		{"WildcardEndLongInput", "FW*", "1FTFW123400000000", (1.0 + 1.0 + 0.5) / 3.0},

		{"CharClassListScore", "F[WX]1", "1FTFW100000000000", (1.0 + 0.8 + 1.0) / 3.0},
		{"CharClassRangeScore", "F[0-9]1", "1FTF5100000000000", (1.0 + 0.7 + 1.0) / 3.0},
		{"CharClassRangeNoMatch", "F[0-4]1", "1FTF5100000000000", 0.0},

		// VIS / Plant Code Scoring
		{"VISPatternScoreExact", "*****|*F", "1FTFW1E83LF000000", 1.0},
		{"VISPatternScoreWildcardMetadata", "*****|**", "1FTFW1E83LA000000", 0.8},
		{"VISPatternScoreNoMatch", "*****|*F", "1FTFW1E83LX000000", 0.0},

		{"MismatchReturnsZero", "FW1", "1FTFX100000000000", 0.0},
		{"InvalidVINReturnsZero", "FW1", "TOO_SHORT", 0.0},

		// Pattern: 1G*[A-C] vs Input: 1GBA...
		// 1(1) + G(1) + *(0.5) + [A-C](0.7) = 3.2/4 = 0.8
		{"ComplexPatternScoring", "1G*[A-C]", "1FT1GBA0000000000", 0.8},

		// Pattern: 1G[A-C]* vs Input: 1GBCDE...
		// 1(1) + G(1) + [A-C](0.7) + *(0.5) = 3.2/4 = 0.8
		{"ComplexPatternScoring2", "1G[A-C]*", "1FT1GBCDE00000000", 0.8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Score(tt.pattern, tt.vin)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("Score(%q, %q) = %v, want %v", tt.pattern, tt.vin, got, tt.expected)
			}
		})
	}
}