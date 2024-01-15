package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	tcpHost := flag.String("t", "", "TCP scan for specified host")
    udpHost := flag.String("u", "", "UDP scan for specified host")
	portRange := flag.String("p", "", "Range of ports to scan")
	flag.Parse()

	scanType := ""

	if *tcpHost != "" {
		scanType = "tcp"
	} else if *udpHost != ""{
		scanType = "udp"
		} else {
			fmt.Println("Usage: tinyscanner [OPTIONS] [HOST] [PORT]")
			fmt.Println("Options:")
			fmt.Println("  -p               Range of ports to scan")
			fmt.Println("  -u               UDP scan")
			fmt.Println("  -t               TCP scan")
			fmt.Println("  --help           Show this message and exit.")
		}

		fmt.Printf("Debug: Port range input is '%s'\n", *portRange)

		
	startPort, endPort, err := findRange(*portRange)
	if err != nil {
		fmt.Println("Error parsing port range:", err)
		return
	}
		
	for port := startPort; port <= endPort; port++ {
        if scanType == "tcp" {
            fmt.Printf("Scanning TCP port %d: %s\n", port, scanPort("tcp", strconv.Itoa(port)))
        } else  {
            fmt.Printf("Scanning UDP port %d: %s\n", port, scanPort("udp", strconv.Itoa(port)))
        }
    }
}

func findRange(portRange string)(int, int, error) {
	if strings.Contains(portRange, "-") {
        parts := strings.Split(portRange, "-")
        if len(parts) != 2 {
            return 0, 0, fmt.Errorf("invalid port range format")
        }
        startPort, err := strconv.Atoi(parts[0])
        if err != nil {
            return 0, 0, fmt.Errorf("invalid start port")
        }
        endPort, err := strconv.Atoi(parts[1])
        if err != nil {
            return 0, 0, fmt.Errorf("invalid end port")
        }
        return startPort, endPort, nil
	} else {
		port, err := strconv.Atoi(portRange)
        if err != nil {
			return  0,0, fmt.Errorf("invalid port")
        }
		return port, port, nil
	}
}

func scanPort(host string, port string) string {
    return "open"
}