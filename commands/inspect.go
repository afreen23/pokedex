package commands

import (
	"fmt"

	"github.com/afreen23/pokedex/internals/userpokedex"
)

func Inspect(name string, p *userpokedex.Pokedex) {
	if name == "" {
		fmt.Println("No name provided")
		return
	}
	data, ok := p.Pokemons[name]
	if !ok {
		fmt.Printf("You have not caught %s\n", name)
		return
	}
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight %d\n", data.Weight)
	fmt.Println("Stats: ")
	for _, s := range data.Stats {
		fmt.Printf("-%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range data.Types {
		fmt.Printf("-%s\n", t.Type.Name)
	}
}
