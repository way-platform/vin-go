package vin

import (
	_ "embed"
	"encoding/json"
	"sync"
)

//go:embed wmi.json
var wmiJSON []byte
var wmiEntries map[string]wmiEntry
var wmiEntriesOnce sync.Once

type wmiEntry struct {
	Code         string `json:"wmi"`
	Manufacturer string `json:"manufacturer"`
	Country      string `json:"country"`
}

func lookupWMI(wmi string) wmiEntry {
	wmiEntriesOnce.Do(func() {
		wmis := make([]wmiEntry, 0)
		err := json.Unmarshal(wmiJSON, &wmis)
		if err != nil {
			panic(err)
		}
		wmiEntries = make(map[string]wmiEntry)
		for _, wmi := range wmis {
			wmiEntries[wmi.Code] = wmi
		}
	})
	entry, ok := wmiEntries[wmi]
	if !ok {
		return wmiEntry{
			Code:         "Unspecified",
			Manufacturer: "Unspecified",
			Country:      "Unspecified",
		}
	}
	return entry
}
