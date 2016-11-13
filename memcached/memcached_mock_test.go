package memcached

import (
	"github.com/bradfitz/gomemcache/memcache"
	testify "github.com/stretchr/testify/mock"
)

//NewGoMemcacheClientMock creates a new mock instance of the cache client component
func NewGoMemcacheClientMock() *GoMemcacheClientMock {
	return new(GoMemcacheClientMock)
}

//GoMemcacheClientMock is a mock client for memcached
type GoMemcacheClientMock struct {
	testify.Mock
}

//Get reads the value associated with the provided key
func (m *GoMemcacheClientMock) Get(key string) (*memcache.Item, error) {
	args := m.Called(key)
	result := args.Get(0)
	if result != nil {
		return result.(*memcache.Item), args.Error(1)
	}
	return nil, args.Error(1)
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (m *GoMemcacheClientMock) Add(item *memcache.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (m *GoMemcacheClientMock) Set(item *memcache.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

//Delete removes the item associated with the provided key
func (m *GoMemcacheClientMock) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}
