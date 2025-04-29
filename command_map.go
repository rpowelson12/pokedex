package main

import (
	"errors"
	"fmt"
)

func commandMap(pokeAPI *PokeApi) error {
	if len(pokeAPI.Next) == 0 {
		result, err := FetchLocationAreas("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
			return errors.New("Error fetching locations")
		}
		pokeAPI.Next = result.Next
		pokeAPI.Previous = result.Previous
		for _, location := range result.Results {
			fmt.Println(location.Name)
		}
	} else {
		result, err := FetchLocationAreas(pokeAPI.Next)
		if err != nil {
			return errors.New("Error fetching locations")
		}
		pokeAPI.Next = result.Next
		pokeAPI.Previous = result.Previous
		for _, location := range result.Results {
			fmt.Println(location.Name)
		}
	}
	return nil
}
