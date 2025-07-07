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
				fmt.Println(err)
			}
		} else {
			fmt.Println("What madness is this command?! Try again!")
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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
