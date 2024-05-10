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

// ResOk setting gin.JSON
func ResOk(c *gin.Context, mst string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "ok",
		Data: data,
	})
	return
}

// ResError setting gin.JSON
func ResError(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  msg,
		Data: data,
	})
	return
}
