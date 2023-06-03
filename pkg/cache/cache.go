package cache

import (
	"log"
	"sync"
	"time"
)

type Entry struct {
	value      interface{}
	expiration time.Time
}

type Cache struct {
	lock   sync.RWMutex
	data   map[string]Entry
	ticker *time.Ticker
}

func NewCache() *Cache {
	cache := &Cache{
		data:   make(map[string]Entry),
		ticker: time.NewTicker(time.Second),
	}

	go cache.gc()

	return cache
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}

	return entry.value, true
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	var (
		expiration = time.Now().Add(ttl)
		entry      = Entry{
			value:      value,
			expiration: expiration,
		}
	)

	c.data[key] = entry
}

func (c *Cache) gc() {
	for range c.ticker.C {
		c.lock.Lock()

		for key, entry := range c.data {
			if time.Now().After(entry.expiration) {
				log.Printf("collecting %s", key)
				delete(c.data, key)
			}
		}

		c.lock.Unlock()
	}
}
