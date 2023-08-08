package main

import (
	"fmt"
	"os"
)

func next20Locations() {}

func prev20Locations() {}

func usage() {
	fmt.Println("Usage:")
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
		"map":  next20Locations,
		"mapb": prev20Locations,
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
