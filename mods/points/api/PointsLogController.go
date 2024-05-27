package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/points/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/points/biz"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
	"strconv"
)

// SignInHandler
// @Summary 签到
// @Description 签到得积分
// @Tags points
// @Produce json
// @Security ApiKeyAuth
// @Success 200 "ok"
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/points/signIn [get]
func SignInHandler(c *gin.Context) {
	newCtx := middleware.GenContextWithInformation(c)

	err := biz.SignIn(&newCtx)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}
	util.ResOk(c)
}

// PointsLogsHandler
// @Summary 查看积分记录
// @Description 查看个人积分记录
// @Tags points
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} []res.PointsLogRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/points/log [get]
func PointsLogsHandler(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 10
	}

	var queryReq = req.QueryPointsLogReq{
		Page: page,
		Size: size,
	}

	newCtx := middleware.GenContextWithInformation(c)
	logs, err := biz.GetPointsLogs(&newCtx, &queryReq)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	util.ResOkWithData(c, logs)
}