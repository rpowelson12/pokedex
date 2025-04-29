package main

import (
	"errors"
	"fmt"
)

func commandMapb(pokeAPI *PokeApi) error {

	if len(pokeAPI.Previous) == 0 {
		fmt.Println("you're on the first page")
	} else {
		result, err := FetchLocationAreas(pokeAPI.Previous)
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
