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

// HostOS retrieves the name of the host operating system by executing a PowerShell command.
// It returns the OS name as a string and an error if the command execution fails.
// Note: This function is intended for use on Windows systems.
func HostOS() (string, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-Command", `Get-CimInstance Win32_Operatingsystem | Select-Object -expand Caption`)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(output)), nil
}

// CoreCount returns (logicalCores, physicalCores, error) in that order.
// CoreCount returns the number of logical and physical CPU cores on the system.
// It returns the logical core count, physical core count, and an error if any occurred
// during the retrieval of core counts.
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

// GetRam retrieves the total amount of system RAM in GiB (Gibibytes).
// It uses the mem.VirtualMemory function to obtain memory statistics.
// If an error occurs while fetching the memory information, the function logs the error and terminates the program.
// Returns the total RAM as a uint64 value representing GiB.
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

// GetLocalAdminUsers retrieves the list of local administrator user names on the current Windows machine.
// It executes a PowerShell command to query the "Administrators" group and returns the names as a slice of strings.
// Returns an error if the command fails or cannot be executed.
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
