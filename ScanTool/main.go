package main

import (
	"encoding/json"
	"fmt"
	"log"
	"scantool/scan"
)

func main() {
	fmt.Println("Scan Tool starting...")

	ms, err := scan.Collect()
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.MarshalIndent(ms, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}

// The above code collects machine scan data and prints it in JSON format.
