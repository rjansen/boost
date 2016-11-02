package cache

import (
	"context"
	"errors"
	"farm.e-pedion.com/repo/config"
	"farm.e-pedion.com/repo/logger"
	"fmt"
)

var (
	configuration     *Configuration
	cacheClientKey    = 30
	ErrInvalidContext = errors.New("The provided Context is invalid")
)

//Setup configures the cache package
func Setup() error {
	if err := config.UnmarshalKey("cache", &configuration); err != nil {
		logger.Error("cache.GetConfigErr", logger.Err(err))
		return err
	}
	return nil
}

//Client is a component to interact with a cache system
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

func SetClient(c context.Context) (context.Context, error) {
	if c == nil {
		return nil, ErrInvalidContext
	}
	cacheClient := NewClient()
	return context.WithValue(c, cacheClientKey, cacheClient), nil
}
