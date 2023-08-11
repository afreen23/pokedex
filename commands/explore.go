package commands

import (
	"encoding/json"
	"fmt"

	"github.com/afreen23/pokedex/internals/pokeapi"
)

func Explore(loc string) {
	if loc == "" {
		fmt.Println("No location provided!")
		fmt.Println("Enter `explore <location-name>` to find pokemons")
		return
	}
	fmt.Printf("Exploring %s city area ......\n", loc)
	res, err := pokeapi.GetLocation(loc)
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
