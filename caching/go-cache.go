package caching

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type GoCache struct {
	gocache *cache.Cache
}

func (c *GoCache) GetCacheName() string {
	return "GoCache"
}

func (c *GoCache) Initialize() {
	defaultExpiry := 1 * time.Minute
	purgeInterval := 60 * time.Minute
	c.gocache = cache.New(defaultExpiry, purgeInterval)
}

func (c *GoCache) GetProgram(name string) Program {
	var prog Program

	value, found := c.gocache.Get(name)
	if found {
		prog = (value.(Program))
	}

	return prog //, found
}

func (c *GoCache) SetProgram(name string, program Program) {
	c.gocache.Set(name, program, cache.DefaultExpiration)
}

func (c *GoCache) Flush() {
	c.gocache.Flush()
	fmt.Println("GoCache is flushed")
}

// g.gocache.Items()
