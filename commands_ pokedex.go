package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {
	fmt.Println("Pokedex:")
	for _, value := range cfg.pokemonCaught {
		fmt.Printf("- %s\n", value.Name)
	}

	return nil
}
