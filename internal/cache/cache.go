package cache

import (
	"errors"
	"sync"
)

// Cache is a struct for storage any value by unique key
type Cache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

// New initialize and return cache pointer
func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set add to storage new value by key
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

// Get return stored value by key if it exists
// Otherwise, returns an error
func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exist := c.data[key]
	if !exist {
		return nil, errors.New("get: unknown key")
	}

	return value, nil
}

// Delete remove value from storage by key if it exists
// Otherwise, returns an error
func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exist := c.data[key]
	if !exist {
		return errors.New("delete: unknown key")
	}
	delete(c.data, key)

	return nil
}
