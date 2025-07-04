package main

import (
	"fmt"
	"scantool/scan"
)

func main() {
	fmt.Println("Scan Tool starting...")
	scan.GetHostname()
	scan.IsVirtual()
	scan.HostOS()
	scan.CoreCount()
	scan.PrintRAM()
	scan.GetHostIP()
}
