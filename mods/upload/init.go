package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/upload/api"
)

func RegisterV1Routers(router *gin.RouterGroup) {
	router.Use(middleware.AuthMiddleware())
	authRouter := router.Group("/oss")
	{
		authRouter.POST("/upload", api.Upload)
		//authRouter.GET("/image", api.Image)
	}
}
