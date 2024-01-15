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

### Why the ports scanning is important in the pentesting?

### How this program works?

This program reads commandline arguments (like -fn, -u, -ip) and provided argument for which we need to investigate on.
For full name it does web scraping from www.whitepages.be, ip address it uses http://ip-api.com and username it checks if that platform userpage exists or not.
