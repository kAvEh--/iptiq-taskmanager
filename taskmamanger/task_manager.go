package taskmamanger

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func NewTaskManage(capacity int) *TaskManager {
	if capacity < 1 || capacity > 10000 {
		capacity = 10
	}
	return &TaskManager{
		MaxCapacity: capacity,
		ProcessList: make([]*MProcess, 0),
	}

}

func (tm *TaskManager) Add(process MProcess) error {
	if len(tm.ProcessList) >= tm.MaxCapacity {
		return errors.New("maximum capacity reached")
	}
	process.time = time.Now()
	tm.ProcessList = append(tm.ProcessList, &process)

	return nil
}

func (tm *TaskManager) AddFIFO(process MProcess) {
	if len(tm.ProcessList) >= tm.MaxCapacity {
		sort.Sort(ByTime(tm.ProcessList))
		tm.ProcessList[0].Process.Kill()
		tm.ProcessList = tm.ProcessList[1:]
	}
	process.time = time.Now()
	tm.ProcessList = append(tm.ProcessList, &process)
}

func (tm *TaskManager) AddPriority(process MProcess) error {
	process.time = time.Now()
	if len(tm.ProcessList) >= tm.MaxCapacity {
		indicator := 0
		for i := 1; i < len(tm.ProcessList); i++ {
			if tm.ProcessList[i].Priority < tm.ProcessList[indicator].Priority {
				indicator = i
			} else if tm.ProcessList[i].Priority == tm.ProcessList[indicator].Priority {
				if tm.ProcessList[i].time.Unix() == tm.ProcessList[indicator].time.Unix() {
					indicator = i
				}
			}
		}
		if tm.ProcessList[indicator].Priority < process.Priority {
			tm.ProcessList = append(tm.ProcessList[:indicator], tm.ProcessList[indicator+1:]...)
			tm.ProcessList = append(tm.ProcessList, &process)
			return nil
		}
		return errors.New("no process found with lower priority")
	}
	tm.ProcessList = append(tm.ProcessList, &process)
	return nil
}

func (tm *TaskManager) List(sorting string) []*MProcess {
	switch sorting {
	case "priority":
		sort.Sort(ByPriority(tm.ProcessList))
	case "id":
		sort.Sort(ByID(tm.ProcessList))
	case "time":
		sort.Sort(ByTime(tm.ProcessList))
	}
	return tm.ProcessList
}

func (tm *TaskManager) Kill(process MProcess) error {
	err := process.Process.Kill()
	if err != nil {
		return err
	}
	for i := 0; i < len(tm.ProcessList); i++ {
		if tm.ProcessList[i].Process.Pid == process.Process.Pid {
			tm.ProcessList = append(tm.ProcessList[:i], tm.ProcessList[i+1:]...)
			fmt.Println("process deleted")
		}
	}

	return nil
}

func (tm *TaskManager) KillByPriority(priority PriorityType) error {
	if !priority.IsValid() {
		return errors.New("priority is invalid")
	}
	for i := 0; i < len(tm.ProcessList); i++ {
		if tm.ProcessList[i].Priority == priority {
			err := tm.ProcessList[i].Process.Kill()
			if err != nil {
				return err
			}
			tm.ProcessList = append(tm.ProcessList[:i], tm.ProcessList[i+1:]...)
		}
	}
	fmt.Println("priority deleted")
	return nil
}

func (tm *TaskManager) KillAll() error {
	for i := 0; i < len(tm.ProcessList); i++ {
		err := tm.ProcessList[i].Process.Kill()
		if err != nil {
			return err
		}
	}
	tm.ProcessList = make([]*MProcess, 0)
	fmt.Println("all processes deleted")
	return nil
}
