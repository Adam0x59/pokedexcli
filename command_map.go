package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var mapOffset int = 0

func commandMap() error {

	fmt.Println()
	get_url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", mapOffset)
	res, err := http.Get(get_url)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	results, ok := result["results"].([]interface{})
	if !ok {
		fmt.Println("results is not a slice")
	}

	for _, i := range results {
		entry, ok := i.(map[string]interface{})
		if !ok {
			fmt.Println("i is not a map")
		}

		name, ok := entry["name"].(string)
		if !ok {
			fmt.Printf("\n%v is not a string", name)
		}
		fmt.Println(name)
	}
	fmt.Println()
	mapOffset += 20

	return err
}
