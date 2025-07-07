package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Giira/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	var output []string
	text = strings.ToLower(text)
	output = strings.Fields(text)
	return output
}

func catchInput(cfg *config, cache *pokecache.Cache) {
	coms := mapCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		boo := scanner.Scan()
		if boo {
			input := scanner.Text()
			input_slice := cleanInput(input)
			switch input_slice[0] {
			case "exit":
				err := coms["exit"].callback(cfg, cache)
				if err != nil {
					fmt.Println(err)
				}
			case "help":
				err := coms["help"].callback(cfg, cache)
				if err != nil {
					fmt.Println(err)
				}
			case "map":
				err := coms["map"].callback(cfg, cache)
				if err != nil {
					fmt.Println(err)
				}
			case "mapb":
				err := coms["mapb"].callback(cfg, cache)
				if err != nil {
					fmt.Println(err)
				}
			default:
				fmt.Print("Unknown command\n")
			}
		}
	}
}
