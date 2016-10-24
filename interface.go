package cache

import (
	"farm.e-pedion.com/repo/config"
	"farm.e-pedion.com/repo/logger"
)

var (
	configuration *Configuration
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
