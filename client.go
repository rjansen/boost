package cache

import (
	"context"
	"errors"
	"fmt"
)

var (
	//Config is the cache package configuraton
	Config            *Configuration
	cacheClientKey    = 30
	ErrInvalidContext = errors.New("The provided Context is invalid")
	ErrInvalidConfig  = errors.New("The provided Configuration is invalid")
)

//Setup configures the cache package
func Setup(config *Configuration) error {
	if config == nil {
		return ErrInvalidConfig
	}
	Config = config
	return nil
}

//Configuration holds cache connections parameters
type Configuration struct {
	Provider string `mapstructure:"provider"`
	URL      string `mapstructure:"url"`
}

func (c Configuration) String() string {
	return fmt.Sprintf("cache.Configuration Provider=%s URL=%v", c.Provider, c.URL)
}

//Client is an interface to interact with the cache
type Client interface {
	//Get reads the value associated with the provided key
	Get(key string) ([]byte, error)
	//Add inserts a new item in the cache, Add throws error if the provided key was already defined
	Add(key string, expires int, item []byte) error
	//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
	Set(key string, expires int, item []byte) error
	//Delete removes the item associated with the provided key
	Delete(key string) error
}

func GetClient(c context.Context) (Client, error) {
	if c == nil {
		return nil, ErrInvalidContext
	}
	cacheClient, ok := c.Value(cacheClientKey).(Client)
	if !ok {
		return nil, fmt.Errorf("ErrInvalidCacheClient client=%+v", cacheClient)
	}
	return cacheClient, nil
}

func SetClient(c context.Context, cacheClient Client) (context.Context, error) {
	if c == nil {
		return nil, ErrInvalidContext
	}
	return context.WithValue(c, cacheClientKey, cacheClient), nil
}
