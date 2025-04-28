package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		
		scanner.Scan()
		text := scanner.Text()
		
		clean := cleanInput(text)
		
		fmt.Println("Your command was:", clean[0])
	}
}
