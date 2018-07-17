package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"

)

var (
	port = flag.IntP("port", "p", 22, "which port")
	verbose = flag.BoolP("verbose", "v", false, "include the process name in a pretty formatted string?") // simply adding the -v flag with no arguments makes this true
	allLocal = flag.BoolP("all", "a", false, "return all locally-served ports, their pids, and process names. Note: overrides -p")
)

func main() {


	flag.Parse()

	if *allLocal {
		getLocalPnP()
	} else {
		// retrieve PID. This is the entire point of the program
		pidFromPort := findPidfromPort(*port)
		response := fmt.Sprint(pidFromPort) // just stringing for type consistency below

		if *verbose {
			procName := nameFromPid(pidFromPort)
			response = fmt.Sprintf("%v - %v", pidFromPort, procName)
		}

		fmt.Println(response)
	}
}

func findPidfromPort (port int) (pid int32) {
	a, _ := net.Connections("tcp")

	for _, resource := range a {
		switch resource.Laddr.IP {
		case "::1", "::", ":::", "0.0.0.0", "127.0.0.1": // there's no place like home, home, home, home, or home
			if resource.Laddr.Port == uint32(port) {
				pid = resource.Pid
			}
		}
	}
	return
}

func getLocalPnP () {
	a, _ := net.Connections("tcp")

	for _, datas := range a {
		switch datas.Laddr.IP {
		case "::1", "::", ":::", "0.0.0.0", "127.0.0.1":
			name := nameFromPid(datas.Pid)
			fmt.Printf("PID: %v Name: %v Port: %v Local: %v\n", datas.Pid, name, datas.Laddr.Port, datas.Laddr.IP)
		}
	}

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
