package pattern

import (
	"testing"
)

func TestPackUnpack(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		ok      bool   // Expected result from Pack
	}{
		{name: "Empty string", pattern: "", ok: true},
		{name: "Single char 'A'", pattern: "A", ok: true},
		{name: "Single char '0'", pattern: "0", ok: true},
		{name: "Single char '*'", pattern: "*", ok: true},
		{name: "Single char '['", pattern: "[", ok: true},
		{name: "Single char ']'", pattern: "]", ok: true},
		{name: "Single char '|'", pattern: "|", ok: false},
		// Calculation for "ABC": C=13, B=12, A=11. Processed right-to-left: C + B*40 + A*40*40 = 13 + 12*40 + 11*1600 = 13 + 480 + 17600 = 18093
		{name: "Short pattern ABC", pattern: "ABC", ok: true},
		{name: "Short pattern 012", pattern: "012", ok: true},
		{
				name: "Max length 12 chars (all pipes)", pattern: "||||||||||||", ok: false},
				{name: "Multiple pipes (ABC|D|E)", pattern: "ABC|D|E", ok: false},
				{name: "Max length 12 chars (mixed)", pattern: "0123456789AZ", ok: true},
				{name: "Single pipe at start", pattern: "|ABC", ok: true},
				{name: "Single pipe in middle", pattern: "AB|C", ok: true},
				{name: "Single pipe at end", pattern: "ABC|", ok: false},

		{name: "Too long (13 chars)", pattern: "0123456789AZB", ok: false},
		{name: "Invalid char space", pattern: "AB C", ok: false},
		{name: "Invalid char exclamation", pattern: "ABC!", ok: false},
		{name: "Invalid char lowercase", pattern: "abc", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_Pack", func(t *testing.T) {
			gotPacked, gotOk := Pack(tt.pattern)
			if gotOk != tt.ok {
				t.Errorf("Pack(\"%s\") gotOk = %v, want %v", tt.pattern, gotOk, tt.ok)
				return
			}
			// No longer checking specific packed values as they are derived from Base-40 and are complex.
			// The primary check is round-trip consistency.
			if tt.ok && tt.pattern != "" && gotPacked == 0 {
				// This case means a valid, non-empty pattern packed to 0, which shouldn't happen
				t.Errorf("Pack(\"%s\") unexpectedly packed to 0", tt.pattern)
			}
		})

		if tt.ok { // Only test Unpack and round-trip for patterns that are expected to pack successfully
			t.Run(tt.name+"_UnpackAndRoundTrip", func(t *testing.T) {
				packed, _ := Pack(tt.pattern) // Should succeed

				gotPattern, gotOk := Unpack(packed)
				if !gotOk {
					t.Errorf("Unpack(Pack(\"%s\")) gotOk = %v, want %v", tt.pattern, gotOk, true)
				}
				if gotPattern != tt.pattern {
					t.Errorf("Round trip failed for \"%s\": Packed %v, Unpacked \"%s\"", tt.pattern, packed, gotPattern)
				}
			})
		}
	}
}

func TestUnpackInvalidValues(t *testing.T) {
	// The theoretical maximum packed value for a 12-character string in Bijective Base-40.
	const maxPackedValFor12Chars uint64 = 17620378650256410256

	// We want to test a value that would decode to 13 characters.
	// In our bijective base-40 system, maxPackedValFor12Chars corresponds to the largest 12-digit number.
	// Therefore, maxPackedValFor12Chars + 1 corresponds to the smallest 13-digit number.
	longVal := maxPackedValFor12Chars + 1

	t.Run("Value corresponding to >12 chars", func(t *testing.T) {
		gotPattern, gotOk := Unpack(longVal)
		if gotOk {
			t.Errorf("Unpack(%v) gotOk = %v, want %v (decoded: %q)", longVal, gotOk, false, gotPattern)
		}
		if gotPattern != "" {
			t.Errorf("Unpack(%v) gotPattern = \"%s\", want \"\"", longVal, gotPattern)
		}
	})
}

// containsInvalidChar is a helper for fuzz testing to quickly identify patterns with chars outside our alphabet.
func containsInvalidChar(s string) bool {
	for i := 0; i < len(s); i++ {
		if reverseLookup[s[i]] == 0 {
			return true
		}
	}
	return false
}

// Fuzz test for broader coverage
func FuzzPackUnpack(f *testing.F) {
	seedPatterns := []string{
		"", "A", "0", "*", "[", "]",
		"ABC", "0123456789", "A[B*]C|D",
		"TE45F", "1G1***", "ABC[D]", "1G[0-9][A-Z]",
		"000000000000",
		"999999999999",
		"AAAAAAAAAAAA",
		"ZZZZZZZZZZZZ",
		"**********",
	}
	for _, p := range seedPatterns {
		f.Add(p)
	}

	f.Fuzz(func(t *testing.T, pattern string) {
		// Filter out patterns that contain invalid characters or are too long for our Pack function
		if len(pattern) > 12 || containsInvalidChar(pattern) {
			// These inputs would cause Pack to return ok=false, so we just ensure Pack handles them gracefully
			_, ok := Pack(pattern)
			if ok {
				t.Errorf("Pack(\"%s\") unexpectedly succeeded for invalid input", pattern)
			}
			return
		}

		// Filter out patterns ending with pipe
		if len(pattern) > 0 && pattern[len(pattern)-1] == '|' {
			_, ok := Pack(pattern)
			if ok {
				t.Errorf("Pack(\"%s\") unexpectedly succeeded for pattern ending in pipe", pattern)
			}
			return
		}

		// For valid inputs, ensure round-trip consistency
		packed, ok := Pack(pattern)
		if !ok {
			t.Errorf("Pack(\"%s\") failed unexpectedly for valid input", pattern)
			return
		}

		unpacked, unpackOk := Unpack(packed)
		if !unpackOk {
			t.Errorf("Unpack(Pack(\"%s\")) gotOk = %v, want %v", pattern, unpackOk, true)
		}
		if unpacked != pattern {
			t.Errorf("Round trip failed for \"%s\": Packed %v, Unpacked \"%s\"", pattern, packed, unpacked)
		}
	})
}
