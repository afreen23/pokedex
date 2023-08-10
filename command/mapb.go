package command

import (
	"encoding/json"
	"fmt"

	"github.com/afreen23/pokedex/internals/pokeapi"
)

func Mapb() {
	if pokeapi.ID == 1 {
		fmt.Printf("No previous locations to look back\n")
		return
	}
	locations := pokeapi.Location{}
	for i := pokeapi.ID - 20; i < pokeapi.ID; i++ {
		res, err := pokeapi.Get(i)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(res, &locations)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(locations)
	}
	pokeapi.ID = pokeapi.ID - 20
}
