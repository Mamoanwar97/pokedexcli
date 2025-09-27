package main

import (
	"fmt"
)

func commandInspect(cfg *config) error {
	value, exists := cfg.pokemonCaught[cfg.params[0]]
	if !exists {
		fmt.Printf("you have not caught %s\n", cfg.params[0])
		return nil
	}

	fmt.Printf("Name: %s\n", value.Name)
	fmt.Printf("Height: %d\n", value.Height)
	fmt.Printf("Weight: %d\n", value.Weight)
	fmt.Printf("Base Experience: %d\n", value.BaseExperience)

	return nil
}
