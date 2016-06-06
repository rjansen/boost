package cache

import (
	"farm.e-pedion.com/repo/config"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	configuration *config.CacheConfig
)

//Client is a component to interact with a cache system
type Client struct {
	*memcache.Client
}

//Get reads the value associated with the provided key
func (c *Client) Get(key string) ([]byte, error) {
	item, err := c.Client.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (c *Client) Add(key string, expires int, item []byte) error {
	return c.Client.Add(&memcache.Item{Key: key, Expiration: int32(expires), Value: item})
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (c *Client) Set(key string, expires int, item []byte) error {
	return c.Client.Set(&memcache.Item{Key: key, Expiration: int32(expires), Value: item})
}

//Delete removes the item associated with the provided key
func (c *Client) Delete(key string) error {
	return c.Client.Delete(key)
}

//Setup configures the cache package
func Setup(cacheConfig *config.CacheConfig) error {
	configuration = cacheConfig
	return nil
}

//NewClient creates a new instance of the cache client component
func NewClient() *Client {
	client := &Client{
		Client: memcache.New(configuration.CacheAddress),
	}
	return client
}
