package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must enter a location name")
	}
	location := args[0]

	locationResp, err := cfg.pokeapiClient.ListLocationPokemon(location)
	if err != nil {
		return err
	}

	if len(locationResp.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon Found")
		return nil
	}

	fmt.Printf("Exploring %v \n", location)
	fmt.Println("Found Pokemon: ")

	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Printf("- %v \n", encounter.Pokemon.Name)
	}

	return nil
}
