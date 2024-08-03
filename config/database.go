// config.go
package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Type    string `yaml:"type"`
		MongoDB struct {
			URI      string `yaml:"uri"`
			Database string `yaml:"database"`
		} `yaml:"mongodb"`
		SQLite struct {
			Path string `yaml:"path"`
		} `yaml:"sqlite"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	var config Config
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
