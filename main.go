package main

import (
	"fmt"

	"github.com/afreen23/pokedex/command"
)

func main() {
	var input, option string
	for {
		fmt.Print("pokedex > ")
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
			command.Catch(option)
		default:
			fmt.Println("Invalid command")
			command.Usage()
		}
	}
}
