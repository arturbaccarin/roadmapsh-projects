package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"tasktracker/json"
	"tasktracker/task"
)

func main() {
	filename := "tasks.json"
	tasks := task.NewTasks()

	if !json.FileExists(filename) {
		err := json.CreateFile(filename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := tasks.LoadFromJSONFile(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("cli> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Clean input
		command := strings.TrimSpace(input)

		// Handle commands
		switch command {
		case "exit", "quit":
			fmt.Println("Exiting CLI...")
			return
		case "hello":
			fmt.Println("Hi there!")
		case "":
			// Ignore empty input
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}
