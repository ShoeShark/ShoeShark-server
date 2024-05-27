package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/user/biz"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

// GetUserByAccountAddress godoc
// @Summary Get User
// @Description Get User  By AccountAddress
// @Tags user
// @Accept json
// @Security ApiKeyAuth
// @Success 200 {object} res.UserInfoRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/user [get]
func GetUserByAccountAddress(c *gin.Context) {
	newCtx := middleware.GenContextWithInformation(c)

	user, err := biz.GetUserByAccountAddress(&newCtx)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	if user == nil {
		util.ResErrorWithMsg(c, "user not found")
	}

	util.ResOkWithData(c, user)
}
