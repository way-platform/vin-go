package fordvin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDecodeDate(t *testing.T) {
	tests := []struct {
		yearCode  byte
		monthCode byte
		wantYear  int
		wantMonth int
		wantErr   bool
		name      string
	}{
		{name: "2024 Jan", yearCode: 'R', monthCode: 'L', wantYear: 2024, wantMonth: 1, wantErr: false},
		{name: "2024 Feb", yearCode: 'R', monthCode: 'Y', wantYear: 2024, wantMonth: 2, wantErr: false},
		{name: "2024 Mar", yearCode: 'R', monthCode: 'S', wantYear: 2024, wantMonth: 3, wantErr: false},
		{name: "2024 Apr", yearCode: 'R', monthCode: 'T', wantYear: 2024, wantMonth: 4, wantErr: false},
		{name: "2024 May", yearCode: 'R', monthCode: 'J', wantYear: 2024, wantMonth: 5, wantErr: false},
		{name: "2024 Jun", yearCode: 'R', monthCode: 'U', wantYear: 2024, wantMonth: 6, wantErr: false},
		{name: "2024 Jul", yearCode: 'R', monthCode: 'M', wantYear: 2024, wantMonth: 7, wantErr: false},
		{name: "2024 Aug", yearCode: 'R', monthCode: 'P', wantYear: 2024, wantMonth: 8, wantErr: false},
		{name: "2024 Sep", yearCode: 'R', monthCode: 'B', wantYear: 2024, wantMonth: 9, wantErr: false},
		{name: "2024 Oct", yearCode: 'R', monthCode: 'R', wantYear: 2024, wantMonth: 10, wantErr: false},
		{name: "2024 Nov", yearCode: 'R', monthCode: 'A', wantYear: 2024, wantMonth: 11, wantErr: false},
		{name: "2024 Dec", yearCode: 'R', monthCode: 'G', wantYear: 2024, wantMonth: 12, wantErr: false},

		{name: "2025 Jan", yearCode: 'S', monthCode: 'C', wantYear: 2025, wantMonth: 1, wantErr: false},
		{name: "2025 Feb", yearCode: 'S', monthCode: 'K', wantYear: 2025, wantMonth: 2, wantErr: false},
		{name: "2025 Dec", yearCode: 'S', monthCode: 'P', wantYear: 2025, wantMonth: 12, wantErr: false},

		{name: "2026 Jan", yearCode: 'T', monthCode: 'B', wantYear: 2026, wantMonth: 1, wantErr: false},
		{name: "2026 Feb", yearCode: 'T', monthCode: 'R', wantYear: 2026, wantMonth: 2, wantErr: false},

		{name: "2027 Jan", yearCode: 'V', monthCode: 'J', wantYear: 2027, wantMonth: 1, wantErr: false},
		{name: "2027 Feb", yearCode: 'V', monthCode: 'U', wantYear: 2027, wantMonth: 2, wantErr: false},

		{name: "Invalid month code for year", yearCode: 'S', monthCode: 'B', wantYear: 0, wantMonth: 0, wantErr: true}, // 'B' is unused in 2025 (Year 2)
		{name: "Invalid year code", yearCode: 'Z', monthCode: 'C', wantYear: 0, wantMonth: 0, wantErr: true},
		{name: "Invalid month char", yearCode: 'S', monthCode: 'Z', wantYear: 0, wantMonth: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYear, gotMonth, err := DecodeDate(tt.yearCode, tt.monthCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.wantYear, gotYear); diff != "" {
					t.Errorf("DecodeDate() year mismatch (-want +got):\n%s", diff)
				}
				if diff := cmp.Diff(tt.wantMonth, gotMonth); diff != "" {
					t.Errorf("DecodeDate() month mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
