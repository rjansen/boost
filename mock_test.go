package cache

import (
	testify "github.com/stretchr/testify/mock"
)

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
func (m *ClientMock) Add(key string, expires int, item []byte) error {
	args := m.Called(key, expires, []byte(item))
	return args.Error(0)
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (m *ClientMock) Set(key string, expires int, item []byte) error {
	args := m.Called(key, expires, []byte(item))
	return args.Error(0)
}

//Delete removes the item associated with the provided key
func (m *ClientMock) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}
