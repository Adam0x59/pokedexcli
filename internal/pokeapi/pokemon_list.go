package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/adam0x59/pokedexcli/internal/pokecache"
)

func (c *Client) ListPokemon(cache *pokecache.Cache, arg string) (RespDeepLocation, error) {

	// Generate url for the api call
	url := baseURL + "/location-area" + "/" + arg

	cachedDat, ok := cache.Get(url)
	var dat []byte
	if !ok {

		// Create a new GET request using generated url
		// store the request in req
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespDeepLocation{}, err
		}

		// Actually send the new GET request to the server
		// Store the response in resp
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespDeepLocation{}, err
		}
		if resp.StatusCode > 299 {
			return RespDeepLocation{}, errors.New("provided Location invalid, please try again")
		}
		// On function exit close the response before exit
		defer resp.Body.Close()

		// Read response body into memory - store in dat
		datNew, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespDeepLocation{}, err
		}
		cache.Add(url, datNew)
		dat = datNew

	} else {
		dat = cachedDat
	}

	// Create an empty struct to store response data
	// Unmarshall the JSON data and store in empty struct
	pokemonResp := RespDeepLocation{}
	err := json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespDeepLocation{}, err
	}

	// Return struct containing unmarshalled JSON data
	return pokemonResp, nil
}
