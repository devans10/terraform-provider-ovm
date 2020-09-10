package ovm

import (
	"log"
	"os"

	"github.com/devans10/go-ovm-helper/ovmHelper"
)

// Config - structure for connecting to OVM Manager
type Config struct {
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Entrypoint string `mapstructure:"entrypoint"`
}

// Client - returns a new client for accessing pingdom.
func (c *Config) Client() (*ovmHelper.Client, error) {

	if v := os.Getenv("OVM_USERNAME"); v != "" {
		c.User = v
	}
	if v := os.Getenv("OVM_PASSWORD"); v != "" {
		c.Password = v
	}
	if v := os.Getenv("OVM_ENTRYPOINT"); v != "" {
		c.Entrypoint = v
	}

	client := ovmHelper.NewClient(c.User, c.Password, c.Entrypoint)

	log.Printf("[INFO] OVM Client configured for user: %s", c.User)

	return client, nil
}
