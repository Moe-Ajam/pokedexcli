package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Type `help` to get started")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]

		availableCommands := getCommands()

		selectedCommand, ok := availableCommands[command]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		if len(cleaned) < 2 {
			selectedCommand.callback(cfg, "")
			continue
		}

		param := cleaned[1]
		selectedCommand.callback(cfg, param)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 areas of the Pokemon world",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 areas of the Pokemon world",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore",
			description: "Show the Pokemons living in the mentioned area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon! be careful though, the higher experince the Pokemon has, the harder it will be to catch it",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show information about a cought Pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the list of Pokemon you have cought so far",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
