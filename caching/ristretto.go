package caching

import (
	"fmt"

	"github.com/dgraph-io/ristretto"
)

type RistrettoCache struct {
	cache *ristretto.Cache
}

func (c *RistrettoCache) GetCacheName() string {
	return "Ristretto"
}

func (c *RistrettoCache) Initialize() {
	var err error
	c.cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // Num keys to track frequency of (10M).
		MaxCost:     1 << 30, // Maximum cost of cache (1GB).
		BufferItems: 64,      // Number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
}

func (c *RistrettoCache) GetProgram(name string) Program {
	var prog Program

	value, found := c.cache.Get(name)
	if found {
		prog = value.(Program)
	}

	return prog
}

func (r *RistrettoCache) SetProgram(name string, value Program) {
	r.SetProgramWithCost(name, value, 1)
}

func (c *RistrettoCache) SetProgramWithCost(name string, value Program, cost int64) {
	c.cache.Set(name, value, cost)
	c.cache.Wait()

	// _, found := c.cache.Get(name)
	// if !found {
	// 	panic("missing value")
	// }
}

func (c *RistrettoCache) Flush() {
	c.cache.Clear()
	fmt.Println("Ristretto is flushed")
}
