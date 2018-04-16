package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents server and database
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("./main/config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
