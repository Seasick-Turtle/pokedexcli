package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c config) saveConfigNext(nextUrl string) {
	Next := c.getConfigNext()
	*Next = nextUrl
}

func (c config) getConfigNext() *string {
	return &c.Next
}

func commandMap(config *config) error {
	var apiUrl string

	if config.Next == "" {
		apiUrl = "https://pokeapi.co/api/v2/location-area/"
	} else {
		apiUrl = config.Next
	}

	res, err := http.Get(apiUrl)

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
	config.saveConfigNext(config.Next)
	config.getLocationNames()

	return nil
}
