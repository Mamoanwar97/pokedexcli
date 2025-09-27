package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	pokemonDetailsResp, err := cfg.pokeapiClient.PokemonDetails(cfg.params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonDetailsResp.Name);

	possiblityToCatch := rand.Intn(pokemonDetailsResp.BaseExperience)


	if possiblityToCatch > pokemonDetailsResp.BaseExperience / 2 {
		fmt.Printf("%s was caught!\n", pokemonDetailsResp.Name)
		cfg.pokemonCaught[pokemonDetailsResp.Name] = pokemonDetailsResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonDetailsResp.Name)
	}

	return nil
}