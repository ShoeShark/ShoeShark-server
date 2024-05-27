package mods

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/content"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts"
	"github.com/shoe-shark/shoe-shark-service/mods/points"
	"github.com/shoe-shark/shoe-shark-service/mods/upload"
	"github.com/shoe-shark/shoe-shark-service/mods/user"
)

func RegisterRouters(router *gin.RouterGroup) {
	router.Static("/static", "resources/static")

	r := router.Group("/api/v1")

	content.RegisterV1Routers(r)
	user.RegisterV1Routers(r)
	contracts.RegisterV1Routers(r)
	upload.RegisterV1Routers(r)
	points.RegisterV1Routers(r)
}
