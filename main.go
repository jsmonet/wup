package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"

)

var (
	port = flag.IntP("port", "p", 22, "which port")
	//includeName = pflag.StringP("name", "n", "n", "include process name?")
	verbose = flag.BoolP("verbose", "v", false, "include the process name in a pretty formatted string?") // simply adding the -v flag with no arguments makes this true

)

func main() {


	flag.Parse()

	// retrieve PID. This is the entire point of the program
	pidFromPort := findPidfromPort(*port)
	response := fmt.Sprint(pidFromPort) // just stringing for type consistency below

	if *verbose {
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