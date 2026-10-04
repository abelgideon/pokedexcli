package main

import (
	"fmt"
)

func commandMapB(cfg *config, args ...string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	maps, err := cfg.pokeapiClient.GetLocations(*cfg.previous)
	if err != nil {
		return err
	}

	cfg.previous = maps.Previous
	cfg.next = maps.Next

	for _, mp := range maps.Results {
		fmt.Println(mp.Name)
	}

	return nil
}
