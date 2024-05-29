package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

// CustomTime is a wrapper around time.Time to handle custom JSON formatting
type CustomTime struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface for CustomTime
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format("2006-01-02 15:04:05"))
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomTime
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}
	parsedTime, err := time.Parse(`"2006-01-02 15:04:05"`, str)
	if err != nil {
		return err
	}
	ct.Time = parsedTime
	return nil
}
