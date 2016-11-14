package itest

import (
	. "farm.e-pedion.com/repo/cache/memcached"
	"farm.e-pedion.com/repo/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	cacheClient *Client
	setted      = false
	key1        = "8b06603b-9b0d-4e8c-8aae-10f988639fe6"
	expires     = 60
	testConfig  *Configuration
)

func init() {
	os.Args = append(os.Args, "-ecf", "etc/cache/cache.yaml")
	if err := config.UnmarshalKey("cache.memcached", &testConfig); err != nil {
		panic(err)
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
	return nil
}

func TestIntegrationNewClient(t *testing.T) {
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	cacheClient = NewClient()
	assert.NotNil(t, cacheClient)
}

func TestIntegrationAddItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	err := cacheClient.Add(key1, expires, []byte("1234567890"))
	assert.Nil(t, err)
}

func TestIntegrationGetItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	item, err := cacheClient.Get(key1)
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestIntegrationDelItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	err := cacheClient.Delete(key1)
	assert.Nil(t, err)
}

func TestIntegrationGetEmptyItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	item, err := cacheClient.Get(key1)
	assert.NotNil(t, err)
	assert.Nil(t, item)
}

func TestIntegrationSetItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	err := cacheClient.Set("cache_test", 120, []byte("golang test"))
	assert.Nil(t, err)
}
