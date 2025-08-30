package cache

import (
	"context"
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]interface{}
	timeLeft map[string]context.CancelFunc
	mu       *sync.RWMutex
}

func (c *Cache) uninstaller(ctx context.Context, duration time.Duration, key string) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(duration):
			c.Delete(key)
			return
		}
	}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cancel, ok := c.timeLeft[key]
	if ok {
		cancel()
		delete(c.timeLeft, key)
	}

	ctx, cancel := context.WithCancel(context.Background())

	c.timeLeft[key] = cancel
	c.cache[key] = value

	go c.uninstaller(ctx, duration, key)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exist := c.cache[key]
	return value, exist
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cancel, ok := c.timeLeft[key]
	if ok {
		cancel()
		delete(c.timeLeft, key)
	}

	delete(c.cache, key)
}

func New() *Cache {
	return &Cache{
		cache:    make(map[string]interface{}),
		timeLeft: make(map[string]context.CancelFunc),
		mu:       &sync.RWMutex{},
	}
}
