package main

import (
	"fmt"
)

func callbackExplore(cfg *config, param string) error {
	resp, err := cfg.pokeapiClient.ListPokemonsInArea(param)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range resp.PokemonEncounters {
		fmt.Print("- ", pokemonEncounter.Pokemon.Name, "\n")
	}

	return nil
}
