package boost

import (
	"testing"
	"time"

	"github.com/rjansen/boost"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClientPoolMock(t *testing.T) {
	pool := NewClientPoolMock()
	assert.NotNil(t, pool)
	assert.Implements(t, (*boost.ClientPool)(nil), pool)

	pool.On("Get").Return(nil, nil).Once()
	pool.On("Close").Return(nil).Once()
	pool.Get()
	pool.Close()

	pool.AssertExpectations(t)

	client := NewClientMock()
	pool.On("Get").Return(client, nil).Once()
	poolClient, err := pool.Get()
	assert.Nil(t, err)
	assert.Exactly(t, client, poolClient)

	pool.AssertExpectations(t)
}

func TestClient(t *testing.T) {
	client := NewClientMock()
	assert.NotNil(t, client)
	assert.Implements(t, (*boost.Client)(nil), client)

	client.On("Add", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	client.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	client.On("Get", mock.Anything).Return(nil, nil).Once()
	client.On("Delete", mock.Anything).Return(nil).Once()
	client.On("Close").Return(nil).Once()
	client.Add("key", time.Second, []byte("value"))
	client.Set("key", time.Second, []byte("value"))
	client.Get("key")
	client.Delete("key")
	client.Close()

	client.AssertExpectations(t)

	mockBytes := []byte("mock_bytes")
	client.On("Get", mock.Anything).Return(mockBytes, nil).Once()
	getBytes, err := client.Get("key")
	assert.Nil(t, err)
	assert.Exactly(t, mockBytes, getBytes)

	client.AssertExpectations(t)
}
