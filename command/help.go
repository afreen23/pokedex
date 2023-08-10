package command

import "fmt"

func Usage() {
	fmt.Println("Usage:")
	fmt.Println("map:  Prints the next 20 location areas of Pokemon world")
	fmt.Println("mapb:  Prints the previous 20 location areas of Pokemon world")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("")
}

func Help() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("-----------------------")
	Usage()
}
