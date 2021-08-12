package taskmamanger

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type mTest struct {
	process MProcess
	list    []*MProcess
	error   error
}

func TestTaskManager_Add(t *testing.T) {
	tm := NewTaskManage(2)

	p1, _ := start("ls")
	mp1 := MProcess{Process: p1, Priority: 1}
	p2, _ := start("ls")
	mp2 := MProcess{Process: p2, Priority: 2}
	p3, _ := start("ls")
	mp3 := MProcess{Process: p3, Priority: 3}

	tt := []mTest{
		{process: mp1, error: nil, list: []*MProcess{&mp1}},
		{process: mp2, error: nil, list: []*MProcess{&mp1, &mp2}},
		{process: mp3, error: errors.New("maximum capacity reached"), list: []*MProcess{&mp1, &mp2}},
	}
	for _, test := range tt {
		testName := fmt.Sprintf("%d", test.process.Process.Pid)
		t.Run(testName, func(t *testing.T) {
			err := tm.Add(test.process)
			if err != nil {
				if test.error != nil {
					if err.Error() != test.error.Error() {
						t.Errorf("got %s, want %s", err, test.error)
					}
				} else {
					t.Errorf("got %s, want %s", err, test.error)
				}
			} else {
				if !check(test.list, tm.ProcessList) {
					t.Errorf("got %v, want %v", tm.ProcessList, test.list)
				}
			}
		})
	}
}

func check(l1 []*MProcess, l2 []*MProcess) bool {
	fmt.Println("*****", len(l1), len(l2))
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		fmt.Println("----", l1[i].Process.Pid, l2[i].Process.Pid)
		if l1[i].Process.Pid != l2[i].Process.Pid {
			return false
		}
	}

	return true
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
