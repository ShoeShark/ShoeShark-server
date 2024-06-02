package points

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/points/api"
)

func RegisterV1Routers(router *gin.RouterGroup) {

	authRouter := router.Group("/points")
	authRouter.Use(middleware.AuthMiddleware())
	{
		authRouter.GET("/log", api.PointsLogsHandler)
		authRouter.GET("/signIn", api.SignInHandler)
		authRouter.GET("/account", api.GetPointsHandler)
		authRouter.GET("/add/publish/content", api.AddPublishContentPointsHandler)
	}
}
