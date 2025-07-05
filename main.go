package main

import (
	"time"

	"github.com/Giira/pokedexcli/internal/pokeapi"
	"github.com/Giira/pokedexcli/internal/pokecache"
)

func main() {
	pokeApiClient := pokeapi.NewClient()
	cache := pokecache.NewCache(5 * time.Second)
	cfg := &config{
		client: pokeApiClient,
	}
	catchInput(cfg, &cache)
}
