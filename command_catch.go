package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must enter a pokemon name")
	}
	pokemon := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	name := pokemonResp.Name

	fmt.Printf("Throwing a Pokeball at %s\n", name)
	if rand.IntN(pokemonResp.BaseExperience) > 45 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)

	return nil
}
