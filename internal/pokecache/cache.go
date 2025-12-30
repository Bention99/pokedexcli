package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mu sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache() *Cache {
	return &Cache{
		entries: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()

	if _, exists := c.entries[key]; exists {
		return
	}

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
}

func (c *Cache) Get(key string) []byte, bool {
	c.mu.Lock()
	entry, ok := c.entries[key]
	c.mu.Unlock()

	if !ok {
		return nil, false
	}

	return entry.val, true
}