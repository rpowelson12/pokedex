package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	// get the pokemon from API
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// get random number to see if we catch the pokemon

	randomNum := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if randomNum > 40 {
		fmt.Printf("%s got away!\n", pokemon.Name)
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
