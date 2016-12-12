package memcached

import (
	"errors"
	"github.com/rjansen/boost"
	"github.com/rjansen/l"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	//Config holds all package parameters instance
	Config *Configuration
)

//Configuration holds cache connections parameters
type Configuration struct {
	URL string `json:"url" mapstructure:"url"`
}

func (c Configuration) String() string {
	return fmt.Sprintf("memcached.Configuration URL=%v", c.URL)
}

//Cache is the interface interact with the memcached
type Cache interface {
	//Get reads the value associated with the provided key
	Get(key string) (*memcache.Item, error)
	//Add inserts a new item in the cache, Add throws error if the provided key was already defined
	Add(item *memcache.Item) error
	//Set inserts a new item in the cache if the key is new or modifies the value associated with the provided key
	Set(item *memcache.Item) error
	//Delete removes the item associated with the provided key
	Delete(key string) error
}

//Setup configures the memcached cache implementor package
func Setup(config *Configuration) error {
	Config = config
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
	return fmt.Sprintf("MemcachedPool Configuration=%s ClientIsNil=%t",
		Config.String(),
		c.client == nil,
	)
}

//Get creates and returns a Client reference
func (c *Pool) Get() (boost.Client, error) {
	if c == nil || c.client == nil {
		return nil, errors.New("SetupMustCalled: Message='You must call Setup with a memcached.Configuration before get a Pool reference')")
	}
	// if c.client.Closed() {
	// 	return nil, fmt.Errorf("cassandra.SessionIsClosedErr")
	// }
	l.Debug("memcached.Get",
		l.String("Pool", c.String()),
		l.Bool("ClientIsNil", c.client == nil),
	)
	return c.client, nil
}

//Close close the database pool
func (c *Pool) Close() error {
	if c == nil || c.client == nil {
		return errors.New("SetupMustCalled: Message='You must call Setup with a CassandraBConfig before get a Cassandrapool reference')")
	}
	l.Info("CloseMemcacheClient",
		l.String("MemcachedPool", c.String()),
	)
	c.client.Close()
	return nil
}
