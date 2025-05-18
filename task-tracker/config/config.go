package config

import "tasktracker/task"

const (
	Filename = "tasks.json"
)

var (
	Tasks = task.NewTasks()
)
