package mods

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/content"
)

func RegisterRouters(router *gin.RouterGroup) {
	r := router.Group("/api/v1")

	content.RegisterV1Routers(r)
}
