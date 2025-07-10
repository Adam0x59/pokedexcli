package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adam0x59/pokedexcli/internal/pokecache"
)

// Function to return a list of locations returned from pokeapi when
// calling locations-area without arguments.
func (c *Client) ListLocations(cache *pokecache.Cache, pageURL *string) (RespShallowLocations, error) {

	// Generate url for the api call
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedDat, ok := cache.Get(url)
	var dat []byte
	if !ok {

		// Create a new GET request using generated url
		// store the request in req
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		// Actually send the new GET request to the server
		// Store the response in resp
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		// On function exit close the response before exit
		defer resp.Body.Close()

		// Read response body into memory - store in dat
		datNew, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		cache.Add(url, datNew)
		dat = datNew
		fmt.Println()
		fmt.Printf("URL: %s", url)
	} else {
		dat = cachedDat
		fmt.Println()
		fmt.Println("Cache Used!")
		fmt.Printf("URL: %s", url)
		fmt.Println()
	}

	// Create an empty struct to store response data
	// Unmarshall the JSON data and store in empty struct
	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Return struct containing unmarshalled JSON data
	return locationsResp, nil
}
