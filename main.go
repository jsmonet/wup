package main

import (
	"flag"
	"github.com/shirou/gopsutil/net"
	"fmt"
)

var (
	port = flag.Int("port", 22, "enter the port to find out what process is consuming it presently")
	//pidonly = flag.Bool("pidonly", true, "bool, true returns just the pid, false returns the pid as well as the process name"
)

func main() {
	flag.Parse()

	pidFromPort := findPidfromPort(*port)
	fmt.Println(pidFromPort)

}

func findPidfromPort (port int) (pid int32) {
	a, _ := net.Connections("TCP")
	//fmt.Println(a)
	for _, resource := range a {
		if resource.Laddr.IP == "::1" {
			if resource.Laddr.Port == uint32(port) {
				pid = resource.Pid
			}
		}
		if resource.Laddr.IP == "127.0.0.1" {
			if resource.Laddr.Port == uint32(port) {
				pid = resource.Pid
			}
		}
	}
	return
}