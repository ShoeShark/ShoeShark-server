package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/biz"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/schema"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

// Create
// @Summary Add a new pet to the store
// @Description add a new pet to the store
// @Tags 测试
// @Accept  json
// @Produce  json
// @Param   demo  body      schema.DemoForm   true  "DemoForm contains all necessary pet information"
// @Router /api/v1/demo/create [put]
func Create(c *gin.Context) {
	ctx := c.Request.Context()

	var item schema.DemoForm
	if err := c.BindJSON(&item); err != nil {
		util.ResError(c, "body serialize error", nil)
		return
	}

	if err := biz.SaveOne(ctx, &item); err != nil {
		util.ResError(c, "Save Error", nil)
		return
	}

	util.ResOk(c, "Save Success", nil)
}
