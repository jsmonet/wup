package main

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

var (
	port = flag.Int("p", 22, "enter the port to find out what process is consuming it presently")
	includeName = flag.String("n", "n", "enter y or n for yes or no. y, default, limits the output to just the PID") // this is probably stupid

)

func main() {
	flag.Parse()

	// retrieve PID. This is the entire point of the program
	pidFromPort := findPidfromPort(*port)
	response := fmt.Sprint(pidFromPort) // just stringing for type consistency below

	if *includeName == "y" {
		procName := nameFromPid(pidFromPort)
		response = fmt.Sprintf("%v is consuming port %v", procName, pidFromPort)
	}

	fmt.Println(response)

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