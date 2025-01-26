package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationsList(pageURL *string) (LocationAreaResp, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok{
		locationsResp := LocationAreaResp{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationsresp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locationsresp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(url, dat)
	return locationsresp, nil
}