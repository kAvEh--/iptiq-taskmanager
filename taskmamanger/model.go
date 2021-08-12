package taskmamanger

import (
	"os"
	"time"
)

// MProcess represent an os Process with priority and creation time
type MProcess struct {
	Process  *os.Process
	Priority PriorityType
	time     time.Time
}

//TaskManager represent a list of tasks and maximum capacity of list
type TaskManager struct {
	MaxCapacity int
	ProcessList []*MProcess
}

//ITaskManager consists of all functions that must be implemented by Task Manager
type ITaskManager interface {
	Add(process MProcess) error
	AddFIFO(process MProcess)
	AddPriority(process MProcess) error
	List(sorting string) []*MProcess
	Kill(process MProcess) error
	KillByPriority(priority PriorityType) error
	KillAll() error
}

//PriorityType is an Int with validation function
type PriorityType int

func (p PriorityType) IsValid() bool {
	if p == 1 || p == 2 || p == 3 {
		return true
	}
	return false
}

type ByPriority []*MProcess

func (s ByPriority) Len() int      { return len(s) }
func (s ByPriority) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPriority) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].time.UnixNano() < s[j].time.UnixNano()
	}
	return s[i].Priority < s[j].Priority
}

type ByTime []*MProcess

func (s ByTime) Len() int      { return len(s) }
func (s ByTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByTime) Less(i, j int) bool {
	return s[i].time.UnixNano() < s[j].time.UnixNano()
}

type ByID []*MProcess

func (s ByID) Len() int      { return len(s) }
func (s ByID) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByID) Less(i, j int) bool {
	return s[i].Process.Pid < s[j].Process.Pid
}
