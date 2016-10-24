package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

//NewClient creates a new instance of the cache client component
func NewClient() Client {
	if configuration == nil {
		panic("cache.Setup never called")
	}
	client := &MemcachedClient{
		Client: memcache.New(configuration.URL),
	}
	return client
}

//MemcachedClient is a component to interact with a cache system
type MemcachedClient struct {
	*memcache.Client
}

//Get reads the value associated with the provided key
func (c *MemcachedClient) Get(key string) ([]byte, error) {
	item, err := c.Client.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (c *MemcachedClient) Add(key string, expires int, item []byte) error {
	return c.Client.Add(&memcache.Item{Key: key, Expiration: int32(expires), Value: item})
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (c *MemcachedClient) Set(key string, expires int, item []byte) error {
	return c.Client.Set(&memcache.Item{Key: key, Expiration: int32(expires), Value: item})
}

//Delete removes the item associated with the provided key
func (c *MemcachedClient) Delete(key string) error {
	return c.Client.Delete(key)
}
