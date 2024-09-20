package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Available commands:")
	commands := getCommands()
	for _, command := range commands {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
