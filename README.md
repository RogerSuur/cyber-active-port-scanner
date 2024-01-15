## General

This is tool that gathers information from open source investigative methods.

## Setup

1. Check if python3 is installed
   python3 --version
   If NOT(ubuntu):
   sudo apt update
   sudo apt install python3

## Commands

### tinyscanner --help

Usage: tinyscanner [OPTIONS] [HOST] [PORT]
Options:
-p Range of ports to scan
-u UDP scan
-t TCP scan
--help Show this message and exit.

### tinyscanner -p 127.0.0.1 -t 80

Does port 80 show as open?

### inyscanner -p 127.0.0.1 -u 80

Does port 80 show as open?

## Audit questions:

https://github.com/01-edu/public/tree/master/subjects/cybersecurity/passive/audit

### Explain clearly what does port mean?

### Why the ports scanning is important in the pentesting?

### How this program works?

This program reads commandline arguments (like -fn, -u, -ip) and provided argument for which we need to investigate on.
For full name it does web scraping from www.whitepages.be, ip address it uses http://ip-api.com and username it checks if that platform userpage exists or not.
