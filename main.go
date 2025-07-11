package main

import (
	"time"

	"github.com/adam0x59/pokedexcli/internal/pokeapi"
	"github.com/adam0x59/pokedexcli/internal/pokecache"
)

func main() {
	// Set system default timing
	clientTimeout := 5 * time.Second
	cacheReapInterval := 100 * time.Second
	// Open a new client
	pokeClient := pokeapi.NewClient(clientTimeout)
	// Create a new cache
	pokeCache := pokecache.NewCache(cacheReapInterval)
	// Create a pokedex to store caught pokemon
	pokeDex := make(map[string]pokeapi.RespDeepPokemon)
	// Create a single state-persistent config struct and
	// add the new client and cache to the struct
	cfg := &config{
		pokeapiClient: pokeClient,
		pokecache:     pokeCache,
		pokedex:       pokeDex,
	}
	// Start the repl and pass in the config struct
	startRepl(cfg)
}
