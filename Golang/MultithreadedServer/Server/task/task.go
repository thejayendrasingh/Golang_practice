package task

/*
	Task package contains task model to be used throughout our server execution
*/

import "time"

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

func CreateNewTask(
	name string,
	processingTimeInSeconds int,
	willUseResource bool,
	whichResource int,
) Task {
	return Task{
		Name:                    name,
		ProcessingTimeInSeconds: processingTimeInSeconds,
		WillUseResource:         willUseResource,
		WhichResource:           whichResource,
		StartedAt:               time.Now(),
		Status:                  Created,
	}
}
