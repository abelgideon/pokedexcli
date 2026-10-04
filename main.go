package main

import (
	"time"

	"github.com/abelgideon/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands      map[string]cliCommand
	next          string
	previous      *string
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	myConfig := config{
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Displays 20 location areas in the Pokemon world",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Displays previous 20 location areas in the Pokemon world",
				callback:    commandMapB,
			},
			"explore": {
				name:        "explore",
				description: "Explore Pokemon in a given location",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Try to catch a Pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect Pokemon you have caught",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "Track Pokemon you have caught",
				callback:    commandPokedex,
			},
		},
		next:          "https://pokeapi.co/api/v2/location-area",
		previous:      nil,
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&myConfig)
}
