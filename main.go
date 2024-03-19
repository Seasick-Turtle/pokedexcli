package main

import "fmt"

func (c config) getLocationNames() {
	for i := range len(c.Results) {
		fmt.Println(c.Results[i].Name)
	}
}

type config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	config := config{}
	startRepl(&config)
}
