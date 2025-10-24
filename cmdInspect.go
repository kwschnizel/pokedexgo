package main

import (
	"fmt"
)

func cmdInspect(config *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Inspect command requires a pokemon name. \n Usage:  inspect <pokemon_name>")
	}

	pokemonName := args[0]
	pokemon, ok := config.pokedex[pokemonName]
	if !ok {
		fmt.Printf("you have not caught a %v\n", pokemonName)
		return nil
	}

	fmt.Printf(" Name: %v\n", pokemon.Name)
	fmt.Printf(" Height: %v\n", pokemon.Height)
	fmt.Printf(" Weight: %v\n", pokemon.Weight)

	fmt.Printf(" Stats:\n")
	for _, v := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", v.Stat.Name, v.BaseStat)
	}

	fmt.Printf(" Types:\n")

	for _, v := range pokemon.Types {
		fmt.Printf("  - %v\n", v.Type.Name)
	}

	return nil
}
