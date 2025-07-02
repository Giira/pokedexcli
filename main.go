package main

import (
	"github.com/Giira/pokedexcli/internal/pokeapi"
)

func main() {
	pokeApiClient := pokeapi.NewClient()
	cfg := &config{
		client: pokeApiClient,
	}
	catchInput(cfg)
}
