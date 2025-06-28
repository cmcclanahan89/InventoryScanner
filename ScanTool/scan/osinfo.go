package scan

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func CoreCount() {
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		fmt.Println("Error getting logical cores:", err)
		return
	}
	fmt.Println("Logical CPU Cores:", logicalCores)

	physicalCores, err := cpu.Counts(false)
	if err != nil {
		fmt.Println("Error getting physical cores:", err)
		return
	}
	fmt.Println("Phsycial CPU Cores:", physicalCores)
}

func PrintRAM() {

	ramAmount, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting RAM Count:", err)
	}
	fmt.Println("RAM Amount:", int64(ramAmount.Total)/(1<<30))
}
