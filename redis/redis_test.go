package memcached

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/rjansen/boost"
	"github.com/rjansen/l"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	cacheClient = new(Client)
	redisClient *CacheMock
	setted      = false
	key1        = "8b06603b-9b0d-4e8c-8aae-10f988639fe6"
	expires     = 60
	testConfig  *Configuration
)

func init() {
	if err := l.Setup(&l.Configuration{}); err != nil {
		panic(err)
	}
	testConfig = &Configuration{
		URL: "mock://cache",
	}
}

func setup() error {
	setupErr := Setup(testConfig)
	if setupErr != nil {
		setted = true
	}
	return setupErr
}

func before() error {
	if !setted {
		if err := setup(); err != nil {
			return err
		}
	}
	redisClient = NewCacheMock()
	cacheClient.cache = redisClient
	return nil
}

func TestNewClientPanic(t *testing.T) {
	assert.Panics(t,
		func() {
			NewClient()
		},
	)
}

func TestNewClient(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotPanics(t,
		func() {
			NewClient()
		},
	)
}

func TestPoolGetClient(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	pool, err := boost.GetPool()
	assert.Nil(t, err)
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
	redisClient.On("SetNX", mock.Anything, mock.Anything, mock.Anything).Return(
		redis.NewBoolResult(true, nil),
	)
	err := cacheClient.Add(key1, expires, []byte("1234567890"))
	assert.Nil(t, err)
}

func TestGetItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	redisClient.On("Get", mock.Anything).Return(
		redis.NewStringResult("CacheMock", nil),
	)
	item, err := cacheClient.Get(key1)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestDelItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	redisClient.On("Del", mock.Anything).Return(
		redis.NewIntResult(1, nil),
	)
	err := cacheClient.Delete(key1)
	assert.Nil(t, err)
}

func TestGetEmptyItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	redisClient.On("Get", mock.Anything).Return(
		redis.NewStringResult("", errors.New("ErrMockGet")),
	)
	item, err := cacheClient.Get(key1)
	assert.NotNil(t, err)
	assert.Nil(t, item)
}

func TestSetItem(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	assert.NotNil(t, cacheClient)
	redisClient.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(
		redis.NewStatusResult("OK", nil),
	)
	err := cacheClient.Set("cache_test", 120, []byte("golang test"))
	assert.Nil(t, err)
}
