package config

type Config struct {
	AppMode string `yaml:"app_mode"`
	Server  struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Logging struct {
		Dir     string `yaml:"dir"`
		Name    string `yaml:"name"`
		Level   string `yaml:"level"`
		MaxSize int    `yaml:"max_size"`
		Enable  bool   `yaml:"enable"`
	} `yaml:"logger"`
	Database struct {
		Pg struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Dbname   string `yaml:"dbname"`
		} `yaml:"pg"`
	} `yaml:"database"`
}
