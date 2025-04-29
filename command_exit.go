package main

import (
	"fmt"
	"os"
)

func commandExit(page *Page) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
