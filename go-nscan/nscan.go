package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

const DIAL_TIMEOUT = 100

// ICMP SCAN
func isHostAvailable(address string) bool {
	pinger, err := probing.NewPinger(address)
	if err != nil {
		fmt.Printf("Error creating pinger: %v\n", err)
		return false
	}
	pinger.Timeout = DIAL_TIMEOUT * time.Millisecond
	pinger.Count = 1
	pinger.SetPrivileged(true) // required for most systems

	if err := pinger.Run(); err != nil {
		fmt.Printf("Error running pinger: %v\n", err)
		return false
	}

	return pinger.Statistics().PacketsRecv > 0
}

func resolveHostnameToIP(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return nil, err
	}

	var ipStrings []string
	for _, ip := range ips {
		// filter out IPv6 addresses by checking the length
		if ip.To4() != nil {
			ipStrings = append(ipStrings, ip.String())
		}
	}

	return ipStrings, nil
}

func scanIPRange(ipRange []string, numWorkers int) []string {
	ipChan := make(chan string, len(ipRange))
	resultChan := make(chan string, len(ipRange))
	var wg sync.WaitGroup

	// start worker goroutines for scanning
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ip := range ipChan {
				if isHostAvailable(ip) {
					resultChan <- ip
				}
			}
		}()
	}

	// send IPs to workers
	go func() {
		for _, ip := range ipRange {
			ipChan <- ip
		}
		close(ipChan)
	}()

	// wait for all workers to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// collect results from resultChan
	availableIPs := make([]string, 0)
	for ip := range resultChan {
		availableIPs = append(availableIPs, ip)
	}

	return availableIPs
}

func findAvailableIps(ipRange []string, hostnames string, numWorkers int) []string {
	var availableIPs []string

	if len(ipRange) > 0 {
		availableIPs = append(availableIPs, scanIPRange(ipRange, numWorkers)...)
	}

	if hostnames != "" {
		for _, hostname := range strings.Split(hostnames, ",") {
			hostIPs, err := resolveHostnameToIP(hostname)
			if err == nil {
				availableIPs = append(availableIPs, hostIPs...)
			}
		}
	}

	return availableIPs
}

