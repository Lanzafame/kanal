package kanal

import (
	"log"

	"github.com/pelletier/go-toml"
)

// Config contains the Toml tree read in from kanal.toml.
type Config struct {
	toml.TomlTree
	FilePath string
}

// NewConfig creates a new Config object.
func NewConfig() *Config {
	c := new(Config)
	return c
}

// loadConfig loads the toml config file into Config.
func (c *Config) LoadConfig(args ...interface{}) error {
	if len(args) > 0 {
		c.FilePath = args[0].(string)
	} else {
		c.FilePath = "kanal.toml"
	}
	config, err := toml.LoadFile(c.FilePath)
	if err != nil {
		log.Fatalf("loadConfig: %s", err)
		return err
	}
	c.TomlTree = *config
	return nil
}
