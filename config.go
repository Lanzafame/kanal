package kanal

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config contains the Toml tree read in from kanal.toml.
type Config struct {
	Sources  []Source  `toml:"sources"`
	Stores   []Store   `toml:"stores"`
	Mappings []Mapping `toml:"mappings"`
}

// LoadConfig loads the toml config file into Config.
func LoadConfig(path string) *Config {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Printf("load config: decode file: %v", err)
	}
	return &config
}
