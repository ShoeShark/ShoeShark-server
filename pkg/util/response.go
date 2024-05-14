package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResOk(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "ok",
		Data: nil,
	})
	return
}

func ResOkWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  msg,
		Data: nil,
	})
	return
}

func ResOkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "ok",
		Data: data,
	})
	return
}

func ResError(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  "fail",
		Data: nil,
	})
	return
}

func ResErrorWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  msg,
		Data: nil,
	})
	return
}
