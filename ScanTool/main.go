package main

import (
	"fmt"
	"log"
	"scantool/scan"
)

func main() {
	fmt.Println("Scan Tool starting...")

	initDBErr := scan.InitDB()
	if initDBErr != nil {
		log.Fatalf("Failed to initialize database: %v", initDBErr)
	}

	ms, err := scan.Collect()
	if err != nil {
		log.Fatalf("Failed to collect Data: %v", err)
	}

	exists, err := scan.GetExistingHostname(ms.Hostname)
	if err != nil {
		log.Fatalf("Failed to check existing hostname: %v", err)
	}

	if exists {
		log.Printf("Hostname %s already exists in the database, updating record... \n", ms.Hostname)
		if err := scan.UpdateMachineScan(ms); err != nil {
			log.Fatalf("Failed to update machine scan: %v", err)
		}
	} else {
		fmt.Printf("Machine scan for %s does not exist. Creating new record.\n", ms.Hostname)
		if err := scan.InsertScan(ms); err != nil {
			log.Fatalf("Failed to insert machine scan: %v", err)
		}
	}

	// SaveScanErr := scan.InsertScan(ms)
	// if SaveScanErr != nil {
	// 	log.Fatalf("Failed to save scan data: %v", SaveScanErr)
	// }

}

// The above code collects machine scan data and prints it in JSON format.
