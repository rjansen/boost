package cache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupErr(t *testing.T) {
	err := Setup(nil)
	assert.NotNil(t, err)
}

func TestSetupSuccess(t *testing.T) {
	err := Setup(
		&Configuration{
			Provider: "mock",
			URL:      "mockURL",
		},
	)
	assert.Nil(t, err)
}

func TestConfiguration(t *testing.T) {
	provider := "mockProvider"
	url := "mock://MockCacheURL"
	cfg := &Configuration{
		Provider: provider,
		URL:      url,
	}
	cfgStr := cfg.String()
	assert.Contains(t, cfgStr, provider)
	assert.Contains(t, cfgStr, url)
}

func TestSetGetClientOnContext(t *testing.T) {
	c := context.Background()
	c, err := SetClient(c, NewClientMock())
	assert.Nil(t, err)
	assert.NotZero(t, c)

	client, err := GetClient(c)
	assert.Nil(t, err)
	assert.NotZero(t, client)
}

func TestSetGetClientOnContextErr(t *testing.T) {
	c, err := SetClient(nil, nil)
	assert.NotNil(t, err)
	assert.Zero(t, c)

	client, err := GetClient(c)
	assert.NotNil(t, err)
	assert.Zero(t, client)

	c = context.Background()
	client, err = GetClient(c)
	assert.NotNil(t, err)
	assert.Zero(t, client)
}
