package pokeapi

type RespDeepPokemon struct {
	Name   string `json:"name"`
	BaseXP int    `json:"base_experience"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Sname    struct {
			Name string
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
