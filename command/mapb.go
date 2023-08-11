package command

import (
	"encoding/json"
	"fmt"

	"github.com/afreen23/pokedex/internals/pokeapi"
)

func Mapb() {
	if pokeapi.LocationID == 1 {
		fmt.Printf("No previous locations to look back\n")
		return
	}
	locations := pokeapi.Location{}
	for i := pokeapi.LocationID - 20; i < pokeapi.LocationID; i++ {
		res, err := pokeapi.GetLocation(i)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(res, &locations)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(locations.Name)
	}
	pokeapi.LocationID = pokeapi.LocationID - 20
}
