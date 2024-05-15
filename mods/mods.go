package mods

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/content"
	"github.com/shoe-shark/shoe-shark-service/mods/user"
)

func RegisterRouters(router *gin.RouterGroup) {
	r := router.Group("/api/v1")

	content.RegisterV1Routers(r)
	user.RegisterV1Routers(r)
}
