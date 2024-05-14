package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var cfg Config

type Config struct {
	AppMode  string    `yaml:"app_mode"`
	Server   *Server   `yaml:"server"`
	Logging  *Logging  `yaml:"logging"`
	Database *Database `yaml:"database"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Logging struct {
	Dir     string `yaml:"dir"`
	Name    string `yaml:"name"`
	Level   string `yaml:"level"`
	MaxSize int    `yaml:"max_size"`
	Enable  bool   `yaml:"enable"`
}

type Database struct {
	Pg *DatabasePg `yaml:"pg"`
}

type DatabasePg struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func GetConfig() *Config {
	return &cfg
}

func InitConfig() {
	data, err := os.ReadFile("resources/application.dev.yml")
	if err != nil {
		panic("parse config fail" + err.Error())
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic("parse config fail" + err.Error())
	}
}
