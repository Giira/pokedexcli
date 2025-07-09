package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Giira/pokedexcli/internal/pokeapi"
	"github.com/Giira/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	client   pokeapi.Client
	cache    *pokecache.Cache
	area     *string
	next     *string
	previous *string
}

func mapCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas. Subsequent calls display the next 20",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemon in an area",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\n\nUsage:\n")
	for _, c := range mapCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	locsRes, err := cfg.client.GetLocAreas(cfg.next, cfg.cache)
	if err != nil {
		return err
	}

	cfg.next = locsRes.Next
	cfg.previous = locsRes.Previous

	for _, loc := range locsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}

	locsRes, err := cfg.client.GetLocAreas(cfg.previous, cfg.cache)
	if err != nil {
		return err
	}

	cfg.next = locsRes.Next
	cfg.previous = locsRes.Previous

	for _, loc := range locsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cfg *config) error {
	areaExps, err := cfg.client.GetAreaExplore(cfg.area, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *cfg.area)
	fmt.Println("Found Pokemon:")
	for _, pok := range areaExps.PokemonEncounters {
		fmt.Printf("- %s\n", pok.Pokemon.Name)
	}
	return nil
}
