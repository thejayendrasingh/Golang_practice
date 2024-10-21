package main

/*
	This utility auto-generate http request traffic and hit go server running on http://localhost:8080
	- during execution will start with askign how many request you want to send to go server, Enter number
	- It auto-generate inputs using sequential task name with random seconds of execution request should take
		with random request execution type ( normal sleep, using resource A(or 1) or resouce B (or 2)
	- All request will be created with 2 seconds delay after every 100 requests, due to avoid connection reset error
*/
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"sync"
	"syscall"
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

type Task struct {
	Name                    string
	ProcessingTimeInSeconds int
	WillUseResource         bool
	WhichResource           int
	StartedAt               time.Time
	CompletedAt             time.Time
	TotalTimeTakenInSeconds int
	Status                  TaskStatus
}
var wg sync.WaitGroup
func main(){
	var requestCount int
	fmt.Println("Enter number of random request to generate : ")
	fmt.Scan(&requestCount)	
	wg.Add(requestCount)
	startTime := time.Now()
	for inx := 0; inx < requestCount; inx++ {
		name := "Task-" + strconv.Itoa(inx)
		 go sendCreateRequest(name, rand.IntN(10)+1, (rand.IntN(2)+1)%2 == 0, rand.IntN(2)+1)
		if inx%100 == 0 {
			time.Sleep(2*time.Second)
		}
	}
	wg.Wait()
	endTime := time.Now()
	fmt.Println(fmt.Sprintf("Time Taken to execute %d random request( < 9s each) : %d seconds", requestCount, int(endTime.Sub(startTime).Seconds())))
	printAllAppStats()
}

// Fetch all statistics from go_server
func printAllAppStats(){
	getUrl := "http://localhost:8080/allStats"
	res, err := http.Get(getUrl)
	if err != nil {
		log.Println("New request creation error : ", err)
		return
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("io.ReadAll error",err)
			return
		}
		bodyString := string(bodyBytes)
		fmt.Println("All stats from go_server : ",bodyString)
	}
}

// Send create task request to go_server
func sendCreateRequest(name string, processingTime int, willUseResource bool, whichResource int){
	defer wg.Done()
	posturl := "http://localhost:8080/createTask"
	t := Task{
		Name: name,
		ProcessingTimeInSeconds: processingTime,
		WillUseResource: willUseResource,
		WhichResource: whichResource,
	}
	body, err := json.Marshal(t)
	if err != nil {
		log.Println("Marshal error : ", err)
		return
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Println("New request creation error : ", err)
		return
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		if errors.Is(err, syscall.ECONNRESET) {
            log.Print("This is connection reset by peer error", err)
        } else {
			log.Println("Client Do error : ", err)
		}
		return
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("io.ReadAll error",err)
			return
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}