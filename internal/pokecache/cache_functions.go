package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()

	return &cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}

		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cache := cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}

	c.cache[key] = cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, exists := c.cache[key]
	if !exists {
		return nil, false
	}

	return data.val, true
}
