package main

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

var (
	port = flag.Int("port", 22, "enter the port to find out what process is consuming it presently")
	pidonly = flag.String("pidonly", "y", "enter y or n for yes or no. y, default, limits the output to just the PID")
)

func main() {
	flag.Parse()

	
	// retrieve PID. This is the entire point of the program
	pidFromPort := findPidfromPort(*port)

	if *pidonly == "y" {
		fmt.Println(pidFromPort)

	} else if *pidonly == "n"  {
		nameFromPid := nameFromPid(pidFromPort)
		prettyOutput := fmt.Sprintf("Pid: %v and process name: %v", pidFromPort, nameFromPid)
		fmt.Println(prettyOutput)
	} else {
		fmt.Println("If you are going to define -pidonly, you must choose either y or n, nothing else")
	}




}

func findPidfromPort (port int) (pid int32) {
	a, _ := net.Connections("TCP")

	for _, resource := range a {
		switch resource.Laddr.IP {
		case "::1", ":::", "0.0.0.0", "127.0.0.1":
			if resource.Laddr.Port == uint32(port) {
				pid = resource.Pid
			}
		}
	}
	return
}

func nameFromPid (pid int32) (name string) {
	if pid == 0 {
		name = "no process found"
	} else {
		newPid, _ := process.NewProcess(pid)
		name, _ = newPid.Name()
	}
	return
}