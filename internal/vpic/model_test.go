package vpic

import (
	"testing"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
)

func TestResolveModel(t *testing.T) {
	tests := []struct {
		modelID int32
		model   vinv1.Model
		ok      bool
	}{
		{modelID: 1703, model: vinv1.Model_SPRINTER, ok: true},
		{modelID: 3120, model: vinv1.Model_SPRINTER, ok: true},
		{modelID: 3608, model: vinv1.Model_TRANSIT, ok: true},
		{modelID: 22426, model: vinv1.Model_E_CANTER, ok: true},
		{modelID: 0, model: vinv1.Model_MODEL_UNSPECIFIED, ok: false},
		{modelID: 999999, model: vinv1.Model_MODEL_UNSPECIFIED, ok: false},
	}
	for _, test := range tests {
		model, ok := ResolveModel(test.modelID)
		if model != test.model {
			t.Errorf("ResolveModel(%d) = %v, want %v", test.modelID, model, test.model)
		}
		if ok != test.ok {
			t.Errorf("ResolveModel(%d) = %v, want %v", test.modelID, ok, test.ok)
		}
	}
}
