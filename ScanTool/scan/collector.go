package scan

import (
	"fmt"
	"time"
)

func Collect() (MachineScan, error) {
	var ms MachineScan

	// Initialize the MachineScan struct

	ms.SchemaVersion = "1.0"          // initial version, bump when the shape changes
	ms.CollectedAt = time.Now().UTC() // set the timestamp

	// collect hostname, OS, architecture, memory, CPU cores, IP address, and local admin users
	// and populate the MachineScan struct
	h, err := GetHostname()
	if err != nil {
		return ms, err
	}
	ms.Hostname = h

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
	ms.Memory = fmt.Sprintf("%d GiB", ram)

	logicalCores, physicalCores, err := CoreCount()
	if err != nil {
		return ms, err
	}
	ms.LogicalCores = logicalCores
	ms.PhysicalCores = physicalCores

	ip := GetHostIP()
	ms.IPAddress = ip.String() // convert net.IP to string

	adminUsers, err := GetLocalAdminUsers()
	if err != nil {
		return ms, fmt.Errorf("failed to collect local admin users: %w", err)
	}
	ms.AdminUsers = adminUsers

	return ms, nil
}
