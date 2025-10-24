package pokeapi

import (
	"encoding/json"
	"fmt"
)

const pokemonUrl = baseUrl + "/pokemon"

func (c *Client) GetPokemonInfo(pokemon string) (PokemonInfo, error) {
	url := pokemonUrl + "/" + pokemon

	body, err := c.GetBody(url)
	if err != nil {
		return PokemonInfo{}, fmt.Errorf("error from Client.GetPokemonInfo: %w", err)
	}

	var monInfo PokemonInfo
	if err := json.Unmarshal(body, &monInfo); err != nil {
		return PokemonInfo{}, fmt.Errorf("error from GetPokemonInfo: error from json.Unmarshal: %w", err)
	}

	return monInfo, nil

}
