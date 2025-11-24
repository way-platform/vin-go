package vpic

import (
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// ResolveBrand resolves the brand of a vPIC make ID.
func ResolveBrand(makeID int32) (vinv1.Brand, bool) {
	values := vinv1.Brand_BRAND_UNSPECIFIED.Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		vpicMakeId, ok := proto.GetExtension(value.Options(), vinv1.E_VpicMakeId).(int32)
		if !ok || vpicMakeId == 0 {
			continue
		}
		if vpicMakeId == makeID {
			return vinv1.Brand(value.Number()), true
		}
	}
	return vinv1.Brand_BRAND_UNSPECIFIED, false
}
