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

func TestPack(t *testing.T) {
	tests := []struct {
		wmi1, wmi2 string
		packed     uint32
		ok         bool
	}{
		{"W1T", "1A9", (41537 << 16) | 1665, true},
		{"3VW", "OQI", (5036 << 16) | 32058, true},
		{"foo", "bar", 0, false},
	}
	for _, test := range tests {
		packed, ok := Pack(test.wmi1, test.wmi2)
		if packed != test.packed {
			t.Errorf("Pack(%s, %s) = %d, want %d", test.wmi1, test.wmi2, packed, test.packed)
		}
		if ok != test.ok {
			t.Errorf("Pack(%s, %s) ok = %v, want %v", test.wmi1, test.wmi2, ok, test.ok)
		}
		if ok {
			w1, w2, okUnpack := Unpack(packed)
			if !okUnpack {
				t.Errorf("Unpack failed for valid packed %d", packed)
			}
			if w1 != test.wmi1 || w2 != test.wmi2 {
				t.Errorf("Unpack(%d) = %s, %s; want %s, %s", packed, w1, w2, test.wmi1, test.wmi2)
			}
		}
	}
}

func Fuzz_PackUnpack(f *testing.F) {
	f.Add("W1T", "1A9")
	f.Add("3VW", "OQI")
	f.Fuzz(func(t *testing.T, w1, w2 string) {
		packed, ok := Pack(w1, w2)
		if !ok {
			return
		}
		out1, out2, ok := Unpack(packed)
		if !ok {
			t.Errorf("Unpack failed for packed value %d (from %s, %s)", packed, w1, w2)
			return
		}
		if out1 != w1 || out2 != w2 {
			t.Errorf("Round trip failed! Input: %s, %s -> Packed: %d -> Output: %s, %s", w1, w2, packed, out1, out2)
		}
	})
}
