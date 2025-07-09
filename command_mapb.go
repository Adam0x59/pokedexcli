package main

import "fmt"

func commandMapb() error {

	if mapOffset > 20 {
		mapOffset -= 40
		return commandMap()
	}
	return fmt.Errorf("you're on the first page")
}
