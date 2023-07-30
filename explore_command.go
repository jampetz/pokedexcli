package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("location_area ID is nil")
	}
	locationId := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationId)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", resp.Name)
	for _, pokemon := range resp.PokemonEncounters {
		pokemonId, err := getIDFromURL(pokemon.Pokemon.URL)
		if err != nil {
			return errors.New("could not get pokemonId")
		}
		fmt.Printf("%v) %s\n", pokemonId, pokemon.Pokemon.Name)
	}

	return nil
}
