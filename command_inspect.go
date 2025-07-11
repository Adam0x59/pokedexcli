package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, arg string) error {

	pokemonData, ok := cfg.pokedex[arg]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println()
	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf(" -%s: %d\n", stat.Sname.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonData.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	fmt.Println()

	return nil
}
