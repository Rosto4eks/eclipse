package database

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	file, err := os.ReadFile("./config/db.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
