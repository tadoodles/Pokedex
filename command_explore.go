package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandExplore(cfg *config) error {
	if len (cfg.args) < 1 {
		return errors.New("usage: explore <location-name>") 
	}
	
	areaName := strings.Join(cfg.args, " ")
	areaName = strings.TrimSpace(areaName)

	if areaName == "" {
		return errors.New("Please provide location after 'explore'")
	}

	pokemonList, err := cfg.pokeapiClient.GetPokemonListbyArea(areaName)
	if err != nil {
		return fmt.Errorf("Failed to fetch Pokemon list from area %s: %w", areaName, err)
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Printf("Found Pokemon: \n")
	for _, name := range pokemonList {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}