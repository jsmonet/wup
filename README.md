# wup is life

What's Using [this] Port gives you your choice of 
- the PID of an application using the port you specify
- the PID and name of that application
- a list of all services advertising locally, formatted to include the PID, Name, Port, and local address of each service

## SECURITY!
This has not been tested for security issues. Treat it as volatile. In fact, don't trust me. Run this sandboxed if you have even a hint of paranoia. I don't claim any responsibility for exploits of third-party libraries I've used, so please do your due diligence before putting this anywhere sensative. 

## INSTALL!

`go install github.com/jsmonet/wup`

This is assuming you have go installed on your system. If not... install it and then run

`go install github.com/jsmonet/wup`

## USE!

Here is some sample output, obviously sanitized a little bit:

```bash
bash-3.2$ wup -p 8080
1498
bash-3.2$ wup -p 8080 -v
1498 - com.docker.vpnki
bash-3.2$ wup -a
PID: 1498 Name: com.docker.charz Port: 9323 Local: ::1
PID: 1498 Name: com.docker.axe Port: 50000 Local: ::1
PID: 1498 Name: com.docker.skz Port: 8080 Local: ::1
PID: 1636 Name: Odds Port: 16123 Local: 127.0.0.1
PID: 1656 Name: Evens Port: 16223 Local: 127.0.0.1

```

If you're ever stuck, just run `wup -h` for the built-in help output:

```bash
Usage of wup:
     -a, --all        return all locally-served ports, their pids, and process names. Note: overrides -p
     -p, --port int   which port (default 22)
     -v, --verbose    include the process name in a pretty formatted string?  
```

## Tested in...

OSX, Centos7, very lightly in Debian9. More to come

## Under the hood

This app uses Shirou's gopsutil package. It queries what network connections are active locally and filters for anything claiming a local IP address of "::1", "::", ":::", "0.0.0.0", or "127.0.0.1". It then tries to match your port argument to connections matching both those local IP addresses as well as the port you specify. As the connection info includes a PID, it then returns the related PID. 

You can optionally request the process name based on the PID. This can be a bit lossy, but is generally useful. Process names are derived using the gopsutil/process bundle. 

## Where's the testing and validation??

Still Coming.

## I have a suggestion

Comment here, raise an issue, and/or hit me up on slack

### Changelog

v1.02: added `-a` flag to return all locally-served processes, their pids/names/ports/sources. 

v1.01: language update

v1: initial release