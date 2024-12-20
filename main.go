package main

import (
	"fmt"
	argutils "port-scanner/args"
	connections "port-scanner/connections"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

func main() {
	var args argutils.Settings = argutils.ParseArgs() // scanning args

	fmt.Println("Scanning ports", *args.StartPort, "to", *args.EndPort)

	var waitGroup sync.WaitGroup         // wait group for sync
	var workers int = 0                  // amount of threads spawned
	var startTime time.Time = time.Now() // time the scan started

	bar := progressbar.Default(int64(*args.EndPort - *args.StartPort))

	for i := *args.StartPort; i < *args.EndPort; i++ {
		if workers%100 == 0 { // prevent "queue full" error with net.Dial. spawn 100 workers at a time
			time.Sleep(9 * time.Millisecond)
		}
		workers++
		waitGroup.Add(1)
		bar.Add(1)
		go connections.IsPortOpen(args, i, &waitGroup)
	}

	waitGroup.Wait()
	var endTime time.Time = time.Now() // time at which the scan completed

	fmt.Println("Scan complete. Took", endTime.UnixMilli()-startTime.UnixMilli(), "ms")
	fmt.Println("Found", connections.PortsOpen(), *args.Protocol, "open ports")
	if !*args.ListOpenPorts {
		return
	}

	fmt.Print("OPEN Ports: ")
	for i, v := range connections.ListOpenPorts() {
		if i%10 != 0 {
			fmt.Printf("%-7d", v)
		} else {
			fmt.Print("\n")
		}
		fmt.Print(" ")
	}
}
