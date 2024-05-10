package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/config"
	_ "github.com/shoe-shark/shoe-shark-service/docs"
	"github.com/shoe-shark/shoe-shark-service/mods"
	"github.com/shoe-shark/shoe-shark-service/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitRouter(config *config.Config) *gin.Engine {

	r := gin.New()

	if config.AppMode == "dev" {
		r.Use(CORSMiddleware())
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(middleware.GinLoggerPrintConf())
	//r.Use(gin.Recovery())

	//a := r.Group("/shoe-shark")
	ssr := r.Group("/shoe-shark")
	mods.RegisterRouters(ssr)

	return r
}
