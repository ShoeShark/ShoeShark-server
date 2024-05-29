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
	"os"
	"path/filepath"
)

func main() {
	log.Info("Init ShoeSharkServer Start....")
	config.InitConfig("resources/application.dev.yml")

	cfg := config.GetConfig()
	if cfg.AppMode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.InitLogger()
	repository.Init()
	eth.InitClient(cfg)

	//获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// 构建资源目录的绝对路径
	resourcesDir := filepath.Join(cwd, "/resources/static")
	cfg.ResourceDir = resourcesDir

	router := routers.InitRouter()

	err = router.Run(fmt.Sprintf("%s%d", ":", cfg.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
