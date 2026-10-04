package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleaned := cleanInput(userInput)

		if len(cleaned) == 0 {
			continue
		}

		command, exists := cfg.commands[cleaned[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, cleaned[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}
