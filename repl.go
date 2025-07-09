package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
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
		comm, exists := getCommand()[command[0]]
		if exists {
			err := comm.callback()
			if err != nil {
				fmt.Println()
				fmt.Println(err)
				fmt.Println()
			}
		} else {
			fmt.Println()
			fmt.Println("What madness is this command?! Try again!\n")
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 map locations",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
