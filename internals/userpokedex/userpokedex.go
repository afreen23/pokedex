package userpokedex

import "github.com/afreen23/pokedex/internals/pokeapi"

type Pokedex struct {
	Pokemons map[string]pokeapi.Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Pokemons: make(map[string]pokeapi.Pokemon),
	}
}
