package scan

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
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
