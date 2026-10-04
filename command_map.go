package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	maps, err := cfg.pokeapiClient.GetLocations(cfg.next)
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
