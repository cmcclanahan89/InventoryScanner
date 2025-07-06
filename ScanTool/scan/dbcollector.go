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

func InsertScan(ms MachineScan) error {
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

func GetExistingHostname(targetHostname string) (bool, error) {
	db, err := sql.Open("sqlite3", "./scans.db")
	if err != nil {
		return false, fmt.Errorf("failed to open DB: %w", err)
	}
	defer db.Close()

	var hostname string
	err = db.QueryRow("SELECT hostname FROM machine_scans WHERE hostname = ? LIMIT 1", targetHostname).Scan(&hostname)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No rows found, return empty string
		}
		return false, fmt.Errorf("failed to query hostname: %w", err)
	}
	return true, nil
}

func UpdateMachineScan(ms MachineScan) error {
	db, err := sql.Open("sqlite3", "./scans.db")
	if err != nil {
		return err
	}
	defer db.Close()

	adminUsersJSON, err := json.Marshal(ms.AdminUsers)
	if err != nil {
		return err
	}

	tableUpdate, err := db.Exec(`
		UPDATE machine_scans
		SET schema_version = ?, collected_at = ?, os = ?, arch = ?, logical_cores = ?, physical_cores = ?, memory = ?, ip_address = ?, admin_users = ?
		WHERE hostname = ?`,
		ms.SchemaVersion,
		ms.CollectedAt,
		ms.OS,
		ms.Arch,
		ms.LogicalCores,
		ms.PhysicalCores,
		ms.Memory,
		ms.IPAddress,
		string(adminUsersJSON),
		ms.Hostname,
	)
	if err != nil {
		fmt.Printf("Error inserting row: %v\n", err)
		return err
	}

	rows, _ := tableUpdate.RowsAffected()
	fmt.Printf("Rows affected: %d\n", rows)
	return nil

}
