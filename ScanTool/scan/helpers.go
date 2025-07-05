package scan

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

func GetHostname() (string, error) {
	return os.Hostname()
}

func IsVirtual() (string, error) {
	// Check for Dell keyword in system model name. All physical hardware is Dell.
	cmd := exec.Command("powershell",
		"-NoProfile", "-Command",
		`if((Get-WmiObject -Class Win32_ComputerSystem).Model -match '*Dell*') { exit 0 } else { exit 1 }`,
	)
	err := cmd.Run()
	if err == nil {
		// PowerShell exited 0 → virtualization detected
		return "Virtual", nil
	}

	// If it’s a non-zero exit code, see whether it’s exactly 1 (physical).
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
		return "Physical", nil // bare-metal
	}

	// Any other error means the command itself failed.
	return "", fmt.Errorf("failed to execute powershell: %w", err)
}

func HostOS() (string, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-Command", `Get-CimInstance Win32_Operatingsystem | Select-Object -expand Caption`)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(output)), nil
}

// CoreCount returns (logicalCores, physicalCores, error) in that order.
func CoreCount() (int, int, error) {
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		return 0, 0, err
	}
	physicalCores, err := cpu.Counts(false)
	if err != nil {
		return 0, 0, err
	}
	return logicalCores, physicalCores, nil
}

func GetRam() uint64 {

	ramAmount, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	return (ramAmount.Total / (1 << 30)) // Convert bytes to GiB
}

func GetHostIP() net.IP {

	hostIP, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer hostIP.Close()

	localAddress := hostIP.LocalAddr().(*net.UDPAddr)
	// fmt.Println("IP Address:", localAddress.IP)
	return localAddress.IP

}

// GetLocalAdminUsers returns a list of local admin users on the machine.
func GetLocalAdminUsers() ([]string, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-Command",
		`Get-LocalGroupMember -Group "Administrators" | Select-Object -ExpandProperty Name`,
	)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get local admin users: %w", err)
	}

	users := strings.Split(strings.TrimSpace(string(output)), "\n")
	for i, user := range users {
		users[i] = strings.TrimSpace(user) // Clean up whitespace
	}
	return users, nil
}
