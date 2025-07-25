// internal/scan/schema.go
package scan

import "time"

// Keep everything in one place so both agent & server import the *same* type.
type MachineScan struct {
	SchemaVersion string      `json:"schema_version"` // bump when the shape changes
	CollectedAt   time.Time   `json:"collected_at"`   // RFC 3339 for readability
	Hostname      string      `json:"hostname"`
	OS            string      `json:"os"`
	Arch          string      `json:"arch"`
	LogicalCores  int         `json:"logical_cores"`  // logical CPU cores
	PhysicalCores int         `json:"physical_cores"` // physical CPU cores
	Memory        string      `json:"RAM"`
	Disks         []DiskStats `json:"disks,omitempty"` // omit if empty
	IPAddress     string      `json:"IP Address,omitempty"`
	AdminUsers    []string    `json:"admin_users,omitempty"` // local admin users, if any
}

type DiskStats struct {
	Mount string  `json:"mount"`
	FS    string  `json:"fs"`
	Used  uint64  `json:"used_bytes"`
	Util  float64 `json:"used_percent"`
}
