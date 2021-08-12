# iptiq-taskmanager

Task Manager written in Go.

Task Manager will handle multiple processes inside an operating system.

### Requirements

Go 1.8+: download and install Go here: https://golang.org/doc/install

### Usage

For building the code use this script:

```shell
go build 
```

For running the code use this script:

```shell
go build & ./iptiq-taskmanager
```

For running tests use this script:

```shell
go test -v ./.. 
```

### Implementation
-AddProcess

-AddProcess with FIFO approach

-AddProcess Priority based approach

-ListProcesses

-KillProcess

-KillProcess by Priority

-KillAll Processes