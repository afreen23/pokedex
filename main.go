package main

import (
	"fmt"

	"github.com/afreen23/pokedex/command"
)

func main() {
	commands := map[string]func(){
		"help": command.Help,
		"exit": command.Exit,
		"map":  command.Map,
		"mapb": command.Mapb,
	}
	var input string
	for {
		fmt.Print("pokedex > ")
		fmt.Scan(&input)
		handler, ok := commands[input]
		if !ok {
			fmt.Println("Invalid command")
			command.Usage()
			break
		}
		handler()
	}
}
