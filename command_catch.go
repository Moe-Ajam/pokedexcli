package main

import (
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, param string) error {
	resp, err := cfg.pokeapiClient.ListPokemonsInformation(param)
	if err != nil {
		return err
	}

	_, isCought := coughtPokemon[param]

	if isCought {
		fmt.Println(param, "is already in your collection!")
		return nil
	}

	fmt.Println("Throwing ball at ", param, "...")

	if catchPokemon(resp.BaseExperience) {
		fmt.Println(param, "cought!")
		coughtPokemon[param] = resp
		return nil
	}

	fmt.Println(param, "escaped!")

	return nil
}

func catchPokemon(baseExperience int) bool {
	maxCatchProbability := 0.8
	minCatchProbability := 0.2

	catchProbability := maxCatchProbability - ((maxCatchProbability - minCatchProbability) * float64(baseExperience) / 1000.0)

	if catchProbability < minCatchProbability {
		catchProbability = minCatchProbability
	}

	randomValue := rand.Float64()

	return randomValue < catchProbability
}
