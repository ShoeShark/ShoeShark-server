package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/api"
)

func RegisterV1Routers(router *gin.RouterGroup) {
	r := router.Group("/demo")
	r.PUT("/create", api.Create)
}
