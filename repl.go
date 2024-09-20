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
		fmt.Print("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		if len(cleanText) == 0 {
			continue
		}
		command := cleanText[0]
		switch command {
		case "help":
			commandHelp()
		case "exit":
			commandExit()
		default:
			fmt.Println("invalid command - enter help for info")
		}
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
