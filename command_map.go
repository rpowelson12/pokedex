package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMap() error {
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

	fmt.Printf("%s", &body)

	return nil
}
