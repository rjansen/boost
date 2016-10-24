package cache

import (
	testify "github.com/stretchr/testify/mock"
)

//NewMockClient creates a new mock instance of the cache client component
func NewMockClient() *MockClient {
	return &MockClient{}
}

//MockClient is a mock client for cache system
type MockClient struct {
	testify.Mock
}

//Get reads the value associated with the provided key
func (m *MockClient) Get(key string) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (m *MockClient) Add(key string, expires int, item []byte) error {
	args := m.Called(key, expires, []byte(item))
	return args.Error(0)
}

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (m *MockClient) Set(key string, expires int, item []byte) error {
	args := m.Called(key, expires, []byte(item))
	return args.Error(0)
}

//Delete removes the item associated with the provided key
func (m *MockClient) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}
