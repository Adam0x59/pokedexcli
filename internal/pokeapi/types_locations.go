package pokeapi

// Struct to store the paginated return value from calling the
// location-area api with no id or name specified
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type RespDeepLocation struct {
	PokemonEncounters []struct {
		Pokemon struct {
			PokeName string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
