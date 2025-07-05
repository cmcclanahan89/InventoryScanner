// internal/scan/schema.go
package scan

import "time"

// Keep everything in one place so both agent & server import the *same* type.
type MachineScan struct {
	SchemaVersion string    `json:"schema_version"` // bump when the shape changes
	CollectedAt   time.Time `json:"collected_at"`   // RFC 3339 for readability
	Hostname      string    `json:"hostname"`
	OS            string    `json:"os"`
	Arch          string    `json:"arch"`

	CPU    CPUStats    `json:"cpu"`
	Memory uint64      `json:"memory"`
	Disks  []DiskStats `json:"disks,omitempty"` // omit if empty
	Net    []NetStats  `json:"net,omitempty"`
}

// Nested types keep the top struct readable.
type CPUStats struct {
	CoresP int `json:"coresp"`
	CoresV int `json:"coresv"`
}

type MemoryStats struct {
	Max uint64 `json:"total_bytes"`
}

type DiskStats struct {
	Mount string  `json:"mount"`
	FS    string  `json:"fs"`
	Used  uint64  `json:"used_bytes"`
	Util  float64 `json:"used_percent"`
}

type NetStats struct {
	IPAddress string `json:"ipaddress"`
}
