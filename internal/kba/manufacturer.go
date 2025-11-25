package kba

import (
	"slices"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// ResolveBrand resolves the brand of a KBA manufacturer ID.
func ResolveBrand(manufacturerID int32) (vinv1.Brand, bool) {
	values := vinv1.Brand_BRAND_UNSPECIFIED.Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		kbaManufacturerID, ok := proto.GetExtension(value.Options(), vinv1.E_KbaManufacturerId).([]int32)
		if !ok || len(kbaManufacturerID) == 0 {
			continue
		}
		if slices.Contains(kbaManufacturerID, manufacturerID) {
			return vinv1.Brand(value.Number()), true
		}
	}
	return vinv1.Brand_BRAND_UNSPECIFIED, false
}
