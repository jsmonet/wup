# wup is life

What's Using [this] Port is highly lossy, oversimplified, and bears rework. It will likely be responsible for your next 3 data breaches because I've never bothered trying to make sure it is secure.

That said, all it's supposed to do is ask you for a TCP port, and give you a process ID (PID) in response--and the name of the process if you ask it nicely.

Currently the program works on my laptop. Currently it hasn't been tested anywhere else. I feel like it deserves the "works on my machine" badge of honor presently, but I will work on shedding that designation shortly.

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

