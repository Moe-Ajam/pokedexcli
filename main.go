package main

import (
	"time"

	"github.com/Moe-Ajam/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLoactionAreaUrl *string
	prevLocationAreaUrl *string
	page                int
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		page:          0,
	}
	startRepl(&cfg)
}
