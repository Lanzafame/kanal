package kanal

import (
	"log"

	"github.com/pelletier/go-toml"
)

type Config struct {
	toml.TomlTree
}

func loadConfig(file string) *Config {
	config, err := toml.LoadFile("kanal.toml")
	if err != nil {
		log.Fatalf("loadConfig: %s", err)
	}
	return &Config{*config}
}
