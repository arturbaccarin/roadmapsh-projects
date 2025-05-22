package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"tasktracker/cli"
	"tasktracker/config"
	"tasktracker/json"
)

/*
TODO
implement sort by ID
fix updatedBy
*/

func main() {
	filename := "tasks.json"

	if !json.FileExists(filename) {
		err := json.CreateFile(filename)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := config.Tasks.LoadFromJSONFile(filename)
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

		args := strings.TrimSpace(input)

		err = cli.Execute(strings.Split(args, " ")...)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
