package boost

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type clientPoolMock struct{}

func (c *clientPoolMock) Get() (Client, error) { return nil, nil }

func (c *clientPoolMock) Close() error { return nil }

func TestClientPool(t *testing.T) {
	pool := new(clientPoolMock)
	assert.NotNil(t, pool)
	assert.Implements(t, (*ClientPool)(nil), pool)
}

type clientMock struct{}

//Get reads the value associated with the provided key
func (c *clientMock) Get(key string) ([]byte, error) { return nil, nil }

//Add inserts a new item in the cache, Add throws error if the provided key was already defined
func (c *clientMock) Add(key string, expires time.Duration, item []byte) error { return nil }

//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
func (c *clientMock) Set(key string, expires time.Duration, item []byte) error { return nil }

//Delete removes the item associated with the provided key
func (c *clientMock) Delete(key string) error { return nil }

//Close closes the client connection
func (c *clientMock) Close() error { return nil }

func TestClient(t *testing.T) {
	client := new(clientMock)
	assert.NotNil(t, client)
	assert.Implements(t, (*Client)(nil), client)
}
