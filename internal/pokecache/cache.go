package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheMap: make(map[string]cacheEntry),
		interval: interval,
		mu: &sync.Mutex{},
	}
	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cacheMap[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for {
	<-ticker.C

	c.mu.Lock()
		for key, entry := range c.cacheMap {
			if time.Now().Sub(entry.createdAt) > c.interval{
				delete(c.cacheMap, key)
			}
		}			
	c.mu.Unlock()
	}	
}