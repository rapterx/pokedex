package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]

		switch command {
		case "help":
			fmt.Println("Welcome to the help menu!")
			fmt.Println("- help")
			fmt.Println("- exit")
			fmt.Println("")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}

		fmt.Println(cleaned)
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
