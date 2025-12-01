package vin

import (
	"sort"

	wmipkg "github.com/way-platform/vin-go/internal/wmi"
	vinv1 "github.com/way-platform/vin-go/proto/gen/go/wayplatform/connect/vin/v1"
	"google.golang.org/protobuf/proto"
)

// LookupManufacturer finds a standard manufacturer by WMI (World Manufacturer Identifier).
func LookupManufacturer(wmi string) (*vinv1.Manufacturer, bool) {
	key, ok := wmipkg.ToBase36(wmi)
	if !ok {
		return nil, false
	}
	i := sort.Search(len(wmiIndex), func(i int) bool {
		return wmiIndex[i].K >= key
	})
	if i >= len(wmiIndex) || wmiIndex[i].K != key {
		return nil, false
	}
	start := wmiIndex[i].O
	end := uint32(len(wmiBlob))
	if i+1 < len(wmiIndex) {
		end = wmiIndex[i+1].O
	}
	var output vinv1.Manufacturer
	if err := proto.Unmarshal(wmiBlob[start:end], &output); err != nil {
		return nil, false
	}
	return &output, true
}
