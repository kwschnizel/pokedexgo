package main

import "fmt"

func cmdExplore(config *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Explore command requires a location area name. \n Usage:  explore <area_name>")
	}

	areaName := args[0]
	fmt.Printf("Exploring %v...\n", areaName)

	locInfo, err := config.pokeClient.GetLocationInfo(areaName)
	if err != nil {
		return fmt.Errorf("error from cmdExplore: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, v := range locInfo.PokemonEncounters {
		fmt.Println(" -", v.Pokemon.Name)
	}

	return nil
}
