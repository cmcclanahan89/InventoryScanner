package scan

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	probing "github.com/prometheus-community/pro-bing"
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

func ScanHost() (string, error) {
	// 1) Prompt for the hostname
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter hostname to ping: ")
	hostInput, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %v", err)
	}
	hostInput = strings.TrimSpace(hostInput)

	// 2) Resolve DNS name
	ips, err := net.LookupIP(hostInput)
	if err != nil {
		return "", fmt.Errorf("DNS lookup failed for %q: %v", hostInput, err)
	}
	resolvedIP := ips[0].String()

	// 3) Set up the pinger
	pinger, err := probing.NewPinger(resolvedIP)
	if err != nil {
		return "", fmt.Errorf("failed to create pinger for %s: %v", resolvedIP, err)
	}
	pinger.Count = 3
	pinger.Timeout = 5 * time.Second
	pinger.SetPrivileged(true) // Windows: run as Administrator

	// 4) Run the ping
	fmt.Printf("Pinging %s (%s)...\n", hostInput, resolvedIP)
	if err := pinger.Run(); err != nil {
		return "", fmt.Errorf("ping error: %v", err)
	}
	stats := pinger.Statistics()

	// 5) Check and return
	if stats.PacketsRecv > 0 {
		fmt.Printf("Host %s is reachable at IP %s\n", hostInput, resolvedIP)
		return resolvedIP, nil
	}
	return "", fmt.Errorf("host %s did not respond to ping", resolvedIP)

}
