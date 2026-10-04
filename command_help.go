package main

import (
	"fmt"
	"strings"
)

func commandHelp(cfg *config, args ...string) error {
	usage := []string{"Welcome to the Pokedex!\n\tUsage:\n"}

	for _, v := range cfg.commands {
		usage = append(usage, fmt.Sprintf("\n\t%v: %v", v.name, v.description))
	}

	fmt.Println(strings.Join(usage, ""))

	return nil
}
