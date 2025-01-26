package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/tadoodles/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient 		pokeapi.Client
	nextLocationsURL 	*string
	prevLocationsURL 	*string
	args				[]string
	caughtPokemon		map[string]pokeapi.Pokemon
}

func startrepl(cfg *config) {	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
    		input := scanner.Text()
			cleanedInput := cleanInput(input)

			if len(cleanedInput) == 0 {
				continue
			}

			if len(cleanedInput) > 1 {
				cfg.args = cleanedInput[1:] // Everything after the command
			} else {
				cfg.args = []string{} // Reset to an empty slice if no arguments provided
			}

			command := cleanedInput[0]
			if cmd, ok := getCommands()[command]; ok {
				err := cmd.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string{
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
	"exit": {
        name:        	"exit",
        description: 	"Exit the Pokedex",
        callback:    	commandExit,
    },
	"help": {
		name:        	"help",
        description: 	"Displays a help message",
        callback:    	commandHelp,
	},
	"map": {
		name:			"map",
		description: 	"Get the next page of locations",
		callback:		commandMapf,
	},
	"mapb": {
		name:			"map",
		description: 	"Get the previous page of locations",
		callback:		commandMapb,
	},
	"explore": {
		name:			"explore",
		description: 	"Explore a chosen location",
		callback:		commandExplore,
	},
	"catch": {
		name:			"catch",
		description:	"Attempt to catch a pokemon",
		callback:		commandCatch,
	},
	"inspect": {
		name:			"inspect",
		description:	"Check stats of caught pokemon",
		callback:		commandInspect,
	},
	"pokedex": {
		name:			"pokedex",
		description:	"Check pokedex",
		callback:		commandPokedex,
	},
	}
}