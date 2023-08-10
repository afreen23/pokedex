package command

import (
	"encoding/json"
	"fmt"

	"github.com/afreen23/pokedex/internals/pokeapi"
)

func Map() {
	locations := pokeapi.Location{}
	limit := 20
	if pokeapi.ID != 1 {
		limit = limit + pokeapi.ID - 1
	}
	for ; pokeapi.ID <= limit; pokeapi.ID++ {
		res, err := pokeapi.Get(pokeapi.ID)
		if err != nil {
			fmt.Print(err)
		}
		err = json.Unmarshal(res, &locations)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(locations.Name)
	}
}
