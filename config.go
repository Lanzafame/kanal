package kanal

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Sources  map[string]SourceConfig  `toml:"sources"`
	Stores   map[string]StoreConfig   `toml:"stores"`
	Mappings map[string]MappingConfig `toml:"mappings"`
}

type SourceConfig struct {
	Name string
	URL  string
}

type StoreConfig struct {
	Name string
	URL  string
}

type MappingConfig struct {
	Stores         []string
	Transformation string
}

// LoadConfig loads the toml config file into Config.
func LoadConfig(path string) *Config {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Printf("load config: decode file: %v", err)
	}

	return &config
}
