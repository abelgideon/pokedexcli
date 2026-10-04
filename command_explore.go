package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you did not provide the right number of arguments")
	}

	locArea, err := cfg.pokeapiClient.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Println("Found Pokemon:")

	for _, la := range locArea.PokemonEncounters {
		fmt.Printf(" - %v\n", la.Pokemon.Name)
	}

	return nil
}
