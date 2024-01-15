package main

import (
	"flag"
	"fmt"
)

func main()  {
	tcp := flag.Bool("t", false, "TCP scan")
	udp := flag.Bool("u", false, "UDP scan")
	portRange := flag.String("p", "", "Range of ports to scan")
	flag.Parse()

	scanType := ""

	if *tcp  {
		scanType = "tcp"
		fmt.Println(scanPort(scanType, *portRange))
	} else if *udp {
		scanType = "udp"
		fmt.Println("port", &portRange)
		fmt.Println(scanPort(scanType, *portRange))
	} else {
		fmt.Println("Usage: tinyscanner [OPTIONS] [HOST] [PORT]")
		fmt.Println("Options:")
		fmt.Println("  -p               Range of ports to scan")
		fmt.Println("  -u               UDP scan")
		fmt.Println("  -t               TCP scan")
		fmt.Println("  --help           Show this message and exit.")
	}
}

func scanPort(host string, port string) string {
    return "open"
}