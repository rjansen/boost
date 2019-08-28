package memcached

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/rjansen/boost"
	"github.com/rjansen/l"
)

var (
	ErrKeyAlreadyExists = errors.New("ErrKeyAlreadyExists{Message='The provided key was already setted'}")
)

//Pool controls how new gocql.Session will create and maintained
type Pool struct {
	config Configuration
	client *Client
}

func NewPool(config Configuration) *Pool {
	return &Pool{
		config: config,
		client: NewClient(config),
	}
}

func (c Pool) String() string {
	return fmt.Sprintf("redis.Pool{config=%s client=%t}",
		c.config.String(), c.client == nil,
	)
}

//Get creates and returns a Client reference
func (c *Pool) Get() (boost.Client, error) {
	l.Debug(context.Background(), "redis.Get",
		l.NewValue("Pool", c.String()),
		l.NewValue("ClientIsNil", c.client == nil),
	)
	return c.client, nil
}

//Close close the database pool
func (c *Pool) Close() error {
	l.Info(context.Background(), "redis.Close",
		l.NewValue("RedisPool", c.String()),
	)
	c.client.Close()
	return nil
}

//NewClient creates a new instance of the cache client component
func NewClient(config Configuration) *Client {
	return &Client{
		config: config,
		cache: redis.NewClient(
			&redis.Options{
				Addr:     config.URL,
				Password: config.Password,
				DB:       0, // use default DB
			},
		),
	}
}

//Client is a component to interact with a cache system
type Client struct {
	config Configuration
	cache  Cache
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
func (c Client) Add(key string, expires time.Duration, item []byte) error {
	/*
		_, err := c.Get(key)
		if err != nil {
			if err == redis.KeyNotFound {
				return ErrAlreadyExists
			}

			return err
		}
	*/

	return c.cache.SetNX(key, item, expires).Err()
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (c Client) Set(key string, expires time.Duration, item []byte) error {
	return c.cache.Set(key, item, expires).Err()
}

//Delete removes the item associated with the provided key
func (c Client) Delete(key string) error {
	return c.cache.Del(key).Err()
}

//Close terminates the memcached connection
func (c Client) Close() error {
	return c.cache.Close()
}
