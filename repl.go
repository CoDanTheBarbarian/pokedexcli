package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CoDanTheBarbarian/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		if len(cleanText) == 0 {
			continue
		}
		commandName := cleanText[0]
		input := strings.Join(cleanText[1:], " ")
		availCommands := getCommands()
		command, ok := availCommands[commandName]
		if !ok {
			fmt.Println("invalid command - enter 'help' for info")
			continue
		}
		if err := command.callback(cfg, input); err != nil {
			fmt.Println("ERROR:", err)
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
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient   pokeapi.Client
	nextPageURL     *string
	previousPageURL *string
	pokedex         map[string]pokeapi.Pokemon
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
		"map": {
			name:        "map",
			description: "Shows a list of 20 location areas - enter again for the next 20 in the list",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Return to the previous list of 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area-name>",
			description: "Lists all the pokemon that can be found in the named area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempt to catch the named pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Displays info about the named pokemon",
			callback:    commandInspect,
		},
	}
}
