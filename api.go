package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchLocationAreas(url string) (*PokeApi, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	pokeApi := PokeApi{}
	err = json.Unmarshal(body, &pokeApi)
	if err != nil {
		return nil, err
	}

	return &pokeApi, nil
}
