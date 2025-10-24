package main

import (
	"fmt"
)

func cmdPokedex(config *cmdConfig, _ []string) error {
	if len(config.pokedex) == 0 {
		fmt.Println("Your pokedex is empty...")
		return nil
	}
	for _, v := range config.pokedex {
		fmt.Printf(" - %v\n", v.Name)
	}
	return nil
}
