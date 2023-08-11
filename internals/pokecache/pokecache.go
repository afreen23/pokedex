package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mux   sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	res, ok := c.cache[key]
	c.mux.Unlock()
	if !ok {
		return nil, false
	}
	return res.val, true
}

func (c *Cache) reap(t time.Duration) {
	c.mux.Lock()
	for k, v := range c.cache {
		if time.Since(v.createdAt) >= t {
			delete(c.cache, k)
		}
	}
	c.mux.Unlock()
}

func (c *Cache) reapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		c.reap(t)
	}

}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}
