package commands

import (
	"encoding/json"
	"fmt"

	"github.com/afreen23/pokedex/internals/pokeapi"
)

func Map() {
	locations := pokeapi.Location{}
	limit := 20
	if pokeapi.LocationID != 1 {
		limit = limit + pokeapi.LocationID - 1
	}
	for ; pokeapi.LocationID <= limit; pokeapi.LocationID++ {
		res, err := pokeapi.GetLocation(pokeapi.LocationID)
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
