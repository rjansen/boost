package cache

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	cacheClient Client
	setted      = false
	key1        = "8b06603b-9b0d-4e8c-8aae-10f988639fe6"
	expires     = 60
)

func init() {
	os.Args = append(os.Args, "-ecf", "./test/etc/cache/cache.yaml")
}

func setup() error {
	setupErr := Setup()
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

func TestNewClient(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	cacheClient = NewClient()
	assert.NotNil(t, cacheClient)
}

func TestAddItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	err := cacheClient.Add(key1, expires, []byte("1234567890"))
	assert.Nil(t, err)
}

func TestGetItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	item, err := cacheClient.Get(key1)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestDelItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	err := cacheClient.Delete(key1)
	assert.Nil(t, err)
}

func TestGetEmptyItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	item, err := cacheClient.Get(key1)
	assert.NotNil(t, err)
	assert.Nil(t, item)
}

func TestSetItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}

	err := cacheClient.Set("cache_test", 120, []byte("golang test"))
	assert.Nil(t, err)
}
