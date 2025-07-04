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
	mu       sync.Mutex
	entry    map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		interval: interval,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for {
		<-ticker.C
		c.mu.Lock()
		for k, v := range c.entry {
			if time.Since(v.createdAt) > c.interval {
				delete(c.entry, k)
			}
		}
		c.mu.Unlock()
	}
}
