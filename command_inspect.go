package main

import "fmt"

func callbackInspect(cfg *config, param string) error {
	// case of Pokemon not in the deck already
	cp, ok := coughtPokemon[param]

	if !ok {
		fmt.Println(param, "is not in your deck")
		return nil
	}

	// else if the Pokemon is cought, show some information about it
	fmt.Println("Name:", cp.Name)
	fmt.Println("Height:", cp.Height)
	fmt.Println("Weight:", cp.Weight)
	fmt.Println("Base Experience:", cp.BaseExperience)

	return nil
}
