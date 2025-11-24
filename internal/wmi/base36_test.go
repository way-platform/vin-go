package wmi

import "testing"

func Fuzz_roundTrip(f *testing.F) {
	// Seed with some known valid WMIs
	f.Add("W1T")     // Toyota
	f.Add("1A9")     // US LVM
	f.Add("3VW")     // VW Mexico
	f.Add("OQI")     // Forbidden letters used in practice
	f.Add("invalid") // Invalid input
	f.Fuzz(func(t *testing.T, input string) {
		base36, ok := ToBase36(input)
		if !ok {
			return
		}
		decoded, ok := FromBase36(base36)
		if !ok {
			return
		}
		if input != decoded {
			t.Errorf("Round trip failed! Input: %s -> Packed: %d -> Output: %s", input, base36, decoded)
		}
	})
}

func TestToBase36(t *testing.T) {
	tests := []struct {
		wmi    string
		base36 uint16
		ok     bool
	}{
		{wmi: "W1T", base36: 41537, ok: true},
		{wmi: "1A9", base36: 1665, ok: true},
		{wmi: "3VW", base36: 5036, ok: true},
		{wmi: "OQI", base36: 32058, ok: true},
		{wmi: "invalid", base36: 0, ok: false},
	}
	for _, test := range tests {
		base36, ok := ToBase36(test.wmi)
		if base36 != test.base36 {
			t.Errorf("ToBase36(%s) = %d, want %d", test.wmi, base36, test.base36)
		}
		if ok != test.ok {
			t.Errorf("ToBase36(%s) = %v, want %v", test.wmi, ok, test.ok)
		}
	}
}
