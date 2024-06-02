package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	pointsContractBiz "github.com/shoe-shark/shoe-shark-service/mods/contracts/biz"
	"github.com/shoe-shark/shoe-shark-service/mods/points/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/points/api/res"
	"github.com/shoe-shark/shoe-shark-service/mods/points/biz"
	"github.com/shoe-shark/shoe-shark-service/mods/points/constants"
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

// GetPointsHandler
// @Summary 账户积分
// @Description 账户积分
// @Tags points
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} res.AccountPointsInfoRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/points/account [get]
func GetPointsHandler(c *gin.Context) {
	newCtx := middleware.GenContextWithInformation(c)

	points, err := biz.GetPoints(&newCtx)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	linkPoints, err := pointsContractBiz.GetPoints(&newCtx)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
	}

	util.ResOkWithData(c, &res.AccountPointsInfoRes{
		Points:     points,
		LinkPoints: linkPoints.Int64(),
	})
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

// AddPublishContentPointsHandler
// @Summary 发布文章添加积分
// @Description 发布文章添加积分
// @Tags points
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 "ok"
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/points/add/publish/content [get]
func AddPublishContentPointsHandler(c *gin.Context) {

	accountAddress := c.GetString("accountAddress")

	err := biz.AddPoints(accountAddress, 20, constants.PUBLISH_CONTENT)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
	}

	util.ResOk(c)
}
