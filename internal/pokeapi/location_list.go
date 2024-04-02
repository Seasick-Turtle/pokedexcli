package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (ShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cacheData, ok := c.cache.Get(url); ok {
		locationsResp := ShallowLocations{}
		err := json.Unmarshal(cacheData, &locationsResp)
		if err != nil {
			return ShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShallowLocations{}, err
	}

	locationsResp := ShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return ShallowLocations{}, err
	}

	c.cache.Add(url, data)
	return locationsResp, nil
}
