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

	SaveScanErr := scan.SaveScan(ms)
	if SaveScanErr != nil {
		log.Fatalf("Failed to save scan data: %v", SaveScanErr)
	}

}

// The above code collects machine scan data and prints it in JSON format.
