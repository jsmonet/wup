# wup is life

What's Using [this] Port is highly lossy, oversimplified, and bears rework.

## SECURITY!
This is a day old and has never been tested for vulnerabilities. I have no intention of doing any such analysis while I'm still hashing out the first iteration. 

## INSTALL!

`go install github.com/jsmonet/wup`

This is assuming you have go installed on your system. If not... install it and then run

`go install github.com/jsmonet/wup`

## USE!

```
wup -p 8080
```

This runs the program and asks it to return the PID of what is serving on port 8080.

```
wup -p 8080 -n y
```

This returns a pretty formatted string telling you the PID and the process name consuming port 8080.

Yeah, I know `-n y` is horrible. I'm blanking, at present, on what bare flag I want to use for this.

This currently works on my laptop and has not been tested anywhere else. 'Nix support coming soon. Windows has been brought up to me, so we'll see how that works out.

## Under the hood

This app uses Shirou's gopsutil package. It queries what network connections are active locally and filters for anything claiming a local IP address of "::1", ":::", "0.0.0.0", or "127.0.0.1". It then tries to match your port argument to connections matching both those local IP addresses as well as the port you specify. As the connection info includes a PID, it then returns the related PID. 

You can optionally request the process name based on the PID. This can be a bit lossy, but is generally useful. Process names are derived using the gopsutil/process bundle. 

## Where's the testing and validation??

Coming. 

## I have a suggestion

Comment here, raise an issue, and/or hit me up on slack