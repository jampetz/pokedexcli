package main

import (
	"fmt"
)

func helpCommand(cfg *config) error {
	fmt.Println("List of available commands:")
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf(" - %v [%v]\n", cmd.name, cmd.description)
	}
	return nil
}
