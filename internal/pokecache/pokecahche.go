package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheEntries map[string]cacheEntry
	mu           sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) cacheEntry {
	return cacheEntry{
		createdAt: time.Now(),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.CacheEntries[key] = cacheEntry{val: val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, isFound := c.CacheEntries[key]
	val := cacheEntry.val
	return val, isFound
}
