package pokecache

import (
	"sync"
	"time"
)

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu: &sync.RWMutex{},
	}

	go cache.reapLoop(duration)

	return cache
}