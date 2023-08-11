package command

import (
	"fmt"

	"github.com/afreen23/pokedex/internals/userpokedex"
)

func Pokedex(p *userpokedex.Pokedex) {
	fmt.Println("Your Pokedex:")
	for pok := range p.Pokemons {
		fmt.Printf("- %s\n", pok)
	}
}
