package cli

import (
	"fmt"
	"os"
	"strings"
	"tasktracker/config"
	"tasktracker/json"
	"tasktracker/task"
)

func Execute(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command provided")
	}

	command := args[0]
	switch command {
	case "add":
		add(args[1:]...)

	case "list":
		listAll(args[1:]...)

	case "exit":
		os.Exit(0)
	}

	return nil
}

func add(args ...string) error {
	description := strings.Join(args, " ")

	task := task.NewTask(description)
	config.Tasks.Add(&task)

	data, err := config.Tasks.ToJSON()
	if err != nil {
		return err
	}

	err = json.UpdateFile(config.Filename, data)
	if err != nil {
		return err
	}

	return nil
}

func listAll(args ...string) {
	if len(args) == 0 {
		config.Tasks.ListAll()
		return
	}

	status := args[0]

	switch status {
	case "open":
		config.Tasks.ListByStatus(task.Open)
	case "in_progress":
		config.Tasks.ListByStatus(task.InProgress)
	case "done":
		config.Tasks.ListByStatus(task.Done)
	default:
		fmt.Println("Invalid status provided")
	}
}
