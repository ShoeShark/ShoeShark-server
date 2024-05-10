// @title shoe-shark-service API
// @description shoe-shark-service API
// @version 1.0
// @BasePath /shoe-shark

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/config"
	"github.com/shoe-shark/shoe-shark-service/global"
	"github.com/shoe-shark/shoe-shark-service/logger"
	"github.com/shoe-shark/shoe-shark-service/repository"
	"github.com/shoe-shark/shoe-shark-service/routers"
	log "github.com/sirupsen/logrus"
)

func main() {

	conf := config.InitConfig()
	logger.InitLogger(conf)
	rep := repository.InitDb(conf)

	global.CONFIG = conf
	global.REP = rep

	if global.CONFIG.AppMode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if global.CONFIG.AppMode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := routers.InitRouter(conf)

	err := router.Run(fmt.Sprintf("%s%d", ":", global.CONFIG.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
