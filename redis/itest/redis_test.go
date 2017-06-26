package itest

import (
	"fmt"
	. "github.com/rjansen/boost/redis"
	"github.com/rjansen/migi"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	cacheClient *Client
	setted      = false
	expires     = 60
	testConfig  *Configuration
)

func init() {
	os.Args = append(os.Args, "-ecf", "etc/cache/cache.yaml")
	if err := migi.UnmarshalKey("cache.redis", &testConfig); err != nil {
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
	uniqueKey := fmt.Sprintf("add.%v", time.Now().UnixNano())
	err := cacheClient.Add(uniqueKey, expires, []byte("1234567890"))
	assert.Nil(t, err)
	err = cacheClient.Delete(uniqueKey)
	assert.Nil(t, err)
}

func TestIntegrationGetItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	uniqueKey := fmt.Sprintf("get.%v", time.Now().UnixNano())
	err := cacheClient.Add(uniqueKey, expires, []byte("1234567890"))
	assert.Nil(t, err)
	item, err := cacheClient.Get(uniqueKey)
	assert.Nil(t, err)
	assert.NotNil(t, item)
	err = cacheClient.Delete(uniqueKey)
	assert.Nil(t, err)
}

func TestIntegrationDelItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	uniqueKey := fmt.Sprintf("del.%v", time.Now().UnixNano())
	err := cacheClient.Add(uniqueKey, expires, []byte("1234567890"))
	assert.Nil(t, err)
	err = cacheClient.Delete(uniqueKey)
	assert.Nil(t, err)
}

func TestIntegrationGetEmptyItem(t *testing.T) {
	assert.NotNil(t, cacheClient)
	if beforeErr := before(); beforeErr != nil {
		assert.Fail(t, beforeErr.Error())
	}
	item, err := cacheClient.Get("invalidkey1")
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
