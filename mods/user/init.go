package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/user/api"
)

func RegisterV1Routers(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.GET("/nonce/:accountAddress", api.GetNonce)
		authRouter.POST("/nonce/verify", api.VerifySignature)
	}
	userRouter := router.Group("/user")
	userRouter.Use(middleware.AuthMiddleware())

	{
		userRouter.GET("/", api.GetUserByAccountAddress)
	}
}
