package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	inputScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Pokedex CLI tool\nType \"help\" to get the list of available commands")
	for {
		fmt.Print(" >")
		inputScanner.Scan()
		text := inputScanner.Text()
		cleanCommand := cleanInput(text)
		if len(cleanCommand) == 0 {
			continue
		}
		args := []string{}
		if len(cleanCommand) > 1 {
			args = cleanCommand[1:]
		}
		commandName := cleanCommand[0]

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command\nType \"help\" to get the list of available commands")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type commands struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]commands {
	return map[string]commands{
		"map": {
			name:        "map",
			description: "Show map (moving forward)",
			callback:    mapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Show map (moving backward)",
			callback:    mapBackward,
		},
		"explore": {
			name:        "explore {locationId or locationName}",
			description: "Explore specified location",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemonId or pokemonName}",
			description: "Catch specified pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemonName}",
			description: "Show detailed information about pokemon",
			callback:    callbackInpect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View caught pokemon in your pokedex",
			callback:    callbackPokedex,
		},
		"help": {
			name:        "help",
			description: "Prints all available commands",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit application",
			callback:    exitCommand,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	commands := strings.Fields(lowered)

	return commands
}
