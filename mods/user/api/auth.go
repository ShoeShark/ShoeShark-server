package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/user/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/user/api/res"
	"github.com/shoe-shark/shoe-shark-service/mods/user/biz"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

// Test godoc
// @Summary Get Nonce
// @Description Get Nonce By Account Address
// @Tags auth
// @Accept json
// @Produce json
// @Param accountAddress path string true "Wallet Address"
// @Success 200 {object} res.GetNonceRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/auth/test/{accountAddress} [get]
func Test(c *gin.Context) {
	accountAddress := c.Param("accountAddress")
	nonce, err := biz.GetNonce(accountAddress)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	util.ResOkWithData(c, &res.GetNonceRes{Nonce: nonce})
}

// GetNonce godoc
// @Summary Get Nonce
// @Description Get Nonce By Account Address
// @Tags auth
// @Accept json
// @Produce json
// @Param accountAddress path string true "Wallet Address"
// @Success 200 {object} res.GetNonceRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/auth/nonce/{accountAddress} [get]
func GetNonce(c *gin.Context) {
	accountAddress := c.Param("accountAddress")
	nonce, err := biz.GetNonce(accountAddress)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	util.ResOkWithData(c, &res.GetNonceRes{Nonce: nonce})
}

// VerifySignature godoc
// @Summary Verify Signature
// @Description Verify the provided signature and return a JWT if valid
// @Tags auth
// @Accept json
// @Produce json
// @Param content body req.VerifySignatureReq true "Content to verify"
// @Success 200 {object} res.VerifySignatureRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/auth/nonce/verify [post]
func VerifySignature(c *gin.Context) {
	var verifyReq = req.VerifySignatureReq{}
	if err := c.BindJSON(&verifyReq); err != nil {
		util.ResErrorWithMsg(c, "Invalid request")
		return
	}

	token, err := biz.VerifySignature(verifyReq.AccountAddress, verifyReq.Signature)
	if err != nil {
		util.ResErrorWithMsg(c, "invalid signature")
		return
	}

	util.ResOkWithData(c, &res.VerifySignatureRes{Token: token})
}
