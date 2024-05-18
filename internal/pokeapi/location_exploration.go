package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LocationAreaInformation struct {
	EncounterMethodRates []struct {
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
	} `json:"location"`
	Name              string `json:"name"`
	Names             []any  `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListPokemonsInArea(param string) (LocationAreaInformation, error) {
	if strings.Trim(param, " ") == "" {
		fmt.Println("Please specify a location")
		return LocationAreaInformation{}, fmt.Errorf("Not a valid location")
	}
	endpoint := "https://pokeapi.co/api/v2/location-area/" + param

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Invalid location")
	}

	locationAreaInformation := LocationAreaInformation{}

	data, ok := c.cache.Get(endpoint)

	if ok {
		json.Unmarshal(data, &locationAreaInformation)

		return locationAreaInformation, nil
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaInformation{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaInformation{}, fmt.Errorf("Bad request: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaInformation{}, err
	}

	json.Unmarshal(data, &locationAreaInformation)

	c.cache.Add(endpoint, data)

	return locationAreaInformation, nil
}
