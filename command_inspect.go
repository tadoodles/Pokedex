package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(cfg *config) error {
	if len (cfg.args) < 1 {
		return errors.New("usage: inspect <pokemon-name>") 
	}
	
	pokemon_name := strings.Join(cfg.args, " ")
	pokemon_name = strings.TrimSpace(pokemon_name)

	if pokemon_name == "" {
		return errors.New("Please provide pokemon after 'inspect'")
	}

	pokemon, exists := cfg.caughtPokemon[pokemon_name]
	if !exists {
		return fmt.Errorf("You have not yet caught that Pokemon")
	}

	
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}	

	return nil
}