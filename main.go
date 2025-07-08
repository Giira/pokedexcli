package main

import (
	"time"

	"github.com/Giira/pokedexcli/internal/pokeapi"
	"github.com/Giira/pokedexcli/internal/pokecache"
)

func main() {
	pokeApiClient := pokeapi.NewClient()
	cfg := &config{
		client: pokeApiClient,
		cache:  pokecache.NewCache(5 * time.Second),
	}
	catchInput(cfg)
}
