package vin

import (
	"sort"

	"github.com/way-platform/vin-go/internal/wmi"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// LookupLowVolumeManufacturer finds a low volume manufacturer by WMI1 and WMI2.
func LookupLowVolumeManufacturer(wmi1, wmi2 string) (*vinv1.Manufacturer, bool) {
	k, ok := wmi.Pack(wmi1, wmi2)
	if !ok {
		return nil, false
	}
	i := sort.Search(len(lvmIndex), func(i int) bool {
		return lvmIndex[i].K >= k
	})
	if i >= len(lvmIndex) || lvmIndex[i].K != k {
		return nil, false
	}
	start := lvmIndex[i].O
	end := uint32(len(lvmBlob))
	if i+1 < len(lvmIndex) {
		end = lvmIndex[i+1].O
	}
	var output vinv1.Manufacturer
	if err := proto.Unmarshal(lvmBlob[start:end], &output); err != nil {
		return nil, false
	}
	return &output, true
}
