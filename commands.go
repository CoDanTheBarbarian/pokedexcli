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
