package main

import (
	"fmt"

	"github.com/afreen23/pokedex/command"
	"github.com/afreen23/pokedex/internals/userpokedex"
)

func main() {
	var input, option string
	var p *userpokedex.Pokedex = userpokedex.NewPokedex()
	for {
		fmt.Print("Pokedex > ")
		fmt.Scanf("%s%s", &input, &option)
		fmt.Println()
		switch input {
		case "help":
			command.Help()
		case "exit":
			command.Exit()
		case "map":
			command.Map()
		case "mapb":
			command.Mapb()
		case "explore":
			command.Explore(option)
		case "catch":
			command.Catch(option, p)
		case "inspect":
			command.Inspect(option, p)
		case "pokedex":
			command.Pokedex(p)
		default:
			fmt.Println("Invalid command")
			fmt.Println("Type `help` to get the list of valid commands")
		}
	}
}
