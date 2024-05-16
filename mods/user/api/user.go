package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/user/biz"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
	"net/http"
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
	accountAddress, exists := c.Get("accountAddress")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}

	user, err := biz.GetUserByAccountAddress(accountAddress.(string))
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	if user == nil {
		util.ResErrorWithMsg(c, "user not found")
	}

	util.ResOkWithData(c, user)
}
