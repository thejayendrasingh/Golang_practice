package task

/*
	Task package to store task (in-memory for now) and their status
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type TaskStatus int

const (
	Created TaskStatus = iota + 1
	Scheduled
	Started
	Running
	Completed
	Failed
)

var taskStatusString = []string{"", "Created", "Scheduled", "Started", "Running", "Completed", "Failed"}

var taskMap = make(map[string]Task)
var mu sync.Mutex

func AddNewTask(taskName string, task Task) (bool, error) {
	if len(taskName) == 0 {
		return false, errors.New("Task Name is Empty")
	}
	if &task == nil {
		return false, fmt.Errorf("Task Object is Empty, taskName : %s", taskName)
	}
	mu.Lock()
	defer mu.Unlock()

	_, ok := taskMap[taskName]
	if ok {
		return false, fmt.Errorf("Task Name is Already Exist, taskName : %s", taskName)
	}
	task.Status = Scheduled
	taskMap[taskName] = task
	log.Println("Task is Scheduled : ", task.Name)
	return true, nil
}

func UpdateTaskCompletedAt(taskName string) (bool, error) {
	if len(taskName) == 0 {
		return false, errors.New("Task Name is Empty")
	}

	mu.Lock()
	defer mu.Unlock()
	task, ok := taskMap[taskName]
	if !ok {
		return false, fmt.Errorf("Task Do Not Exist : %s", taskName)
	}
	task.CompletedAt = time.Now()
	task.TotalTimeTakenInSeconds = int(task.CompletedAt.Sub(task.StartedAt).Seconds())
	taskMap[taskName] = task
	log.Printf("Task %s : started at %s, and competed at %s", task.Name, task.StartedAt, task.CompletedAt)
	return true, nil
}

func UpdateTaskStatus(taskName string, status TaskStatus) (bool, error) {
	if len(taskName) == 0 {
		return false, errors.New("Task Name is Empty")
	}
	if &status == nil {
		return false, fmt.Errorf("TaskStatus is Empty, taskName : %s", taskName)
	}

	mu.Lock()
	defer mu.Unlock()
	task, ok := taskMap[taskName]
	if !ok {
		return false, fmt.Errorf("Task Do Not Exist : %s", taskName)
	}
	task.Status = status
	taskMap[taskName] = task
	log.Printf("Task %s : Status changed to %s", task.Name, taskStatusString[status])
	return true, nil
}

func GetAllTasksInJson() ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()
	return json.Marshal(taskMap)
}

func GetTaskInJson(taskName string) ([]byte, error) {
	if len(taskName) == 0 {
		return nil, errors.New("Task Name is Empty")
	}

	mu.Lock()
	defer mu.Unlock()
	task, ok := taskMap[taskName]
	if !ok {
		return nil, fmt.Errorf("Task Do Not Exist : %s", taskName)
	}
	return json.Marshal(task)
}
