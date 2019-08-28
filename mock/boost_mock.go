package boost

import (
	"time"

	"github.com/rjansen/boost"
	testify "github.com/stretchr/testify/mock"
)

func NewClientPoolMock() *ClientPoolMock {
	return new(ClientPoolMock)
}

//ClientPoolMock is a mock for a cache client pool
type ClientPoolMock struct {
	testify.Mock
}

//Get returns a cache Client instance
func (m *ClientPoolMock) Get() (boost.Client, error) {
	args := m.Called()
	result := args.Get(0)
	if result != nil {
		return result.(boost.Client), args.Error(1)
	}
	return nil, args.Error(1)
}

//Close finalizes the pool instance
func (m *ClientPoolMock) Close() error {
	args := m.Called()
	return args.Error(0)
}

//NewClientMock creates a new mock instance of the cache client component
func NewClientMock() *ClientMock {
	return new(ClientMock)
}

//MockClient is a mock client for cache system
type ClientMock struct {
	testify.Mock
}

//Get reads the value associated with the provided key
func (m *ClientMock) Get(key string) ([]byte, error) {
	args := m.Called(key)
	result := args.Get(0)
	if result != nil {
		return result.([]byte), args.Error(1)
	}
	return nil, args.Error(1)
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (m *ClientMock) Add(key string, expires time.Duration, item []byte) error {
	args := m.Called(key, expires, []byte(item))
	return args.Error(0)
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (m *ClientMock) Set(key string, expires time.Duration, item []byte) error {
	args := m.Called(key, expires, []byte(item))
	return args.Error(0)
}

//Delete removes the item associated with the provided key
func (m *ClientMock) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

//Close finalizes the client instance
func (m *ClientMock) Close() error {
	args := m.Called()
	return args.Error(0)
}
