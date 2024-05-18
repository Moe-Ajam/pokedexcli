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

func NewCache(interval time.Duration) Cache {
	c := Cache{
		CacheEntries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)
	return c
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

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)

	for key, cacheEntry := range c.CacheEntries {
		if cacheEntry.createdAt.Before(timeAgo) {
			delete(c.CacheEntries, key)
		}
	}
}
