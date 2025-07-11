package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, arg string) error {
	pokemonData, _ /*pokeDat*/, err := cfg.pokeapiClient.CatchPokemon(cfg.pokecache, arg)
	if err != nil {
		return err
	}

	pokeDex := cfg.pokedex

	fmt.Println()
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	chance := 0.0
	if pokemonData.BaseXP < 400 && pokemonData.BaseXP >= 30 {
		chance = 0.9 - (float64(pokemonData.BaseXP) / 600)
	} else if pokemonData.BaseXP < 30 {
		chance = 0.9
	} else if pokemonData.BaseXP >= 400 {
		chance = 0.235
	}
	if rand.Float64() < chance {
		fmt.Printf("%s was caught!\n\n", pokemonData.Name)
		pokeDex[pokemonData.Name] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n\n", pokemonData.Name)
	}

	/*
		fmt.Println("Pokedex currently contains:")
		for item := range pokeDex {
			fmt.Printf("%s\n", pokeDex[item].Name)
		}

		fmt.Println()
		fmt.Printf("%s has %d baseXP\n", pokemonData.Name, pokemonData.BaseXP)
		fmt.Println()
	*/

	return nil
}
