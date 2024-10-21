package routes

/*
	Routes package to handle all routes defined on this go_server
	- It maps url path to corresponding handler func
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/executor"
	"server/resource"
	"server/task"
)

type AppStats struct {
	TotalTaskCount           int
	TotalResourceATasks      int
	TotalResourceBTasks      int
	TotalResourceUnkownCount int
	TotalExecutionOnlyTask   int
	ResourceAPoolSize        int
	ResourceBPoolSize        int
}

func StartServer() {
	http.HandleFunc("/createTask", createTask)
	http.HandleFunc("/allTasks", getAllTasks)
	http.HandleFunc("/allStats", getAllStats)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}

func getAllStats(w http.ResponseWriter, _ *http.Request) {
	var appStat = AppStats{
		TotalTaskCount:           executor.TotalTaskCount,
		TotalResourceATasks:      executor.TotalResourceATasks,
		TotalResourceBTasks:      executor.TotalResourceBTasks,
		TotalResourceUnkownCount: executor.TotalResourceUnkownCount,
		TotalExecutionOnlyTask:   executor.TotalExecutionOnlyTask,
		ResourceAPoolSize:        resource.GetTotalResourceACount(),
		ResourceBPoolSize:        resource.GetTotalResourceBCount(),
	}

	statsJson, err := json.Marshal(appStat)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(statsJson)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create task request received")
	decoder := json.NewDecoder(r.Body)
	var t task.Task
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	t = task.CreateNewTask(t.Name, t.ProcessingTimeInSeconds, t.WillUseResource, t.WhichResource)
	log.Println(t)
	executor.Execute(t)

	userJson, err := task.GetTaskInJson(t.Name)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func getAllTasks(w http.ResponseWriter, _ *http.Request) {
	allStatus, err := task.GetAllTasksInJson()
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Failed to return response"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(allStatus)
	}
}
