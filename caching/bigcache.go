package caching

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
)

type BigCache struct {
	cache *bigcache.BigCache
}

func (c *BigCache) GetCacheName() string {
	return "BigCache"
}

func (c *BigCache) Initialize() {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		panic("Unable to create BigCache")
	}

	c.cache = cache
}

func (c *BigCache) GetProgram(name string) Program {
	var prog Program

	byteValue, err := c.cache.Get(name)
	if err == nil {
		json.Unmarshal(byteValue, &prog)
	} else if err.Error() != "Entry not found" {
		fmt.Printf("Error is: %s", err)
		panic("Unable to get Program data")
	}

	return prog
}

func (c *BigCache) SetProgram(name string, prog Program) {
	value, err := json.Marshal(prog)
	if err != nil {
		panic("Unable to marshall Program data")
	}

	val := []byte(value)
	c.cache.Set(name, val)
}

func (c *BigCache) Flush() {
	c.cache.Reset()
	fmt.Println("BigCache is flushed")
}
