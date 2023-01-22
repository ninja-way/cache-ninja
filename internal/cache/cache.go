package cache

import "errors"

type Cache struct {
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, exist := c.data[key]
	if !exist {
		return nil, errors.New("get: unknown key")
	}

	return value, nil
}

func (c *Cache) Delete(key string) error {
	_, exist := c.data[key]
	if !exist {
		return errors.New("delete: unknown key")
	}
	delete(c.data, key)
	return nil
}
