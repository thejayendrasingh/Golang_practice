package executor

/*
	This is Executor package to execute http request
	- It simulates an orchestartor that fetch resource if required
	- Put goroutine to sleep to simulate executin time or resource usage time
	- It also keeps status udpated in central storage of all tasks
*/

import (
	"fmt"
	"log"
	"server/resource"
	"server/task"
	"time"
)

var TotalTaskCount int
var TotalResourceATasks int
var TotalResourceBTasks int
var TotalResourceUnkownCount int
var TotalExecutionOnlyTask int

func Execute(t task.Task) error {
	TotalTaskCount++
	log.Println("Task added to status map : ", t.Name)
	_, err := task.AddNewTask(t.Name, t)
	if err != nil {
		return err
	}
	if t.WillUseResource {
		log.Println("Task wants to use some of limited resource : ", t.Name)
		if t.WhichResource == 1 {
			TotalResourceATasks++
			log.Println("Task wants to use resource 1 : ", t.Name)
			task.UpdateTaskStatus(t.Name, task.Started)
			var res = resource.FetchFromResourceA()
			task.UpdateTaskStatus(t.Name, task.Running)
			time.Sleep(time.Duration(t.ProcessingTimeInSeconds) * time.Second)
			resource.ReturnToResourceA(&res)
			task.UpdateTaskCompletedAt(t.Name)
			task.UpdateTaskStatus(t.Name, task.Completed)
		} else if t.WhichResource == 2 {
			TotalResourceBTasks++
			log.Println("Task wants to use resource 2 : ", t.Name)
			task.UpdateTaskStatus(t.Name, task.Started)
			var res = resource.FetchFromResourceB()
			task.UpdateTaskStatus(t.Name, task.Running)
			time.Sleep(time.Duration(t.ProcessingTimeInSeconds) * time.Second)
			resource.ReturnToResourceB(&res)
			task.UpdateTaskCompletedAt(t.Name)
			task.UpdateTaskStatus(t.Name, task.Completed)
		} else {
			TotalResourceUnkownCount++
			log.Println("Task wants to use resource <UNKOWN> : ", t.Name)
			task.UpdateTaskCompletedAt(t.Name)
			task.UpdateTaskStatus(t.Name, task.Failed)
			return fmt.Errorf("unkown resource, %s ", t.Name)
		}
	} else {
		TotalExecutionOnlyTask++
		log.Println("Task wants to just execute : ", t.Name)
		task.UpdateTaskStatus(t.Name, task.Started)
		task.UpdateTaskStatus(t.Name, task.Running)
		time.Sleep(time.Duration(t.ProcessingTimeInSeconds) * time.Second)
		task.UpdateTaskCompletedAt(t.Name)
		task.UpdateTaskStatus(t.Name, task.Completed)
	}
	return nil
}