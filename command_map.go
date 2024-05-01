package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func callbackMap() {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/?offset=20&limit=20")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Request failed with status code: %d\nbody: %v", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	location := Location{}

	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location)
}
