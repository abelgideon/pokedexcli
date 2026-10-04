package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you did not provide the right number of arguments")
	}

	pok, exists := cfg.caughtPokemon[args[0]]
	if !exists {
		return fmt.Errorf("you have not caught that Pokemon yet")
	}

	fmt.Printf("Name: %v\n", pok.Name)
	fmt.Printf("Height: %v\n", pok.Height)
	fmt.Printf("Weight: %v\n", pok.Weight)
	fmt.Println("Stats:")

	for _, stat := range pok.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, tp := range pok.Types {
		fmt.Printf(" - %v\n", tp.Type.Name)
	}

	return nil
}
