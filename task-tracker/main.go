package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
