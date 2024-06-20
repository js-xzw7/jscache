package jscache

import (
	"jscache/strategy"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *strategy.LruCache
	cacheBytes int64
}

func (c *cache) Add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		c.lru = strategy.NewLruCache(c.cacheBytes, nil)
	}

	c.lru.Add(key, value)
}

func (c *cache) Get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
