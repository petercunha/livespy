# livespy
LiveSpy is an application that monitors a users computer and sends back screenshots, process lists, window titles, and other information to a sneaky hacker every three seconds. This is a project I'm writing to learn websockets in Go.

*This project is currently under development for Macs and Linux machines only.*

-

How to compile:
```shell
git clone https://github.com/petercunha/livespy/
cd livespy/ && go build main.go
./main
```

Once livespy is running, visit `http://localhost:8080` in your favorite web browser.
