package main

import (
	"fmt"
	"errors"
	"math/rand"
	"strings"
)

func commandCatch(cfg *config) error {
	if len (cfg.args) < 1 {
		return errors.New("usage: catch <pokemon-name>") 
	}

	pokemon_name := strings.ToLower(strings.ReplaceAll(strings.Join(cfg.args, " "), " ", "-"))
	pokemon_name = strings.TrimSpace(pokemon_name)

	if pokemon_name == "" {
		return errors.New("Please provide pokemon after 'catch'")
	}

	pokemon_info, err := cfg.pokeapiClient.GetPokemon(pokemon_name)
	if err != nil {
		return fmt.Errorf("Failed to fetch Pokemon list from area %s: %w", pokemon_name, err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)
	if catchingChance(pokemon_info.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemon_name)
		cfg.caughtPokemon[pokemon_name] = pokemon_info
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon_name)
	}

	return nil
}

func catchingChance (baseExp int) bool {
	catchRate := 100 - baseExp/2
	if catchRate < 10 {
		catchRate = 10
	}

	return rand.Intn(100) < catchRate
}