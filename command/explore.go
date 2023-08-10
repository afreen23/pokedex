package command

import (
	"encoding/json"
	"fmt"

	"github.com/afreen23/pokedex/internals/pokeapi"
)

func Explore(loc string) {
	fmt.Printf("Exploring %s city area ......\n", loc)
	res, err := pokeapi.Get(loc)
	if err != nil {
		println("Encountered error", err)
		return
	}
	results := pokeapi.Location{}
	err = json.Unmarshal(res, &results)
	if err != nil {
		println("Encountered error", err)
		return
	}
	fmt.Println("Found Pokemons:")
	for _, v := range results.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}
}
