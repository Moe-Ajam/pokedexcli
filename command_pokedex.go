package main

import "fmt"

func callbackPokedex(cfg *config, param string) error {
	if len(coughtPokemon) == 0 {
		fmt.Println("You haven't cought any Pokemon yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range coughtPokemon {
		fmt.Println("	-", pokemon.Name)
	}
	return nil
}
