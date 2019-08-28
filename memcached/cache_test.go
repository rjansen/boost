package memcached

import (
	"errors"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	cacheClient    = new(Client)
	memcacheClient *CacheMock
	key1           = "8b06603b-9b0d-4e8c-8aae-10f988639fe6"
	expires        = time.Second * 60
	testConfig     = &Configuration{
		URL: "mock://cache",
	}
)

func before() error {
	memcacheClient = NewCacheMock()
	cacheClient.cache = memcacheClient
	return nil
}

func TestNewClient(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotPanics(t,
		func() {
			NewClient(*testConfig)
		},
	)
}

func TestNewPool(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	pool := NewPool(*testConfig)
	assert.NotNil(t, pool)
	client, err := pool.Get()
	assert.Nil(t, err)
	assert.NotNil(t, client)
	err = client.Close()
	assert.Nil(t, err)
	err = pool.Close()
	assert.Nil(t, err)
}

func TestAddItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Add", mock.Anything).Return(nil)
	err := cacheClient.Add(key1, expires, []byte("1234567890"))
	assert.Nil(t, err)
}

func TestGetItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Get", mock.Anything).Return(&memcache.Item{Value: []byte("CacheMock")}, nil)
	item, err := cacheClient.Get(key1)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestDelItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Delete", mock.Anything).Return(nil)
	err := cacheClient.Delete(key1)
	assert.Nil(t, err)
}

func TestGetEmptyItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Get", mock.Anything).Return(nil, errors.New("ErrMockGet"))
	item, err := cacheClient.Get(key1)
	assert.NotNil(t, err)
	assert.Nil(t, item)
}

func TestSetItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	memcacheClient.On("Set", mock.Anything).Return(nil)
	err := cacheClient.Set("cache_test", 120, []byte("golang test"))
	assert.Nil(t, err)
}
