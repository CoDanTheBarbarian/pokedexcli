package main

import (
	"fmt"
	"math/rand"
	"os"
)

func commandHelp(cfg *config, input string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Available commands:")
	commands := getCommands()
	for _, command := range commands {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}

func commandExit(cfg *config, input string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, input string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextPageURL)
	if err != nil {
		return err
	}
	cfg.nextPageURL = resp.Next
	cfg.previousPageURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}

func commandMapb(cfg *config, input string) error {
	if cfg.previousPageURL == nil {
		return fmt.Errorf("can not go back from page one")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousPageURL)
	if err != nil {
		return err
	}
	cfg.nextPageURL = resp.Next
	cfg.previousPageURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}

func commandExplore(cfg *config, input string) error {
	if input == "" {
		return fmt.Errorf("no area specified")
	}
	resp, err := cfg.pokeapiClient.ExploreAreaResponse(input)
	if err != nil {
		return err
	}
	for _, data := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", data.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("no pokemon specified")
	}
	pokemonData, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a pokeball at %s...\n", pokemonData.Name)
	const threshold = 60
	throw := rand.Intn(pokemonData.BaseExperience)
	if throw > threshold {
		fmt.Printf("%s escaped", pokemonData.Name)
		return nil
	} else {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
	}
	cfg.pokedex[pokemon] = pokemonData

	return nil
}

func commandInspect(cfg *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("no pokemon specified")
	}
	pokemonData, ok := cfg.pokedex[pokemon]
	if !ok {
		fmt.Printf("You have not caught %s yet\n", pokemon)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Printf("Stats:\n")
	for _, instance := range pokemonData.Stats {
		fmt.Printf("  - %s: %d\n", instance.Stat.Name, instance.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, instance := range pokemonData.Types {
		fmt.Printf("  - %s\n", instance.Type.Name)
	}
	return nil
}
