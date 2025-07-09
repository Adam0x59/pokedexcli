package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println()
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	fmt.Println()
	os.Exit(0)
	return nil
}
