package main

import "fmt"

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range getCommand() {
		name := key
		description := value.description
		fmt.Printf("%s: %s\n", name, description)
	}
	return nil
}
