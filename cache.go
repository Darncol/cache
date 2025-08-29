package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]interface{}
	timeLeft map[string]chan struct{}
	mu       *sync.RWMutex
}

func (c *Cache) uninstaller(timeLeft time.Duration, key string, timer chan struct{}) {
	for {
		select {
		case <-timer:
			return
		case <-time.After(timeLeft):
			c.Delete(key)
			return
		}
	}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()

	oldTimer, ok := c.timeLeft[key]
	if ok {
		oldTimer <- struct{}{}
	}

	timer := make(chan struct{})

	c.timeLeft[key] = timer
	c.cache[key] = value

	go c.uninstaller(duration, key, timer)

	c.mu.Unlock()
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache[key]
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	timer, ok := c.timeLeft[key]
	if ok {
		timer <- struct{}{}
	}

	close(timer)

	delete(c.cache, key)
	delete(c.timeLeft, key)
}

func New() *Cache {
	return &Cache{
		cache:    make(map[string]interface{}),
		timeLeft: make(map[string]chan struct{}),
		mu:       &sync.RWMutex{},
	}
}
