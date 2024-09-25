package main

import (
	"time"

	"github.com/CoDanTheBarbarian/pokedexcli/internal/pokeapi"
)

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(time.Second*5, time.Minute*5),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&config)
}
