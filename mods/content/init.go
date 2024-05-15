package content

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api"
)

func RegisterV1Routers(router *gin.RouterGroup) {
	r := router.Group("/content")
	r.Use(middleware.AuthMiddleware())

	r.POST("/save", api.CreateContent)
	r.DELETE("/:contentId", api.DeleteContent)
	r.GET("/:contentId", api.GetContent)
	r.PUT("/edit", api.UpdateContent)
}
