package main

import (
	"errors"
	"fmt"
)

func callbackInpect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("pokemonName is nil")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("pokemon was not found in your pokedex")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Species: %s\n", pokemon.Species.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Print("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}

	return nil
}
