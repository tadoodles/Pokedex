package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonListbyArea(areaName string) ([]string, error) {
	url := baseURL + "/location-area/" + areaName

	if cached, found := c.cache.Get(url); found {
		var exploreresp ExploreLocationResp
		err := json.Unmarshal(cached, &exploreresp)
		if err != nil {
			return nil, err
		}
		return extractPokemonNames(exploreresp), nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var exploreresp ExploreLocationResp
	err = json.Unmarshal(dat, &exploreresp)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, dat)
	return extractPokemonNames(exploreresp), nil
}

func extractPokemonNames(resp ExploreLocationResp) []string {
	var names []string
	for _, encounter := range resp.PokemonEncounters {
		names = append(names, encounter.Pokemon.Name)
	}
	return names
}