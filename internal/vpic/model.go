package vpic

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// ResolveModel resolves the model of a vPIC model ID.
func ResolveModel(modelID int32) (vinv1.Model, bool) {
	values := vinv1.Model_MODEL_UNSPECIFIED.Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		ext := proto.GetExtension(value.Options(), vinv1.E_VpicModelId)
		if ids, ok := ext.([]int32); ok {
			for _, id := range ids {
				if id == modelID {
					return vinv1.Model(value.Number()), true
				}
			}
		} else if id, ok := ext.(int32); ok {
			if id == modelID {
				return vinv1.Model(value.Number()), true
			}
		}
	}
	return vinv1.Model_MODEL_UNSPECIFIED, false
}
