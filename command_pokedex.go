package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	for _, pok := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", pok.Name)
	}
	return nil
}
