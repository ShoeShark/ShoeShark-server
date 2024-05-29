package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

var cfg Config

type Config struct {
	AppMode     string    `yaml:"app_mode"`
	Server      *Server   `yaml:"server"`
	Logging     *Logging  `yaml:"logging"`
	Database    *Database `yaml:"database"`
	ResourceDir string    `yaml:"resource_dir"`
	Eth         ETH       `yaml:"eth"`
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

type ETH struct {
	PrivateKey string     `yaml:"private_key"`
	Fuji       FujiConfig `yaml:"fuji"`
}

type FujiConfig struct {
	Url string `yaml:"url"`
}

func GetConfig() *Config {
	return &cfg
}

func InitConfig(resourceDirectory string) {
	ymlData, err := os.ReadFile(resourceDirectory + "/application.yml")
	if err != nil {
		panic("parse config fail" + err.Error())
	}

	if err := yaml.Unmarshal(ymlData, &cfg); err != nil {
		panic("parse config fail" + err.Error())
	}

	if cfg.AppMode == "dev" {
		err := godotenv.Load("config/local.env")
		if err != nil {
			log.Error("Error loading .env file:\t", err.Error())
		}
	}

	appMode := os.Getenv("APP_MODE")
	cfg.AppMode = appMode

	pgHost := os.Getenv("PG_HOST")
	pgUsername := os.Getenv("PG_USERNAME")
	pgPassword := os.Getenv("PG_PASSWORD")

	cfg.Database.Pg.Host = pgHost
	cfg.Database.Pg.Username = pgUsername
	cfg.Database.Pg.Password = pgPassword

	privateKey := os.Getenv("PRIVATE_KEY")
	fujiUrl := os.Getenv("FUJI_URL")
	cfg.Eth.PrivateKey = privateKey
	cfg.Eth.Fuji.Url = fujiUrl

	log.Info("load config: ", cfg)
	log.Info("load config success")
}
