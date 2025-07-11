package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adam0x59/pokedexcli/internal/pokeapi"
	"github.com/adam0x59/pokedexcli/internal/pokecache"
)

// Struct to store cli-commands
type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

// Struct to store callback config, allows config data to persist outside repl loops...
type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokecache        *pokecache.Cache
	pokedex          map[string]pokeapi.RespDeepPokemon
}

func startRepl(cfg *config) {
	r := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		if !r.Scan() {
			break
		}
		command := cleanInput(r.Text())
		if len(command) <= 0 {
			continue
		}
		var arg string
		if len(command) > 1 {
			arg = command[1]
		}
		comm, exists := getCommand()[command[0]]
		if exists {
			err := comm.callback(cfg, arg)
			if err != nil {
				fmt.Println()
				fmt.Println(err)
				fmt.Println()
			}
		} else {
			fmt.Println()
			fmt.Println("What madness is this command?! Try again!")
			fmt.Println()
		}
	}
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List the next 20 map locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area. Useage: explore <area-name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon. Usage: catch <pokemon-name>",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
