package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
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

func GenericConvert(src, dest interface{}) error {
	srcVal := reflect.ValueOf(src)
	destVal := reflect.ValueOf(dest).Elem()

	if srcVal.Kind() != reflect.Slice || destVal.Kind() != reflect.Slice {
		return errors.New("src and dest should be slices")
	}

	for i := 0; i < srcVal.Len(); i++ {
		srcElem := srcVal.Index(i)
		destElem := reflect.New(destVal.Type().Elem()).Elem()

		for j := 0; j < srcElem.NumField(); j++ {
			srcField := srcElem.Type().Field(j)
			srcFieldName := srcField.Name

			if destField := destElem.FieldByName(srcFieldName); destField.IsValid() && destField.CanSet() {
				destField.Set(srcElem.Field(j))
			}
		}

		destVal.Set(reflect.Append(destVal, destElem))
	}

	return nil
}
