package main

import (
	"os"
	"time"
	"github.com/tadoodles/pokedexcli/internal/pokeapi"
)

func main () {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: 	pokeClient,
		args:			os.Args[1:],
		caughtPokemon:	make(map[string]pokeapi.Pokemon),
	}

	startrepl(cfg)
}