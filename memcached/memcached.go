package memcached

import (
	"farm.e-pedion.com/repo/logger"
	"github.com/bradfitz/gomemcache/memcache"
)

//NewClient creates a new instance of the cache client component
func NewClient() *Client {
	if Config == nil {
		logger.Panic("memcached.Setup never called")
	}
	client := &Client{
		cache: memcache.New(Config.URL),
	}
	return client
}

//Client is a component to interact with a cache system
type Client struct {
	cache Cache
}

//Get reads the value associated with the provided key
func (c Client) Get(key string) ([]byte, error) {
	item, err := c.cache.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (c Client) Add(key string, expires int, item []byte) error {
	return c.cache.Add(
		&memcache.Item{
			Key:        key,
			Expiration: int32(expires),
			Value:      item,
		},
	)
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (c Client) Set(key string, expires int, item []byte) error {
	return c.cache.Set(
		&memcache.Item{
			Key:        key,
			Expiration: int32(expires),
			Value:      item,
		},
	)
}

//Delete removes the item associated with the provided key
func (c Client) Delete(key string) error {
	return c.cache.Delete(key)
}

//Close terminates the memcached connection
func (c Client) Close() error {
	return nil
}
