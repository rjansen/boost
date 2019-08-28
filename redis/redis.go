package memcached

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//Configuration holds cache connections parameters
type Configuration struct {
	URL      string `json:"url" mapstructure:"url"`
	Password string
	DB       int
}

func (c Configuration) String() string {
	return fmt.Sprintf("redis.Configuration{URL=%v}", c.URL)
}

//Cache is the interface interact with the memcached
type Cache interface {
	//Ping validate the cache connection
	Ping() *redis.StatusCmd
	//Get reads the value associated with the provided key
	Get(string) *redis.StringCmd
	//SetNX inserts a new item in the cache, Add throws error if the provided key was already defined
	SetNX(string, interface{}, time.Duration) *redis.BoolCmd
	//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	//Del removes the item associated with the provided keys
	Del(...string) *redis.IntCmd
	//Close releases all resources and terminate this client
	Close() error
}
