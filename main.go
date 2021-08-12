package main

import (
	"fmt"
	"github.com/kAvEh--/iptiq-taskmanager/taskmamanger"
	"os"
	"os/exec"
)

func main() {
	tm := taskmamanger.NewTaskManage(2)

	p1, _ := start("ping", "-c 1", "www.google.com")
	p2, _ := start("ping", "-c 1", "www.dell.com")
	p3, _ := start("ping", "-c 1", "www.microsoft.com")
	mp1 := taskmamanger.MProcess{
		Process:  p1,
		Priority: 1,
	}
	tm.Add(mp1)
	tm.Add(taskmamanger.MProcess{
		Process:  p3,
		Priority: 2,
	})
	e := tm.Add(taskmamanger.MProcess{
		Process:  p2,
		Priority: 3,
	})
	if e != nil {
		fmt.Println(e)
	}
}

func start(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}
