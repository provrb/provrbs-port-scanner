package connections

import (
	"net"
	argUtils "port-scanner/args"
	"strconv"
	"sync"
	"time"
)

var openPortsMutex sync.Mutex
var openPortsInt int = 0
var openPortsList []int

func ListOpenPorts() []int {
	return openPortsList
}

func PortsOpen() int {
	return openPortsInt
}

func IsPortOpen(info argUtils.Settings, port int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	config := net.Dialer{}

	if *info.Timeout != 0 {
		config = net.Dialer{Timeout: time.Duration(*info.Timeout) * time.Second} // timeout
	}

	address := *info.IP + ":" + strconv.Itoa(port) // ip + port

	_, err := config.Dial(*info.Protocol, address)
	if err == nil {
		openPortsMutex.Lock()
		openPortsInt++
		openPortsMutex.Unlock()
		if *info.ListOpenPorts {
			openPortsList = append(openPortsList, port)
		}
	}
}
