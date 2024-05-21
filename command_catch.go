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

	fmt.Println(param, "has an experience level of", resp.BaseExperience)
	fmt.Println("Catching Pokemon...")
	if catchPokemon(resp.BaseExperience) {
		fmt.Println("Pokemon cought!")
		return nil
	}

	fmt.Println("Missed :(")

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
