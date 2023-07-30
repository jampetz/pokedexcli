package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you haven't caught pokemon yet")
	}

	fmt.Println("Pokemon in pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
