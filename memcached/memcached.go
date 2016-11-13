package memcached

import (
	"farm.e-pedion.com/repo/cache"
	"farm.e-pedion.com/repo/logger"
	"github.com/bradfitz/gomemcache/memcache"
)

//GoMemcacheClient is the interface interact with the memcached
type GoMemcacheClient interface {
	//Get reads the value associated with the provided key
	Get(key string) (*memcache.Item, error)
	//Add inserts a new item in the cache, Add throws error if the provided key was already defined
	Add(item *memcache.Item) error
	//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
	Set(item *memcache.Item) error
	//Delete removes the item associated with the provided key
	Delete(key string) error
}

//NewClient creates a new instance of the cache client component
func NewClient() *MemcachedClient {
	if cache.Config == nil {
		logger.Panic("cache.Setup never called")
	}
	client := &MemcachedClient{
		client: memcache.New(cache.Config.URL),
	}
	return client
}

//MemcachedClient is a component to interact with a cache system
type MemcachedClient struct {
	client GoMemcacheClient
}

//Get reads the value associated with the provided key
func (c *MemcachedClient) Get(key string) ([]byte, error) {
	item, err := c.client.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (c *MemcachedClient) Add(key string, expires int, item []byte) error {
	return c.client.Add(
		&memcache.Item{
			Key:        key,
			Expiration: int32(expires),
			Value:      item,
		},
	)
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (c *MemcachedClient) Set(key string, expires int, item []byte) error {
	return c.client.Set(
		&memcache.Item{
			Key:        key,
			Expiration: int32(expires),
			Value:      item,
		},
	)
}

//Delete removes the item associated with the provided key
func (c *MemcachedClient) Delete(key string) error {
	return c.client.Delete(key)
}
