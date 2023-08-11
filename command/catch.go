package command

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/afreen23/pokedex/internals/pokeapi"
	"github.com/afreen23/pokedex/internals/userpokedex"
)

func Catch(pokemon string, p *userpokedex.Pokedex) {
	if pokemon == "" {
		fmt.Println("No pokemon name provided!")
		fmt.Println("Enter `catch <pokemon_name>` to catch a pokemon")
		return
	}

	res, err := pokeapi.GetPokemon(pokemon)
	if err != nil {
		fmt.Println("Encountered Error", err)
		return
	}
	data := pokeapi.Pokemon{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Println("Encountered Error", err)
		return
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	r := rand.Intn(data.BaseExperience)

	if r > 40 {
		fmt.Printf("%s escaped\n", pokemon)
		return
	} else {
		fmt.Printf("%s was caught\n", pokemon)
		p.Pokemons[pokemon] = data
	}

}
