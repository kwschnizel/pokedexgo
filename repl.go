package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kwschnizel/pokedexgo/internal/pokeapi"
)

type cmdConfig struct {
	nextLocationUrl string
	prevLocationUrl string
	pokeClient      pokeapi.Client
	pokedex         map[string]pokeapi.PokemonInfo
}

type cmdFunc func(*cmdConfig, []string) error

type cliCmd struct {
	name        string
	alias       []string
	description string
	callback    cmdFunc
}

var cmdList = []cliCmd{}
var cmds = make(map[string]cmdFunc)

func init() {
	// Init cmdList for the help command
	cmdList = []cliCmd{
		{
			name:        "exit",
			alias:       []string{"quit", "q"},
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
		{
			name:        "help",
			alias:       []string{"h"},
			description: "Display the help message",
			callback:    cmdHelp,
		},
		{
			name:        "map",
			alias:       []string{"mapn", "mn"},
			description: "Display the names of the next 20 locations in the PokeWorld",
			callback:    cmdMap,
		},
		{
			name:        "mapb",
			alias:       []string{"mapp", "mp", "mb"},
			description: "Display the names of the previous 20 locations in the PokeWorld",
			callback:    cmdMapb,
		},
		{
			name:        "explore",
			alias:       []string{"ex"},
			description: "Display info of a given <area_name> or <area_id>",
			callback:    cmdExplore,
		},
		{
			name:        "catch",
			alias:       []string{"c"},
			description: "Attempt to catch a given <pokemon_name>",
			callback:    cmdCatch,
		},
		{
			name:        "inspect",
			alias:       []string{"i"},
			description: "Lookup Info on <pokemon_name> from pokemon you've captured",
			callback:    cmdInspect,
		},
		{
			name:        "pokedex",
			alias:       []string{"p", "ls"},
			description: "List all the pokemon in your pokedex",
			callback:    cmdPokedex,
		},
	}

	// Building command map
	for _, c := range cmdList {
		cmds[c.name] = c.callback
		for _, a := range c.alias {
			cmds[a] = c.callback
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	res := strings.Fields(text)
	return res
}

func cmdExit(_ *cmdConfig, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return fmt.Errorf("Error: program should have exited")
}

func cmdHelp(_ *cmdConfig, _ []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("-----------------------")
	fmt.Println()

	for _, v := range cmdList {
		fmt.Printf(" %v (alias: %v): \n   %v\n\n", v.name, strings.Join(v.alias, ", "), v.description)

	}
	return nil
}
