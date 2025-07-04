package collector

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

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

func IsVirtual() (bool, error) {
	// Check for Dell keyword in system model name. All physical hardware is Dell.
	cmd := exec.Command("powershell",
		"-NoProfile", "-Command",
		`if((Get-WmiObject -Class Win32_ComputerSystem).Model -match '*Dell*') { exit 0 } else { exit 1 }`,
	)
	err := cmd.Run()
	if err == nil {
		// PowerShell exited 0 → virtualization detected
		return true, nil
	}
	fmt.Println("Virtual")

	// If it’s a non-zero exit code, see whether it’s exactly 1 (physical).
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
		return false, nil // bare-metal
	}
	fmt.Println("Physical")

	// Any other error means the command itself failed.
	return false, fmt.Errorf("failed to execute powershell: %w", err)
}

func HostOS() {
	cmd := exec.Command("powershell",
		"-NoProfile", "-Command", `Get-CimInstance Win32_Operatingsystem | Select-Object -expand Caption`)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(strings.Trim(string(output), "\r\n"))
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
