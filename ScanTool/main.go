package main

import (
	"fmt"
	collector "scantool/scan"
)

func main() {
	fmt.Println("Scan Tool starting...")
	collector.GetHostname()
	collector.IsVirtual()
	collector.HostOS()
	collector.CoreCount()
	collector.PrintRAM()
	collector.GetHostIP()
}
