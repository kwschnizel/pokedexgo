package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type LocationsList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const locationUrl = baseUrl + "/location-area"

func (c *Client) ListLocations(pageUrl string) (LocationsList, error) {
	url := pageUrl
	if url == "" {
		url = locationUrl
	}

	body, ok := c.cache.Get(url)
	if !ok {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return LocationsList{}, fmt.Errorf("Error in ListLocations GET resp: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return LocationsList{}, fmt.Errorf("Error in ListLocations: Response failed with status code %v", res.StatusCode)
		}
		if err != nil {
			return LocationsList{}, fmt.Errorf("Error in ListLocations from io.ReadAll: %w", err)
		}

		c.cache.Add(url, body)
	}

	var locations LocationsList
	if err := json.Unmarshal(body, &locations); err != nil {
		return LocationsList{}, fmt.Errorf("Error in commandMap from json.Unmarshal: %w", err)
	}

	return locations, nil

}
