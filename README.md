## General

This is tool that gathers information from open source investigative methods.

## Setup

1. Check if python3 is installed
   python3 --version
   If NOT(ubuntu):
   sudo apt update
   sudo apt install python3

## Commands

### go run main.go --help

Usage: tinyscanner [OPTIONS] [HOST] [PORT]
Options:
-p Range of ports to scan
-u UDP scan
-t TCP scan
--help Show this message and exit.

### go run main.go -t 127.0.0.1 -p 80

Does port 80 show as open?

### Run a local server using udp protocol with the port 5500 and run tinyscanner -t 127.0.0.1 -p 5500

Does port 80 show as open?

### go run main.go -t 127.0.0.1 -p 1604

Does it show the following?
Port 1604 is closed

### go run main.go -t 10.53.224.5 -p 80-83

Does it display this result?
Port 80 is open
Port 81 is open
Port 82 is close
Port 83 is open

## Audit questions:

https://github.com/01-edu/public/tree/master/subjects/cybersecurity/passive/audit

### Explain clearly what does port mean?

Ports are virtual places within an operating system where network connections start and end. They help computers sort the network traffic they receive.

### Why the ports scanning is important in the pentesting?

The overarching goal of port scanning is to find ports, hosts, and server locations vulnerable to an attack and to diagnose those points' security levels. Malicious hackers also use the same process to identify port security weaknesses that they can use to compromise a network's security and gain access

### How this program works?

This program reads commandline arguments (like -t, -u, -p) and provided port we need to scan.
The ports are scanned using udp or tcp scan. Tcp scan gives close result if no tcp connection or handshake is
established between the server. Udp tries to provoke answer from the server by sending a small package and if it doesnt receive anything it assumes the server is closed.
