package main

import (
	"fmt"
	"os"

	"github.com/afreen23/pokedex/internals/pokeapi"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("map:  Prints the next 20 location areas of Pokemon world")
	fmt.Println("mapb:  Prints the previous 20 location areas of Pokemon world")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("")
}

func help() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("-----------------------")
	usage()
}

func exit() {
	os.Exit(0)
}

func main() {
	commands := map[string]func(){
		"help": help,
		"exit": exit,
		"map":  pokeapi.Next20Locations,
		"mapb": pokeapi.Prev20Locations,
	}
	var input string
	for {
		fmt.Print("pokedex > ")
		fmt.Scan(&input)
		handler, ok := commands[input]
		if !ok {
			fmt.Println("Invalid command")
			usage()
			break
		}
		handler()
	}
}
