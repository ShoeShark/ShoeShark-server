package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/biz"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

// MintWhiteList godoc
// @Summary mint nft
// @Description mint nft , this mint must in white list
// @Tags contract_nft
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} res.ContentInfoRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/contract/nft/mint/white [get]
func MintWhiteList(c *gin.Context) {
	newCtx := middleware.GenContextWithInformation(c)

	err := biz.MintWhiteList(&newCtx)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}
	util.ResOk(c)
}
