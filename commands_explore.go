package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	exploredLocationResp, err := cfg.pokeapiClient.ExploreLocation(cfg.params[0])
	if err != nil {
		return err
	}

	for _, val := range exploredLocationResp.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}
	return nil
}