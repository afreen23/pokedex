package commands

import "fmt"

func Help() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("-----------------------")

	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Prints the next 20 location areas of Pokemon world")
	fmt.Println("mapb: Prints the previous 20 location areas of Pokemon world")
	fmt.Println("catch: Takes the name of pokemon and throws a Pokeball to catch it!\n e.g: `catch pikachu`")
	fmt.Println("explore: Takes a locations and provides the list of found pokemons at that location\n e.g: `explore canalave-city-area`")
	fmt.Println("inspect: Inspect a cuaght pokeman.\n e.g: `inspect pikachu`")
	fmt.Println("pokedex: Displays all the caught pokemons.")
	fmt.Println("")
}
