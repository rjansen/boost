package memcached

import (
	"context"
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rjansen/boost"
	"github.com/rjansen/l"
)

//Pool controls how new gocql.Session will create and maintained
type Pool struct {
	config Configuration
	client *Client
}

func NewPool(config Configuration) *Pool {
	pool := &Pool{
		config: config,
		client: NewClient(config),
	}
	return pool
}

func (c Pool) String() string {
	return fmt.Sprintf("memcached.Pool{config=%s, client=%t}", c.config.String(), c.client == nil)
}

//Get creates and returns a Client reference
func (c *Pool) Get() (boost.Client, error) {
	l.Debug(context.Background(), "memcached.Get",
		l.NewValue("Pool", c.String()),
		l.NewValue("ClientIsNil", c.client == nil),
	)
	return c.client, nil
}

//Close close the database pool
func (c *Pool) Close() error {
	l.Info(context.Background(), "CloseMemcacheClient",
		l.NewValue("MemcachedPool", c.String()),
	)
	c.client.Close()
	return nil
}

//NewClient creates a new instance of the cache client component
func NewClient(config Configuration) *Client {
	client := &Client{
		config: config,
		cache:  memcache.New(config.URL),
	}
	return client
}

//Client is a component to interact with a cache system
type Client struct {
	config Configuration
	cache  Cache
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
func (c Client) Add(key string, expires time.Duration, item []byte) error {
	return c.cache.Add(
		&memcache.Item{Key: key, Expiration: int32(expires), Value: item},
	)
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (c Client) Set(key string, expires time.Duration, item []byte) error {
	return c.cache.Set(
		&memcache.Item{Key: key, Expiration: int32(expires), Value: item},
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
