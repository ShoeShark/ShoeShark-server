package mods

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/demo"
)

func RegisterRouters(router *gin.RouterGroup) {
	r := router.Group("/api/v1")

	demo.RegisterV1Routers(r)
}
