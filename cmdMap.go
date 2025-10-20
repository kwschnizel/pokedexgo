package main

import (
	"fmt"
)

func cmdMap(config *cmdConfig, _ []string) error {
	locations, err := config.pokeClient.ListLocations(config.nextLocationUrl)
	if err != nil {
		return err
	}

	config.nextLocationUrl = locations.Next
	config.prevLocationUrl, _ = locations.Previous.(string)
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func cmdMapb(config *cmdConfig, _ []string) error {
	if config.prevLocationUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := config.pokeClient.ListLocations(config.prevLocationUrl)
	if err != nil {
		return err
	}

	config.nextLocationUrl = locations.Next
	config.prevLocationUrl, _ = locations.Previous.(string)
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
