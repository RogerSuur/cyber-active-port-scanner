package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
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


		
	startPort, endPort, err := findRange(*portRange)
	if err != nil {
		fmt.Println("Error parsing port range:", err)
		return
	}
		
	for port := startPort; port <= endPort; port++ {
        if scanType == "tcp" {
            fmt.Printf("Scanning TCP port %d: %s\n", port, scanTcp(*tcpHost, strconv.Itoa(port)))
        } else  {
            fmt.Printf("Scanning UDP port %d: %s\n", port, scanUdp(*udpHost, port))
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

func scanTcp(host string, port string) string {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
	return "closed"
	}
	conn.Close()
	return "open"
}


func scanUdp(host string, port int) string {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("udp", address, 1*time.Second)
	if err != nil {
		return "closed"
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		return "closed"
	}

	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	_, err = conn.Read(buffer)
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return "open/filtered"
	}

	return "closed"
}