// TCP SCAN
func grabServiceBanner(targetAddress string, port int, network string) (string, error) {
	address := fmt.Sprintf("%s:%d", targetAddress, port)
	conn, err := net.Dial(network, address)
	if err != nil {
		return "", fmt.Errorf("failed to connect to %s: %w", address, err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	var banner []byte
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	for {
		n, err := conn.Read(buf)
		if n > 0 {
			banner = append(banner, buf[:n]...)
		}
		if err != nil {
			if n == 0 {
				break // no more data to read
			}
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				break // timed out
			}
			return "", fmt.Errorf("error reading from %s: %w", address, err)
		}
	}

	if len(banner) == 0 {
		return "", fmt.Errorf("no banner received from %s", address)
	}

	return string(banner), nil
}

func isPortOpen(targetAddress string, port int) bool {
	address := fmt.Sprintf("%s:%d", targetAddress, port)
	conn, err := net.DialTimeout("tcp", address, 2*DIAL_TIMEOUT*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func scanPorts(targetAddress string, numWorkers int, ports []int) []int {
	portChan := make(chan int, len(ports))
	results := make(chan int, len(ports))

	// start worker goroutines for port scanning
	for i := 0; i < numWorkers; i++ {
		go func() {
			for port := range portChan {
				if isPortOpen(targetAddress, port) {
					results <- port
				} else {
					results <- 0
				}
			}
		}()
	}

	// send ports to workers
	for _, port := range ports {
		portChan <- port
	}
	close(portChan)

	// collect the results
	openPorts := make([]int, 0, len(ports))
	for range ports {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	return openPorts
}

// UTILS
func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func parseInputPortRange(input string) ([]int, error) {
	var result []int
	if strings.Contains(input, "-") {
		ports := strings.Split(input, "-")
		start, err1 := strconv.Atoi(ports[0])
		end, err2 := strconv.Atoi(ports[1])
		if err1 != nil || err2 != nil || start > end {
			return nil, errors.New("invalid port range")
		}
		for i := start; i <= end; i++ {
			result = append(result, i)
		}
	} else if strings.Contains(input, ",") {
		for _, p := range strings.Split(input, ",") {
			port, err := strconv.Atoi(p)
			if err != nil {
				return nil, errors.New("invalid port number")
			}
			result = append(result, port)
		}
	} else {
		port, err := strconv.Atoi(input)
		if err != nil {
			return nil, errors.New("invalid port")
		}
		result = append(result, port)
	}
	return result, nil
}

func parseInputIpRange(input string) ([]string, error) {
	var result []string
	if strings.Contains(input, "/") {
		_, ipNet, err := net.ParseCIDR(input)
		if err != nil {
			return nil, fmt.Errorf("invalid CIDR notation: %v", err)
		}
		ip := ipNet.IP
		for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
			result = append(result, ip.String())
		}
	} else if strings.Contains(input, "-") {
		ips := strings.Split(input, "-")
		startIP := net.ParseIP(ips[0])
		endIP := net.ParseIP(ips[1])
		for ip := startIP; !startIP.Equal(endIP); incrementIP(ip) {
			result = append(result, ip.String())
		}
		result = append(result, endIP.String())
	} else if strings.Contains(input, ",") {
		for _, ip := range strings.Split(input, ",") {
			if net.ParseIP(ip) == nil {
				return nil, fmt.Errorf("invalid IP: %s", ip)
			}
			result = append(result, ip)
		}
	} else {
		result = append(result, input)
	}
	return result, nil
}

func printIpRange(ipRange []string) {
	if len(ipRange) > 0 {
		fmt.Println("Available IPs:")
		for _, ip := range ipRange {
			fmt.Println(ip)
		}
	} else {
		fmt.Println("No available IPs were found.")
	}
}

func printOpenPorts(ip string, openPorts []int, grabBanner bool) {
	var common = map[int]string{
		7:    "echo",
		20:   "ftp-data",
		21:   "ftp",
		22:   "ssh",
		23:   "telnet",
		25:   "smtp",
		43:   "whois",
		53:   "dns",
		67:   "dhcp-server",
		68:   "dhcp-client",
		80:   "http",
		110:  "pop3",
		111:  "rpcbind",
		123:  "ntp",
		137:  "netbios-ns",
		138:  "netbios-dgm",
		139:  "netbios-ssn",
		143:  "imap4",
		161:  "snmp",
		162:  "snmptrap",
		443:  "https",
		445:  "microsoft-ds",
		513:  "rlogin",
		554:  "rtsp",
		587:  "smtp",
		631:  "ipp",
		873:  "rsync",
		902:  "vmware",
		989:  "ftps-data",
		990:  "ftps",
		1194: "openvpn",
		1433: "mssql",
		3306: "mysql",
		3389: "rdp",
		8080: "http-alt",
		6379: "redis",
	}

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("|- %d %s\n", port, common[port])
		if grabBanner {
			fmt.Println(grabServiceBanner(ip, port, "tcp"))
		}
	}
}

// MAIN
func main() {
	ipRangeFlag := flag.String("r", "", "Specify a range or list of IPs (comma separated, CIDR notation, or dash separated range).")
	hostnamesFlag := flag.String("n", "", "Comma separated list of hostnames to scan.")
	portsFlag := flag.String("p", "0-1024", "Specify port(s) to scan (single, comma separated list, or hyphen separated range).")
	workerFlag := flag.Int("w", 200, "Number of concurrent worker goroutines.")
	grabFlag := flag.Bool("b", false, "Grab service banners from open ports.")

	flag.Parse()

	ipRange, err := parseInputIpRange(*ipRangeFlag)
	if err != nil {
		fmt.Println("Error parsing IP range:", err)
		return
	}

	ports, err := parseInputPortRange(*portsFlag)
	if err != nil {
		fmt.Println("Error parsing ports:", err)
		return
	}

	availableIPs := findAvailableIps(ipRange, *hostnamesFlag, *workerFlag)
	printIpRange(availableIPs)

	for _, ip := range availableIPs {
		fmt.Printf("\nScanning IP: %s\n", ip)
		openPorts := scanPorts(ip, *workerFlag, ports)
		printOpenPorts(ip, openPorts, *grabFlag)
	}
}
