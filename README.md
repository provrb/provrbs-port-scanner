# Provrbs Port Scanner

A cross-platform port scanner written in Go, compatible with both Windows and Linux machines for the command line. It scans a variety of ports to determine if they're open or closed, and supports multiple network protocols such as TCP, UDP, and IP.

## Installation

You can easily download the pre-built binaries for Windows and Linux from the [Releases Page](https://github.com/provrb/provrbs-port-scanner/releases)

1. Go to the Provrbs Port Scanner [Releases page](https://github.com/provrb/provrbs-port-scanner/releases).
2. Download the appropriate binary for your operating system:

Once downloaded, follow these steps to get started:
### For Windows:

1. Download ```pps.exe```.
2. Open a command prompt and navigate to the directory where pps.exe is located.
3. Run the program using the appropriate command, for example:
```
pps.exe -start 10 -end 3102 -protocol tcp -retrycon 3 -print true -timeout 3
```

### For Linux:

1. Download the Linux binary ```pps```.
2. Make the binary executable with the following command:
```
chmod +x pps
```
3. Run the program using the appropriate command, for example:
```
./pps -end 3102 -protocol udp -timeout 3
```
## Usage
### Example Commands:
**Windows:**
```
pps.exe -start 10 -end 3102 -protocol tcp -retrycon 3 -print true -timeout 3
```
**Linux:**
```
$ pps -end 3102 -protocol udp -timeout 3
```
### Command-Line Arguments
```
-start
```
Optional. Specifies the starting port to scan in a range. Default: 0.
```
-end
```
Optional. Specifies the ending port to scan in a range. Default: 65535.
```
-timeout
```
Optional. The number of seconds to wait before determining if a port is closed. Default: 3.
```
-protocol
```
Optional. Specifies the protocol to use when scanning for open ports. Options: tcp, udp, ip. Default: tcp.
```
-ip
```
Optional. The IP address to scan for open ports. Default: localhost (127.0.0.1).
```
-delay
```
Optional. The number of seconds to wait before checking the next port. Default: 0.
```
-print
```
Optional. If set to true, the program will print all open ports found. Default: false.
```
-retrycon
```
Optional. The number of times to retry a closed port before determining it's truly closed. Default: 1.