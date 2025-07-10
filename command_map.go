package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	// Get next list of locations, store in locationsResp,
	// If there is no nextLocationsURL first 20 will be returned
	// by default.
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// Store next and previous URLs returned by api in the config struct
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	//Print locations to the terminal
	fmt.Println()
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(cfg *config) error {
	// Check if there is a previous page stored in the config
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Get the previous list of locations
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// Store next and previous URLs returned by api in the config struct
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	//Print locations to the terminal
	fmt.Println()
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}
