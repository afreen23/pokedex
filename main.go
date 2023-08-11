package main

import (
	"fmt"

	"github.com/afreen23/pokedex/commands"
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
			commands.Help()
		case "exit":
			commands.Exit()
		case "map":
			commands.Map()
		case "mapb":
			commands.Mapb()
		case "explore":
			commands.Explore(option)
		case "catch":
			commands.Catch(option, p)
		case "inspect":
			commands.Inspect(option, p)
		case "pokedex":
			commands.Pokedex(p)
		default:
			fmt.Println("Invalid command")
			fmt.Println("Type `help` to get the list of valid commands")
		}
	}
}
