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
	cfg.page++

	fmt.Print("The location names:\n")
	fmt.Println("Page:", cfg.page)
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLoactionAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient
	if cfg.prevLocationAreaUrl == nil {
		fmt.Println("No more pages to go back to! try to use the `help` command for more information")
		return nil
	}
	resp, err := pokeapiClient.PrevLocationAreas(cfg.prevLocationAreaUrl)

	if err != nil {
		log.Fatal("Something went wrong")
	}

	cfg.page--

	fmt.Print("The location names:\n")
	fmt.Println("Page:", cfg.page)
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLoactionAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
