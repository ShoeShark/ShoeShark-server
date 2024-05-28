package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shoe-shark/shoe-shark-service/docs"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
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

func InitRouter() *gin.Engine {
	r := gin.New()

	if os.Getenv("GIN_MODE") == gin.DebugMode {
		r.Use(CORSMiddleware())
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(middleware.GinLoggerPrintConf())

	ssr := r.Group("/shoe-shark")
	mods.RegisterRouters(ssr)

	return r
}
