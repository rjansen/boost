package memcached

import (
	"github.com/go-redis/redis"
	testify "github.com/stretchr/testify/mock"
	"time"
)

//NewCacheMock creates a new mock instance of the cache client component
func NewCacheMock() *CacheMock {
	return new(CacheMock)
}

//ClientMock is a mock client for memcached
type CacheMock struct {
	testify.Mock
}

//Ping validates the connection with the cache
func (m *CacheMock) Ping() *redis.StatusCmd {
	args := m.Called()
	result := args.Get(0)
	if result != nil {
		return result.(*redis.StatusCmd)
	}
	return nil
}

//Get reads the value associated with the provided key
func (m *CacheMock) Get(key string) *redis.StringCmd {
	args := m.Called(key)
	result := args.Get(0)
	if result != nil {
		return result.(*redis.StringCmd)
	}
	return nil
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (m *CacheMock) SetNX(key string, value interface{}, expires time.Duration) *redis.BoolCmd {
	args := m.Called(key, value, expires)
	result := args.Get(0)
	if result != nil {
		return result.(*redis.BoolCmd)
	}
	return nil
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (m *CacheMock) Set(key string, value interface{}, expires time.Duration) *redis.StatusCmd {
	args := m.Called(key, value, expires)
	result := args.Get(0)
	if result != nil {
		return result.(*redis.StatusCmd)
	}
	return nil
}

//Delete removes the item associated with the provided key
func (m *CacheMock) Del(keys ...string) *redis.IntCmd {
	args := m.Called(keys)
	result := args.Get(0)
	if result != nil {
		return result.(*redis.IntCmd)
	}
	return nil
}

//Close releases all resources with the cache and terminate this client
func (m *CacheMock) Close() error {
	args := m.Called()
	return args.Error(0)
}
