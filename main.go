package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/kwschnizel/pokedexgo/internal/pokeapi"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	config := cmdConfig{
		pokeClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := cmds[input[0]]
		if !ok {
			fmt.Println("Unknown Command: enter 'help' or 'h' for a list of commands")
			continue
		}

		var args []string
		if len(input) > 1 {
			args = input[1:]
		}

		err := command(&config, args)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
