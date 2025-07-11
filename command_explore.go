package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, arg string) error {

	if arg == "" {
		return errors.New("please provide a location to explore")
	}

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(cfg.pokecache, arg)
	if err != nil {
		return err
	}

	//Print locations to the terminal
	fmt.Println()
	fmt.Printf("Exploring %s...\n", arg)
	fmt.Println("Pokemon Found:")
	for _, poke := range pokemonResp.PokemonEncounters {
		fmt.Printf(" - %s\n", poke.Pokemon.PokeName)
	}
	fmt.Println()

	return nil
}
