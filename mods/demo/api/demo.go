package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/demo/biz"
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

	var item req.DemoCreateReq
	if err := c.BindJSON(&item); err != nil {
		util.ResErrorWithMsg(c, "body serialize error")
		return
	}

	if err := biz.CreateDemo(ctx, &item); err != nil {
		util.ResErrorWithMsg(c, "Save Error")
		return
	}

	util.ResOkWithMsg(c, "Save Success")
}
