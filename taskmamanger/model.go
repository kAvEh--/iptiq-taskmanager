package taskmamanger

import (
	"os"
	"time"
)

type MProcess struct {
	Process  *os.Process
	Priority PriorityType
	time     time.Time
}

type TaskManager struct {
	ProcessList []*MProcess
}

type ITaskManager interface {
	Add(process MProcess)
	AddFIFO(process MProcess)
	AddPriority(process MProcess)
	List(sorting string) []*MProcess
	Kill(process MProcess)
	KillByPriority(priority PriorityType)
	KillAll()
}

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
	return s[i].Priority < s[j].Priority
}

type ByTime []*MProcess

func (s ByTime) Len() int      { return len(s) }
func (s ByTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByTime) Less(i, j int) bool {
	return s[i].time.Unix() < s[j].time.Unix()
}

type ByID []*MProcess

func (s ByID) Len() int      { return len(s) }
func (s ByID) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByID) Less(i, j int) bool {
	return s[i].Process.Pid < s[j].Process.Pid
}
