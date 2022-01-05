package caching

import (
	"encoding/json"
	"fmt"

	"github.com/coocood/freecache"
)

type FreeCache struct {
	cache *freecache.Cache
}

func (c *FreeCache) GetCacheName() string {
	return "FreeCache"
}

func (c *FreeCache) Initialize() {
	cacheSize := 100 * 1024 * 1024
	c.cache = freecache.NewCache(cacheSize)
}

func (c *FreeCache) GetProgram(name string) Program {
	var prog Program

	key := []byte(name)
	byteValue, err := c.cache.Get(key)
	if err == nil {
		json.Unmarshal(byteValue, &prog)
	} else if err.Error() != "Entry not found" {
		fmt.Printf("Error is: %s", err)
		panic("Unable to get Program data")
	}

	return prog
}

func (c *FreeCache) SetProgram(name string, prog Program) {
	value, err := json.Marshal(prog)
	if err != nil {
		panic("Unable to marshall Program data")
	}

	key := []byte(name)
	val := []byte(value)
	c.cache.Set(key, val, 1)
}

func (c *FreeCache) Flush() {
	c.cache.Clear()
	fmt.Println("Freecache is flushed")
}
