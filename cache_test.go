package boost

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPoolErr(t *testing.T) {
	pool, err := GetPool()
	assert.NotNil(t, err)
	assert.Nil(t, pool)
}

func TestSetupErr(t *testing.T) {
	err := Setup(nil)
	assert.NotNil(t, err)
}

func TestSetupSuccess(t *testing.T) {
	mockClient := NewClientMock()
	mockClient.On("Close").Return(nil)
	mockPool := NewClientPoolMock()
	mockPool.On("Get").Return(mockClient, nil)
	mockPool.On("Close").Return(nil)
	err := Setup(mockPool)
	assert.Nil(t, err)
}

func TestGetPool(t *testing.T) {
	pool, err := GetPool()
	assert.Nil(t, err)
	assert.NotNil(t, pool)
}

func TestConfiguration(t *testing.T) {
	provider := "mockProvider"
	cfg := &Configuration{
		Provider: provider,
	}
	cfgStr := cfg.String()
	assert.Contains(t, cfgStr, provider)
}

func TestSetGetClientOnContext(t *testing.T) {
	c := context.Background()
	c, err := SetClient(c)
	assert.Nil(t, err)
	assert.NotZero(t, c)

	client, err := GetClient(c)
	assert.Nil(t, err)
	assert.NotZero(t, client)
}

func TestSetGetClientOnContextErr(t *testing.T) {
	c, err := SetClient(nil)
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

func TestExecuteContext(t *testing.T) {
	err := ExecuteContext(
		func(c context.Context) error {
			assert.NotNil(t, c)
			client, err := GetClient(c)
			assert.Nil(t, err)
			assert.NotNil(t, client)
			return nil
		},
	)
	assert.Nil(t, err)
}

func TestExecuteContextErr(t *testing.T) {
	err := ExecuteContext(
		func(c context.Context) error {
			assert.NotNil(t, c)
			client, err := GetClient(c)
			assert.Nil(t, err)
			assert.NotNil(t, client)
			return errors.New("MockExecuteContextErr")
		},
	)
	assert.NotNil(t, err)
}

func TestExecute(t *testing.T) {
	err := Execute(
		func(c Client) error {
			assert.NotNil(t, c)
			return nil
		},
	)
	assert.Nil(t, err)
}

func TestExecuteClientErr(t *testing.T) {
	err := Execute(
		func(c Client) error {
			assert.NotNil(t, c)
			return errors.New("MockExecuteErr")
		},
	)
	assert.NotNil(t, err)
}
