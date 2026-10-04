package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you did not provide the right number of arguments")
	}

	pok, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon: %v", err)
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pok.Name)

	roll := rand.Intn((pok.BaseExperience / 35) + 1)

	if roll == 0 {
		cfg.caughtPokemon[pok.Name] = pok
		fmt.Printf("%v was caught!\n", pok.Name)
	} else {
		fmt.Printf("%v escaped!\n", pok.Name)
	}

	return nil
}
