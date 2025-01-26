package pokeapi

type LocationAreaResp struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

type ExploreLocationResp struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	BaseExperience 	int `json:"base_experience"`
	Name			string `json:"name"`
	ID				int `json:"id"`
	Height			int `json:"height"`
	Weight			int `json:"weight"`
	Stats []struct {
		BaseStat	int  `json:"base_stat"`
		Stat struct {
			Name 	string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"` 
		} `json:"type"`
	} `json:"types"`
}