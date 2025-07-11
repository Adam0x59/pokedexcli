package main

import "fmt"

func commandPokedex(cfg *config, arg string) error {
	fmt.Println()
	fmt.Println("Your Pokedex:")
	for _, item := range cfg.pokedex {
		fmt.Printf(" - %s\n", item.Name)
	}
	fmt.Println()
	return nil
}
