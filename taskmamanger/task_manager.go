package taskmamanger

import (
	"sort"
	"time"
)

func NewTaskManage() *TaskManager {
	return &TaskManager{
		ProcessList: make([]*MProcess, 0),
	}

}

func (tm *TaskManager) Add(process MProcess) {
	process.time = time.Now()
	tm.ProcessList = append(tm.ProcessList, &process)
}

func (tm *TaskManager) AddFIFO(process MProcess) {

}

func (tm *TaskManager) AddPriority(process MProcess) {

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

func (tm *TaskManager) Kill(process MProcess) {

}

func (tm *TaskManager) KillByPriority(priority PriorityType) {

}

func (tm *TaskManager) KillAll() {

}
