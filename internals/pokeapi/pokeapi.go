package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/afreen23/pokedex/internals/pokecache"
)

var location_endpoint = "https://pokeapi.co/api/v2/location-area/"

var pokemon_endpoint = "https://pokeapi.co/api/v2/pokemon/"

type Location struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

var ID = 1

var cache *pokecache.Cache = pokecache.NewCache(5 * time.Second)

func Get[T int | string](id T) ([]byte, error) {
	fetch_url := fmt.Sprintf("%s/%v", location_endpoint, id)
	// finding in cache first
	res, ok := cache.Get(fetch_url)
	// making GET request if not found in cache
	if !ok {
		res, err := http.Get(fetch_url)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("response failed with status code: %d and \nbody: %s\n url: %s", res.StatusCode, body, fetch_url)
		}
		if err != nil {
			return nil, err
		}
		cache.Add(fetch_url, body)
		return body, nil
	}
	return res, nil
}

func GetPokemon(name string) ([]byte, error) {
	fetch_url := fmt.Sprintf("%s/%v", pokemon_endpoint, name)

	res, err := http.Get(fetch_url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and \nbody: %s\n url: %s", res.StatusCode, body, fetch_url)
	}
	if err != nil {
		return nil, err
	}

	return body, nil

}
