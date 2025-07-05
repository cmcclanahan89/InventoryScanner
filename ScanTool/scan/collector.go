package scan

import "time"

func Collect() (MachineScan, error) {
	var ms MachineScan
	// Initialize the MachineScan struct

	ms.SchemaVersion = "1.0"          // initial version, bump when the shape changes
	ms.CollectedAt = time.Now().UTC() // set the timestamp

	h, err := GetHostname() // call the helper
	if err != nil {
		return ms, err // bubble the problem up
	}
	ms.Hostname = h // fill the struct field

	// ... collect other fields ...
	v, err := IsVirtual()
	if err != nil {
		return ms, err
	}
	ms.Arch = v

	os, err := HostOS()
	if err != nil {
		return ms, err
	}
	ms.OS = os

	ram := GetRam()
	ms.Memory = ram

	return ms, nil
}
