package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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

		command.callback()
	}
}

type commands struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]commands {
	return map[string]commands{
		"exit": {
			name:        "exit",
			description: "Exit application",
			callback:    exitCommand,
		},
		"help": {
			name:        "help",
			description: "Prints all available commands",
			callback:    helpCommand,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	commands := strings.Fields(lowered)

	return commands
}
