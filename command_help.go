package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range getCommand() {
		name := key
		description := value.description
		fmt.Printf("%s: %s\n", name, description)
	}
	fmt.Println()
	return nil
}
