package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database Database `yaml:"database"`
	Web      Web      `yaml:"web"`

	// Modules
	Spotify Spotify `yaml:"spotify"`
}

func Load() (*Config, error) {
	var config Config

	configPath, err := FindConfigPath()
	if err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
