package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c config) saveConfigPrevious(prevUrl string) {
	Previous := c.getConfigPrevious()
	*Previous = prevUrl
}

func (c config) getConfigPrevious() *string {
	return &c.Previous
}

func commandMapB(config *config) error {

	if config.Previous == "" {
		return errors.New("Pokedex is at the first set of locations, can't go back.")
	}

	res, err := http.Get(config.Previous)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Println(err)
	}
	config.saveConfigPrevious(config.Previous)
	config.getLocationNames()

	return nil
}
