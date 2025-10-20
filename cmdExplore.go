package main

import "fmt"

func cmdExplore(config *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Explore command requires a location area name. \n Usage:  explore <area_name>")
	}

	return nil
}
