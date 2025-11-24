package wmi

import (
	"strings"

	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// ResolveRegion resolves the region of a WMI.
func ResolveRegion(wmi string) (vinv1.Region, bool) {
	values := vinv1.Region_REGION_UNSPECIFIED.Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		value := values.Get(i)
		wmiPrefixes, ok := proto.GetExtension(value.Options(), vinv1.E_WmiPrefix).([]string)
		if !ok {
			continue
		}
		for _, wmiPrefix := range wmiPrefixes {
			if strings.HasPrefix(wmi, wmiPrefix) {
				return vinv1.Region(value.Number()), true
			}
		}
	}
	return vinv1.Region_REGION_UNSPECIFIED, false
}
