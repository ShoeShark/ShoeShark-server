// @title shoe-shark-service API
// @description shoe-shark-service API
// @version 1.0
// @BasePath /shoe-shark
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/config"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/logger"
	"github.com/shoe-shark/shoe-shark-service/repository"
	"github.com/shoe-shark/shoe-shark-service/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.InitConfig("resources/application.dev.yml")

	cfg := config.GetConfig()
	if cfg.AppMode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if cfg.AppMode == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.InitLogger()
	repository.Init()
	eth.InitClient(cfg)

	router := routers.InitRouter()

	err := router.Run(fmt.Sprintf("%s%d", ":", cfg.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
