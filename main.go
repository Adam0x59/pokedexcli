package main

import (
	"time"

	"github.com/adam0x59/pokedexcli/internal/pokeapi"
)

func main() {
	// Open a new client
	pokeClient := pokeapi.NewClient(5 * time.Second)
	// Create a single state-persistent config struct and
	// add the new client to the struct
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	// Start the repl and pass in the config struct
	startRepl(cfg)
}
