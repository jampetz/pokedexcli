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

		availableCommands := getCommands()
		command, ok := availableCommands[cleanCommand[0]]
		if !ok {
			fmt.Println("Unknown command\nType \"help\" to get the list of available commands")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type commands struct {
	name        string
	description string
	callback    func(cfg *config) error
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
