package boost

import (
	"context"
	"errors"
	"fmt"
)

var (
	cacheClientKey    = 30
	ErrInvalidState   = errors.New("The cache current state is invalid. Setup never called")
	ErrInvalidContext = errors.New("The provided Context is invalid")
	ErrInvalidConfig  = errors.New("The provided Configuration is invalid")
	pool              ClientPool
)

//Setup configures the cache package
func Setup(p ClientPool) error {
	if p == nil {
		return ErrInvalidConfig
	}
	pool = p
	return nil
}

//Configuration holds cache connections parameters
type Configuration struct {
	Provider string `mapstructure:"provider"`
}

func (c Configuration) String() string {
	return fmt.Sprintf("cache.Configuration Provider=%s", c.Provider)
}

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
	Add(key string, expires int, item []byte) error
	//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
	Set(key string, expires int, item []byte) error
	//Delete removes the item associated with the provided key
	Delete(key string) error
	//Close closes the client connection
	Close() error
}

func GetPool() (ClientPool, error) {
	if pool == nil {
		return nil, ErrInvalidState
	}
	return pool, nil
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
	cacheClient, err := pool.Get()
	if err != nil {
		return nil, err
	}
	return context.WithValue(c, cacheClientKey, cacheClient), nil
}

//ContextFunc is a functions with context olny parameter
type ContextFunc func(context.Context) error

//ExecuteContext preapres a Client and set it inside context to call the provided function
func ExecuteContext(ctxFunc ContextFunc) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, err := SetClient(ctx)
	if err != nil {
		return err
	}
	return ctxFunc(ctx)
}

//ClientFunc is a functions with context olny parameter
type ClientFunc func(Client) error

//Execute gets a Client from the ClientPool and calls the provided function with the Client instance
func Execute(cliFunc ClientFunc) error {
	cacheClient, err := pool.Get()
	if err != nil {
		return err
	}
	defer cacheClient.Close()
	return cliFunc(cacheClient)
}
