package pokeapi

import (
	"encoding/json"
	"fmt"
)

const locationUrl = baseUrl + "/location-area"

func (c *Client) GetLocationsList(pageUrl string) (LocationsList, error) {
	url := pageUrl
	if url == "" {
		url = locationUrl
	}

	body, err := c.GetBody(url)
	if err != nil {
		return LocationsList{}, fmt.Errorf("error from Client.GetLocationsList: %w", err)
	}

	var locations LocationsList
	if err := json.Unmarshal(body, &locations); err != nil {
		return LocationsList{}, fmt.Errorf("error from ListLocations: error from json.Unmarshal: %w", err)
	}

	return locations, nil

}

func (c *Client) GetLocationInfo(area string) (LocationInfo, error) {
	url := locationUrl + "/" + area
	body, err := c.GetBody(url)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("error from Client.GetLocationInfo: %w", err)
	}

	var locInfo LocationInfo
	if err := json.Unmarshal(body, &locInfo); err != nil {
		return LocationInfo{}, fmt.Errorf("error from Client.GetLocationInfo: error from json.Unmarshal: %w", err)
	}

	return locInfo, nil
}
