package memcached

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

//Configuration holds cache connections parameters
type Configuration struct {
	URL string `json:"url" mapstructure:"url"`
}

func (c Configuration) String() string {
	return fmt.Sprintf("memcached.Configuration{URL=%v}", c.URL)
}

//Cache is the interface interact with the memcached
type Cache interface {
	//Get reads the value associated with the provided key
	Get(key string) (*memcache.Item, error)
	//Add inserts a new item in the cache, Add throws error if the provided key was already defined
	Add(item *memcache.Item) error
	//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
	Set(item *memcache.Item) error
	//Delete removes the item associated with the provided key
	Delete(key string) error
}
