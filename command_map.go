package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Page struct {
	Next     string
	Previous string
}

func commandMap(page *Page) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	request, err := http.Get(url)
	if err != nil {
		return errors.New("Unable to access PokeApi")
	}
	body, err := io.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return errors.New("Error reading body")
	}
	// Have the JSON for the location areas now
	failed := json.Unmarshal(body, &page)
	if failed != nil {
		return errors.New("Error assigning JSON to struct")
	}
	fmt.Printf(page.Next)
	return nil
}
