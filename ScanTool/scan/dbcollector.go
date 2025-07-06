package scan

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	db, err := sql.Open("sqlite3", "./scans.db")
	if err != nil {
		return err
	}
	defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS machine_scans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		schema_version TEXT,
		collected_at   TIMESTAMP,
		hostname       TEXT,
		os             TEXT,
		arch           TEXT,
		logical_cores  INTEGER,
		physical_cores INTEGER,
		memory         TEXT,
		ip_address     TEXT,
		admin_users    TEXT
	);
	`
	_, err = db.Exec(createTable)
	return err
}

func SaveScan(ms MachineScan) error {
	db, err := sql.Open("sqlite3", "./scans.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Convert admin_users to JSON
	adminUsersJSON, err := json.Marshal(ms.AdminUsers)
	if err != nil {
		return err
	}

	fmt.Printf("Saving: %+v\n", ms)

	result, err := db.Exec(`
		INSERT INTO machine_scans
		(schema_version, collected_at, hostname, os, arch, logical_cores, physical_cores, memory, ip_address, admin_users)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		ms.SchemaVersion,
		ms.CollectedAt,
		ms.Hostname,
		ms.OS,
		ms.Arch,
		ms.LogicalCores,
		ms.PhysicalCores,
		ms.Memory,
		ms.IPAddress,
		string(adminUsersJSON),
	)
	if err != nil {
		fmt.Printf("Error inserting row: %v\n", err)
		return err
	}

	rows, _ := result.RowsAffected()
	fmt.Printf("Rows affected: %d\n", rows)
	return nil
}
