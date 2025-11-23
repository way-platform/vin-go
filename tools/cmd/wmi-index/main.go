package main

import "strconv"

// Record represents a record in the WMI index.
type Record struct {
	WMI          string `json:"wmi"`
	Base33       uint16 `json:"base33"`
	Manufacturer string `json:"manufacturer"`
	Country      string `json:"country"`
	Model        string `json:"model"`
}

// CSV returns the record as a CSV record.
func (r *Record) CSV() []string {
	return []string{
		r.WMI,
		strconv.FormatUint(uint64(r.Base33), 10),
		r.Manufacturer,
		r.Country,
		r.Model,
	}
}
