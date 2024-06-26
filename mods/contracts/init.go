package contracts

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/api"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/job"
)

func RegisterV1Routers(router *gin.RouterGroup) {

	authRouter := router.Group("/contract/nft")
	authRouter.Use(middleware.AuthMiddleware())
	{
		authRouter.GET("/mint/white", api.MintWhiteList)
	}

	contractJob := job.NewJob(eth.GetClient(), eth.GetPrivateKey())

	contractJob.StartJob()
}
