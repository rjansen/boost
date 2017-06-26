package memcached

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/rjansen/l"
	"time"
)

var (
	ErrKeyAlreadyExists = errors.New("ErrKeyAlreadyExists: Message='The provided key was already setted'")
)

//NewClient creates a new instance of the cache client component
func NewClient() *Client {
	if Config == nil {
		l.Panic("memcached.Setup never called")
	}
	client := &Client{
		cache: redis.NewClient(
			&redis.Options{
				Addr:     Config.URL,
				Password: Config.Password,
				DB:       0, // use default DB
			},
		),
	}
	return client
}

//Client is a component to interact with a cache system
type Client struct {
	cache Cache
}

//Get reads the value associated with the provided key
func (c Client) Get(key string) ([]byte, error) {
	b, err := c.cache.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return b, nil
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (c Client) Add(key string, expires int, item []byte) error {
	created, err := c.cache.SetNX(key, item, time.Second*time.Duration(expires)).Result()
	if err != nil {
		return err
	}
	if !created {
		return ErrKeyAlreadyExists
	}
	return nil
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (c Client) Set(key string, expires int, item []byte) error {
	return c.cache.Set(key, item, time.Second*time.Duration(expires)).Err()
}

//Delete removes the item associated with the provided key
func (c Client) Delete(key string) error {
	return c.cache.Del(key).Err()
}

//Close terminates the memcached connection
func (c Client) Close() error {
	return c.cache.Close()
}
