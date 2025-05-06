package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Api Api `yaml:"api"`
}

type Api struct {
	Host     string `yaml:"host"`
	Endpoint string `yaml:"endpoint"`
	Token    string `yaml:"token"`
	Filter   Filter `yaml:"filter"`
}

type Filter struct {
	Kind       string   `yaml:"kind"`
	Currencies []string `yaml:"currencies"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	config.Api.Token = os.ExpandEnv(config.Api.Token)

	return &config, nil
}
