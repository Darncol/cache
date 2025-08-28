package cache

import "sync"

type Cache struct {
	cache map[string]interface{}
	mu    *sync.RWMutex
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.cache[key] = value
	c.mu.Unlock()
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache[key]
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.cache, key)
	c.mu.Unlock()
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}
