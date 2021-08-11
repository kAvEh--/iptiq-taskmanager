package taskmamanger

import (
	"os"
	"time"
)

type MProcess struct {
	Process  *os.Process
	Priority PriorityType
	Time     time.Time
}

type TaskManager struct {
}

type ITaskManager interface {
	Add(process MProcess)
	AddFIFO(process MProcess)
	AddPriority(process MProcess)
	List(sort int)
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
