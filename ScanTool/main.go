package main

import (
	"fmt"
	"scantool/scan"
)

func main() {
	fmt.Println("Scan Tool starting...")
	scan.CoreCount()
	scan.PrintRAM()
}
