package main

import (
	"fmt"
)

func callbackHelp(cfg *config) error {
	availableCommands := getCommands()
	fmt.Println("Welcome to the Pokedex menu!")
	fmt.Println("Here are the available commands:")
	for _, command := range availableCommands {
		fmt.Println(" - ", command.name, ":", command.description)
	}
	fmt.Println("")
	return nil
}
