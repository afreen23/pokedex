package pokecache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/afreen23/pokedex/internals/pokecache"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}

}

func TestReapLoop(t *testing.T) {
	input_interval := 5 * time.Second
	wait_time := input_interval + 5*time.Second

	cache := pokecache.NewCache(input_interval)
	cache.Add("foo", []byte("bar"))

	time.Sleep(wait_time)

	_, ok := cache.Get("foo")
	if !ok {
		t.Errorf("Function has not deleted the value created before %v time", input_interval)
	}
}
