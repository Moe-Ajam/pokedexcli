package main

import (
	"fmt"
	"log"
)

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Location struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

func callbackMap(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	resp, err := pokeapiClient.ListLocationAreas(cfg.nextLoactionAreaUrl)

	if err != nil {
		log.Fatal("Something went wrong")
	}

	fmt.Print("The location names:\n")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLoactionAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	resp, err := pokeapiClient.PrevLocationAreas(cfg.prevLocationAreaUrl)

	if err != nil {
		log.Fatal("Something went wrong")
	}

	fmt.Print("The location names:\n")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLoactionAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
