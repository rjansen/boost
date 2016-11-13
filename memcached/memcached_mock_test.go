package memcached

import (
	"github.com/bradfitz/gomemcache/memcache"
	testify "github.com/stretchr/testify/mock"
)

//NewCacheMock creates a new mock instance of the cache client component
func NewCacheMock() *CacheMock {
	return new(CacheMock)
}

//ClientMock is a mock client for memcached
type CacheMock struct {
	testify.Mock
}

//Get reads the value associated with the provided key
func (m *CacheMock) Get(key string) (*memcache.Item, error) {
	args := m.Called(key)
	result := args.Get(0)
	if result != nil {
		return result.(*memcache.Item), args.Error(1)
	}
	return nil, args.Error(1)
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (m *CacheMock) Add(item *memcache.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (m *CacheMock) Set(item *memcache.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

//Delete removes the item associated with the provided key
func (m *CacheMock) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}
