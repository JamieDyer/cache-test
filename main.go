package main

import (
	"caching/caching"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var ristretto caching.RistrettoCache
	var storage caching.Storage
	var goCache caching.GoCache
	var freeCache caching.FreeCache
	var bigCache caching.BigCache

	var storageCount, storageTiming int64
	var ristrettoCount, ristrettoTiming int64
	var goCacheCount, goCacheTiming int64
	var freeCacheCount, freeCacheTiming int64
	var bigCacheCount, bigCacheTiming int64

	var cacheHits = 30000
	var programNames = []string{"Skull Mountain", "Raspberry", "Butterfly"}

	storage.Initialize()
	ristretto.Initialize()
	goCache.Initialize()
	freeCache.Initialize()
	bigCache.Initialize()

	for i := 0; i < cacheHits; i++ {
		rand.Seed(time.Now().UnixNano())
		randoCache := rand.Intn(5)
		randoKey := rand.Intn(3)
		if randoCache == 0 {
			storageCount++
			_, storageTimingTemp := GetProgramFromSomewhere(&storage, programNames[randoKey])
			storageTiming += storageTimingTemp
		} else if randoCache == 1 {
			ristrettoCount++
			_, ristrettoTimingTemp := GetProgramFromSomewhere(&ristretto, programNames[randoKey])
			ristrettoTiming += ristrettoTimingTemp
		} else if randoCache == 2 {
			goCacheCount++
			_, goCacheTimingTemp := GetProgramFromSomewhere(&goCache, programNames[randoKey])
			goCacheTiming += goCacheTimingTemp
		} else if randoCache == 3 {
			freeCacheCount++
			_, freeCacheTimingTemp := GetProgramFromSomewhere(&freeCache, programNames[randoKey])
			freeCacheTiming += freeCacheTimingTemp
		} else {
			bigCacheCount++
			_, bigCacheTimingTemp := GetProgramFromSomewhere(&bigCache, programNames[randoKey])
			bigCacheTiming += bigCacheTimingTemp
		}

		// Flush halfway through
		if i == cacheHits/2 {
			ristretto.Flush()
			goCache.Flush()
			freeCache.Flush()
			bigCache.Flush()
		}
	}

	// Print a program to prove the cache actually has the correct data???

	fmt.Printf("Storage %d interations, time taken: %d micro-seconds\n", storageCount, storageTiming)
	fmt.Printf("Ristretto %d interations, time taken: %d micro-seconds\n", ristrettoCount, ristrettoTiming)
	fmt.Printf("Go-Cache %d interations, time taken: %d micro-seconds\n", goCacheCount, goCacheTiming)
	fmt.Printf("FreeCache %d interations, time taken: %d micro-seconds\n", freeCacheCount, freeCacheTiming)
	fmt.Printf("BigCache %d interations, time taken: %d micro-seconds\n", bigCacheCount, bigCacheTiming)
}

func GetProgramFromSomewhere(cache caching.Cache, key string) (caching.Program, int64) {
	start := time.Now()

	prog := cache.GetProgram(key)
	if prog.ID == "" {
		fmt.Printf("Cache miss for %s in: %s\n", key, cache.GetCacheName())
		// Get from the storage
		var fileStore caching.Storage
		fileStore.Initialize()
		prog = fileStore.GetProgram(key)

		cache.SetProgram(key, prog)
	}

	return prog, time.Since(start).Microseconds()
}
