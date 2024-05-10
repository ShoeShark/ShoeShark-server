package logging

import "github.com/shoe-shark/shoe-shark-service/config"

func InitLogConfig(config *config.Config) {

	err := InitLogger(config)
	if err != nil {
		panic("init logger err:" + err.Error())
	}
}
