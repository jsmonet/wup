package main

import (
	"flag"
	"github.com/shirou/gopsutil/net"
	"fmt"
	"github.com/shirou/gopsutil/process"
)

var (
	port = flag.Int("port", 22, "enter the port to find out what process is consuming it presently")
	//pidonly = flag.Bool("pidonly", true, "bool, true returns just the pid, false returns the pid as well as the process name"
)

func main() {
	flag.Parse()

	pidFromPort := findPidfromPort(*port)
	nameFromPid := nameFromPid(pidFromPort)
	prettyOutput := fmt.Sprintf("Pid: %v and process name: %v", pidFromPort, nameFromPid)
	fmt.Println(prettyOutput)


}

func findPidfromPort (port int) (pid int32) {
	a, _ := net.Connections("TCP")
	//fmt.Println(a)
	for _, resource := range a {

		switch resource.Laddr.IP {
		case "::1", ":::", "0.0.0.0", "127.0.0.1":
			if resource.Laddr.Port == uint32(port) {
				pid = resource.Pid
			}
		default:
			pid = 0
		}

		//if resource.Laddr.IP == "::1" {
		//	if resource.Laddr.Port == uint32(port) {
		//		pid = resource.Pid
		//	}
		//}
		//if resource.Laddr.IP == "127.0.0.1" {
		//	if resource.Laddr.Port == uint32(port) {
		//		pid = resource.Pid
		//	}
		//}
		//if resource.Laddr.IP == "0.0.0.0" {
		//	if resource.Laddr.Port == uint32(port) {
		//		pid = resource.Pid
		//	}
		//}
	}
	return
}

func nameFromPid (pid int32) (name string) {
	newPid, _ := process.NewProcess(pid)
	name, _ = newPid.Name()
	return
}