package caching

type Cache interface {
	GetCacheName() string
	Initialize()
	GetProgram(name string) Program
	SetProgram(name string, prog Program)
	Flush()
}
