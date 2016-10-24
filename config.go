package cache

import (
	"fmt"
)

//Configuration holds cache connections parameters
type Configuration struct {
	Provider string `mapstructure:"provider"`
	URL      string `mapstructure:"url"`
}

func (c Configuration) String() string {
	return fmt.Sprintf("cache.Configuration Provider=%s URL=%v", c.Provider, c.URL)
}
