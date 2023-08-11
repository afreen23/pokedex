package command

import "fmt"

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("map:  Prints the next 20 location areas of Pokemon world")
	fmt.Println("mapb:  Prints the previous 20 location areas of Pokemon world")
	fmt.Println("catch: Takes the name of pokemon and throws a Pokeball to catch it!\n e.g: `catch pikachu`")
	fmt.Println("explore:  Takes a locations and provides the list of found pokemons at that location\n e.g: `explore canalave-city-area`")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("")
}

func Help() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("-----------------------")
	Usage()
}
