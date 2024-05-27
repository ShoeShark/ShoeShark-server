package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/config"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
	"io"
	"os"
	"path/filepath"
)

// Upload
// 上传图片接口
// @Tags 上传
// @Summary 上传图片
// @Description 上传图片
// @Accept  multipart/form-data
// @Produce text/plain
// @Security ApiKeyAuth
// @Param file formData file true "要上传的图片文件"
// @Success 200 {string} string "文件上传成功的哈希值"
// @Failure 400 {string} string "获取表单错误"
// @Failure 500 {string} string "服务器错误"
// @Router /api/v1/oss/upload [post]
func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		util.ResErrorWithMsg(c, fmt.Sprintf("Get form err: %s", err.Error()))
		return
	}
	defer file.Close()

	// 计算文件的哈希值
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		util.ResErrorWithMsg(c, fmt.Sprintf("Hash file err: %s", err.Error()))
		return
	}
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	// 获取文件后缀
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".file" // 如果没有后缀，使用默认后缀
	}

	// 检查文件是否已经存在
	filepath := config.GetConfig().ResourceDir + "/" + hashString + ext
	if _, err := os.Stat(filepath); err == nil {
		util.ResErrorWithMsg(c, fmt.Sprintf("File %s already exists", header.Filename))
		return
	}

	// 重新读取文件
	file.Seek(0, io.SeekStart)
	out, err := os.Create(filepath)
	if err != nil {
		util.ResErrorWithMsg(c, fmt.Sprintf("Create file err: %s", err.Error()))
		return
	}
	defer out.Close()

	// 将上传的文件复制到目标文件
	if _, err := io.Copy(out, file); err != nil {
		util.ResErrorWithMsg(c, fmt.Sprintf("Save file err: %s", err.Error()))
		return
	}

	//c.String(http.StatusOK, fmt.Sprintf("File uploaded successfully with hash %s", hashString))
	util.ResOkWithData(c, "/shoe-shark/static/"+hashString+ext)
}

//func Image(c *gin.Context) {
//	hash := c.Param("hash")
//	filepath := "./resources/static/" + hash
//
//	// 检查文件是否存在
//	if _, err := os.Stat(filepath); os.IsNotExist(err) {
//		c.String(http.StatusNotFound, "File not found")
//		return
//	}
//
//	// 发送文件
//	c.File(filepath)
//}
