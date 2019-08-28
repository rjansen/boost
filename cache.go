package boost

import "time"

//ClientPool is an interface for cache pool contract
type ClientPool interface {
	Get() (Client, error)
	Close() error
}

//Client is an interface to interact with the cache
type Client interface {
	//Get reads the value associated with the provided key
	Get(key string) ([]byte, error)
	//Add inserts a new item in the cache, Add throws error if the provided key was already defined
	Add(key string, expires time.Duration, item []byte) error
	//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
	Set(key string, expires time.Duration, item []byte) error
	//Delete removes the item associated with the provided key
	Delete(key string) error
	//Close closes the client connection
	Close() error
}
