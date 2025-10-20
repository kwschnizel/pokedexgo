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
	mux      sync.Mutex
	cacheMap map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		mux:      sync.Mutex{},
		cacheMap: make(map[string]cacheEntry),
	}

	go newCache.reapLoop(interval)

	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cacheMap[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	res, ok := c.cacheMap[key]
	c.mux.Unlock()

	return res.val, ok
}

// TODO: Implement a better way to clean the cache because this is awful
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		for k, v := range c.cacheMap {
			if time.Since(v.createdAt) > interval {
				delete(c.cacheMap, k)
			}
		}
		c.mux.Unlock()
	}
}
