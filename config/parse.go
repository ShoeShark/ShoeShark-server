package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func InitConfig() *Config {

	var config Config

	data, err := os.ReadFile("resources/application.dev.yml")
	if err != nil {
		panic("parse config fail" + err.Error())
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		panic("parse config fail" + err.Error())
	}

	return &config
}
