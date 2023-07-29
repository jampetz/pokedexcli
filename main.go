package main

import "github.com/jampetz/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaULR *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
