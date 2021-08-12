package taskmamanger

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type mTest struct {
	process  MProcess
	list     []*MProcess
	error    error
	sort     string
	priority int
}

func TestTaskManager_Add(t *testing.T) {
	tm := NewTaskManage(2)

	p1, _ := start("echo", ">>>>> t1")
	p2, _ := start("echo", ">>>>> t2")
	p3, _ := start("echo", ">>>>> t3")
	mp1 := MProcess{Process: p1, Priority: 1}
	mp2 := MProcess{Process: p2, Priority: 2}
	mp3 := MProcess{Process: p3, Priority: 3}

	tt := []mTest{
		{process: mp1, error: nil, list: []*MProcess{&mp1}},
		{process: mp2, error: nil, list: []*MProcess{&mp1, &mp2}},
		{process: mp3, error: errors.New("maximum capacity reached")},
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

func TestTaskManager_AddFIFO(t *testing.T) {
	tm := NewTaskManage(2)

	p1, _ := start("echo", ">>>>> t1")
	p2, _ := start("echo", ">>>>> t2")
	p3, _ := start("echo", ">>>>> t3")
	mp1 := MProcess{Process: p1, Priority: 1}
	mp2 := MProcess{Process: p2, Priority: 2}
	mp3 := MProcess{Process: p3, Priority: 3}

	tt := []mTest{
		{process: mp1, list: []*MProcess{&mp1}},
		{process: mp2, list: []*MProcess{&mp1, &mp2}},
		{process: mp3, list: []*MProcess{&mp2, &mp3}},
	}
	for _, test := range tt {
		testName := fmt.Sprintf("%d", test.process.Process.Pid)
		t.Run(testName, func(t *testing.T) {
			tm.AddFIFO(test.process)
			if !check(test.list, tm.ProcessList) {
				t.Errorf("got %v, want %v", tm.ProcessList, test.list)
			}
		})
	}
}

func TestTaskManager_AddPriority(t *testing.T) {
	tm := NewTaskManage(2)

	p1, _ := start("echo", ">>>>> t1")
	p2, _ := start("echo", ">>>>> t2")
	p3, _ := start("echo", ">>>>> t3")
	p4, _ := start("echo", ">>>>> t4")
	p5, _ := start("echo", ">>>>> t5")
	mp1 := MProcess{Process: p1, Priority: 1}
	mp2 := MProcess{Process: p2, Priority: 2}
	mp3 := MProcess{Process: p3, Priority: 2}
	mp4 := MProcess{Process: p4, Priority: 2}
	mp5 := MProcess{Process: p5, Priority: 3}

	tt := []mTest{
		{process: mp1, error: nil, list: []*MProcess{&mp1}},
		{process: mp2, error: nil, list: []*MProcess{&mp1, &mp2}},
		{process: mp3, error: nil, list: []*MProcess{&mp2, &mp3}},
		{process: mp4, error: errors.New("no process found with lower priority")},
		{process: mp5, error: nil, list: []*MProcess{&mp3, &mp5}},
	}
	for _, test := range tt {
		testName := fmt.Sprintf("%d", test.process.Process.Pid)
		t.Run(testName, func(t *testing.T) {
			err := tm.AddPriority(test.process)
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

func TestTaskManager_List(t *testing.T) {
	tm := NewTaskManage(5)

	p1, _ := start("echo", ">>>>> t1")
	p2, _ := start("echo", ">>>>> t2")
	p3, _ := start("echo", ">>>>> t3")
	p4, _ := start("echo", ">>>>> t4")
	p5, _ := start("echo", ">>>>> t5")
	mp1 := MProcess{Process: p1, Priority: 3}
	mp2 := MProcess{Process: p2, Priority: 2}
	mp3 := MProcess{Process: p3, Priority: 3}
	mp4 := MProcess{Process: p4, Priority: 2}
	mp5 := MProcess{Process: p5, Priority: 1}

	tm.AddFIFO(mp1)
	tm.AddFIFO(mp2)
	tm.AddFIFO(mp3)
	tm.AddFIFO(mp4)
	tm.AddFIFO(mp5)

	tt := []mTest{
		{sort: "priority", list: []*MProcess{&mp5, &mp2, &mp4, &mp1, &mp3}},
		{sort: "time", list: []*MProcess{&mp1, &mp2, &mp3, &mp4, &mp5}},
	}
	for _, test := range tt {
		testName := fmt.Sprintf("%s", test.sort)
		t.Run(testName, func(t *testing.T) {
			tmp := tm.List(test.sort)
			if !check(test.list, tmp) {
				t.Errorf("got %v, want %v", tm.ProcessList, test.list)
			}
		})
	}
}

func TestTaskManager_Kill(t *testing.T) {
	tm := NewTaskManage(5)

	p1, _ := start("echo", ">>>>> t1")
	p2, _ := start("echo", ">>>>> t2")
	p3, _ := start("echo", ">>>>> t3")
	p4, _ := start("echo", ">>>>> t4")
	p5, _ := start("echo", ">>>>> t5")
	mp1 := MProcess{Process: p1, Priority: 3}
	mp2 := MProcess{Process: p2, Priority: 2}
	mp3 := MProcess{Process: p3, Priority: 3}
	mp4 := MProcess{Process: p4, Priority: 2}
	mp5 := MProcess{Process: p5, Priority: 1}

	tm.AddFIFO(mp1)
	tm.AddFIFO(mp2)
	tm.AddFIFO(mp3)
	tm.AddFIFO(mp4)
	tm.AddFIFO(mp5)

	tt := []mTest{
		{process: mp1, list: []*MProcess{&mp2, &mp3, &mp4, &mp5}},
		{process: mp2, list: []*MProcess{&mp3, &mp4, &mp5}},
		{process: mp3, list: []*MProcess{&mp4, &mp5}},
		{process: mp4, list: []*MProcess{&mp5}},
		{process: mp5, list: []*MProcess{}},
	}
	for _, test := range tt {
		testName := fmt.Sprintf("%s", test.sort)
		t.Run(testName, func(t *testing.T) {
			_ = tm.Kill(test.process)
			if !check(test.list, tm.ProcessList) {
				t.Errorf("got %v, want %v", tm.ProcessList, test.list)
			}
		})
	}
}

func TestTaskManager_KillByPriority(t *testing.T) {
	tm := NewTaskManage(5)

	p1, _ := start("echo", ">>>>> t1")
	p2, _ := start("echo", ">>>>> t2")
	p3, _ := start("echo", ">>>>> t3")
	p4, _ := start("echo", ">>>>> t4")
	p5, _ := start("echo", ">>>>> t5")
	mp1 := MProcess{Process: p1, Priority: 3}
	mp2 := MProcess{Process: p2, Priority: 2}
	mp3 := MProcess{Process: p3, Priority: 3}
	mp4 := MProcess{Process: p4, Priority: 2}
	mp5 := MProcess{Process: p5, Priority: 1}

	tm.AddFIFO(mp1)
	tm.AddFIFO(mp2)
	tm.AddFIFO(mp3)
	tm.AddFIFO(mp4)
	tm.AddFIFO(mp5)

	tt := []mTest{
		{priority: 1, list: []*MProcess{&mp1, &mp2, &mp3, &mp4}},
		{priority: 2, list: []*MProcess{&mp1, &mp3}},
		{priority: 3, list: []*MProcess{}},
	}
	for _, test := range tt {
		testName := fmt.Sprintf("%d", test.priority)
		t.Run(testName, func(t *testing.T) {
			_ = tm.KillByPriority(PriorityType(test.priority))
			if !check(test.list, tm.ProcessList) {
				t.Errorf("got %v, want %v", tm.ProcessList, test.list)
			}
		})
	}
}

func TestTaskManager_KillAll(t *testing.T) {
	tm := NewTaskManage(5)

	p1, _ := start("echo", ">>>>> t1")
	p2, _ := start("echo", ">>>>> t2")
	p3, _ := start("echo", ">>>>> t3")
	p4, _ := start("echo", ">>>>> t4")
	p5, _ := start("echo", ">>>>> t5")
	mp1 := MProcess{Process: p1, Priority: 3}
	mp2 := MProcess{Process: p2, Priority: 2}
	mp3 := MProcess{Process: p3, Priority: 3}
	mp4 := MProcess{Process: p4, Priority: 2}
	mp5 := MProcess{Process: p5, Priority: 1}

	tm.AddFIFO(mp1)
	tm.AddFIFO(mp2)
	tm.AddFIFO(mp3)
	tm.AddFIFO(mp4)
	tm.AddFIFO(mp5)

	tt := []mTest{
		{error: nil},
	}
	for _, test := range tt {
		testName := fmt.Sprintf("%d", test.priority)
		t.Run(testName, func(t *testing.T) {
			err := tm.KillAll()
			if err != nil {
				t.Errorf("got error %s", err.Error())
			}
			if !check(test.list, tm.ProcessList) {
				t.Errorf("got %v, want %v", tm.ProcessList, test.list)
			}
		})
	}
}

func check(l1 []*MProcess, l2 []*MProcess) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
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
