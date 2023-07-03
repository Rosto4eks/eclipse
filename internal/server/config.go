package server

import (
	"os"

	"gopkg.in/yaml.v3"
)

// config for server
type Config struct {
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	LogLevel string `yaml:"loglevel"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfgfile, err := os.ReadFile("./config/server.yaml")
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(cfgfile, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
