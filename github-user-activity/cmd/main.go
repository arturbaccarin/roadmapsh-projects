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

		args := strings.TrimSpace(input)

		err = cli.Execute(strings.Split(args, " ")...)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
