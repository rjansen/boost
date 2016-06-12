package cache_test

import (
	"testing"

	"farm.e-pedion.com/repo/cache"
	"farm.e-pedion.com/repo/config"
	"github.com/stretchr/testify/assert"
)

var (
	cacheClient *cache.Client
	setted      = false
	key1        = "8b06603b-9b0d-4e8c-8aae-10f988639fe6"
	expires     = 60
)

func setup() error {
	setupErr := cache.Setup(&config.CacheConfig{
		CacheAddress: "127.0.0.1:11211",
	})
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
	return nil
}

func Test_NewClient(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	cacheClient = cache.NewClient()
	assert.NotNil(t, cacheClient)
}

func Test_AddItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	err := cacheClient.Add(key1, expires, []byte("1234567890"))
	assert.Nil(t, err)
}

func Test_GetItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	item, err := cacheClient.Get(key1)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func Test_DelItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	err := cacheClient.Delete(key1)
	assert.Nil(t, err)
}

func Test_GetEmptyItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	item, err := cacheClient.Get(key1)
	assert.NotNil(t, err)
	assert.Nil(t, item)
}

func Test_SetItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	err := cacheClient.Set("cache_test", 120, []byte("golang test"))
	assert.Nil(t, err)
}
