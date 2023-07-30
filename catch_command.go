package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("pokemonName is nil")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	catchChance := rand.Intn(pokemon.BaseExperience)
	if catchChance > threshold {
		return fmt.Errorf("failed to catch %s (%v EXP)\nwind was too strong (%v / %v)", pokemon.Name, pokemon.BaseExperience, catchChance, threshold)
	}

	fmt.Printf("%s (%v EXP) was caught!\nyou've got lucky (%v / %v)\n", pokemon.Name, pokemon.BaseExperience, catchChance, threshold)
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
