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

List of functionality implemented in this repo:

-AddProcess: add a process to Task Manager`s list

-AddProcess with FIFO approach: add a process to Task Manager`s list with FIFO approach

-AddProcess Priority based approach: add a process to Task Manager`s list priority based

-ListProcesses: return list of all processes

-KillProcess: kill a process and remove from list

-KillProcess by Priority: kill all processes with specified priority

-KillAll Processes: kill all processes in Task Manager`s list