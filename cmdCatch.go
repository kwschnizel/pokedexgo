package main

import (
	"fmt"
	"math/rand"
	"time"
)

func cmdCatch(config *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Catch command requires a pokemon name. \n Usage:  catch <pokemon_name> or <pokemon_id>")
	}

	pokemon := args[0]

	monInfo, err := config.pokeClient.GetPokemonInfo(pokemon)
	if err != nil {
		return fmt.Errorf("error from cmdCatch: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", monInfo.Name)

	//INFO: Highest baseExp value is 635 but most pokemon are under 300 with legendary pokemon not breaching 400
	// With CatchValue at 1000 most pokemon have a <30% to escape and We are giving 3 attempts to escape
	// This should result in ~22% capture rate for the hardest pokemon
	const randCatchValue = 1000

	fmt.Println("...")
	time.Sleep(400 * time.Millisecond)
	if rand.Intn(randCatchValue) < monInfo.BaseExperience {
		fmt.Printf("%v escaped the Pokeball\n", monInfo.Name)
		return nil
	}

	fmt.Println("....")
	time.Sleep(400 * time.Millisecond)
	if rand.Intn(randCatchValue) < monInfo.BaseExperience {
		fmt.Printf("%v escaped the Pokeball\n", monInfo.Name)
		return nil
	}

	fmt.Println(".....")
	time.Sleep(400 * time.Millisecond)
	if rand.Intn(randCatchValue) < monInfo.BaseExperience {
		fmt.Printf("%v escaped the Pokeball\n", monInfo.Name)
		return nil
	}

	fmt.Printf("Success! %v was captured and added to your Pokedex\n", monInfo.Name)
	config.pokedex[monInfo.Name] = monInfo
	return nil
}
