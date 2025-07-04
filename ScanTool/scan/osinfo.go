package scan

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hostname:", hostname)
}

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

func GetHostIP() net.IP {

	hostIP, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Error getting IP Address:", err)
	}
	defer hostIP.Close()

	localAddress := hostIP.LocalAddr().(*net.UDPAddr)
	fmt.Println("IP Address:", localAddress.IP)
	return localAddress.IP

}
