package main

import "pokedexcli/internal/pokeapi"

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&config)
}
