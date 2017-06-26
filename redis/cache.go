package memcached

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rjansen/boost"
	"github.com/rjansen/l"
	"time"
)

var (
	//Config holds all package parameters instance
	Config             *Configuration
	ErrSetupMustCalled = errors.New("SetupMustCalled: Message='You must call Setup with a redis.Configuration before get a Pool reference'")
)

//Configuration holds cache connections parameters
type Configuration struct {
	URL      string `json:"url" mapstructure:"url"`
	Password string
	DB       int
}

func (c Configuration) String() string {
	return fmt.Sprintf("redis.Configuration URL=%v", c.URL)
}

//Cache is the interface interact with the memcached
type Cache interface {
	//Ping validate the cache connection
	Ping() *redis.StatusCmd
	//Get reads the value associated with the provided key
	Get(string) *redis.StringCmd
	//SetNX inserts a new item in the cache, Add throws error if the provided key was already defined
	SetNX(string, interface{}, time.Duration) *redis.BoolCmd
	//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	//Del removes the item associated with the provided keys
	Del(...string) *redis.IntCmd
	//Close releases all resources and terminate this client
	Close() error
}

//Setup configures the memcached cache implementor package
func Setup(config *Configuration) error {
	Config = config
	//Config.Password = ""
	//Config.Addrs = []string{"", "", ""}
	//Config.DB = 0
	pool := &Pool{
		client: NewClient(),
	}
	if err := boost.Setup(pool); err != nil {
		return err
	}
	return nil
}

//Pool controls how new gocql.Session will create and maintained
type Pool struct {
	client *Client
}

func (c Pool) String() string {
	return fmt.Sprintf("RedisPool Configuration=%s ClientIsNil=%t",
		Config.String(),
		c.client == nil,
	)
}

//Get creates and returns a Client reference
func (c *Pool) Get() (boost.Client, error) {
	if c == nil || c.client == nil {
		return nil, ErrSetupMustCalled
	}
	// if c.client.Closed() {
	// 	return nil, fmt.Errorf("cassandra.SessionIsClosedErr")
	// }
	l.Debug("redis.Get",
		l.String("Pool", c.String()),
		l.Bool("ClientIsNil", c.client == nil),
	)
	return c.client, nil
}

//Close close the database pool
func (c *Pool) Close() error {
	if c == nil || c.client == nil {
		return ErrSetupMustCalled
	}
	l.Info("redis.Close",
		l.String("RedisPool", c.String()),
	)
	c.client.Close()
	return nil
}
