package cache

import "errors"

// Cache is a struct for storage any value by unique key
type Cache struct {
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
	c.data[key] = value
}

// Get return stored value by key if it exists
// Otherwise, returns an error
func (c *Cache) Get(key string) (interface{}, error) {
	value, exist := c.data[key]
	if !exist {
		return nil, errors.New("get: unknown key")
	}

	return value, nil
}

// Delete remove value from storage by key if it exists
// Otherwise, returns an error
func (c *Cache) Delete(key string) error {
	_, exist := c.data[key]
	if !exist {
		return errors.New("delete: unknown key")
	}
	delete(c.data, key)
	return nil
}
