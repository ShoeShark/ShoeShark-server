package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/user/biz"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

// GetUserByAccountAddress godoc
// @Summary Get User
// @Description Get User  By AccountAddress
// @Tags user
// @Accept json
// @Success 200 {object} util.Response{Msg=string}
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/user [get]
func GetUserByAccountAddress(c *gin.Context) {
	address := c.Param("accountAddress")

	user, err := biz.GetUserByAccountAddress(address)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	util.ResOkWithData(c, user)
}
