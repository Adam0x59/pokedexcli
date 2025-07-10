package pokecache

import (
	"time"
)

// NewCache creates a new Cache with a reaping goroutine.
func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		interval: interval,
		cache:    make(map[string]cacheEntry),
	}
	// Start the reapLoop in a background goroutine.
	// Use go to run the reapLoop concurrently (In it's own thread.)
	go newCache.reapLoop()
	return &newCache
}

// Add puts a value into the cache under the provided key.
func (c *Cache) Add(key string, value []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = entry
}

// Get retrieves an entry from the cache.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.cache[key]
	if exists {
		return value.val, exists
	} else {
		return nil, exists
	}
}

// reapLoop periodically deletes expired entries from the cache.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for i := range c.cache {
			if time.Since(c.cache[i].createdAt) > c.interval {
				delete(c.cache, i)
			}
		}
		c.mu.Unlock()
	}
}
