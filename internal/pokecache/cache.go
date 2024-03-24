package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntry[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.cacheEntry[key]; ok {
		return entry.val, true
	}

	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key := range c.cacheEntry {
		end := time.Now()
		difference := end.Sub(c.cacheEntry[key].createdAt)
		if difference > interval {
			delete(c.cacheEntry, key)
			fmt.Println(c.cacheEntry)
		}
	}
}

type Time struct {
	interval time.Duration
}

type Ticker struct {
	C <-chan Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheEntry: map[string]cacheEntry{},
		mu:         &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}
