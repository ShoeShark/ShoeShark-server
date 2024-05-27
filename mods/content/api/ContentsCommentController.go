package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/content/biz"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
	"net/http"
)

// CreateContentComments
// @Summary 评论内容
// @Description 给内容添加评论
// @Tags contents
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param content body req.CreateContentCommentReq true "内容评论结构体"
// @Success 200 {object} util.Response{Msg=string}
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/content/comments [post]
func CreateContentComments(c *gin.Context) {
	var commentReq req.CreateContentCommentReq
	if err := c.ShouldBindJSON(&commentReq); err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	newCtx := middleware.GenContextWithInformation(c)
	if err := biz.InsertComment(&newCtx, &commentReq); err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	util.ResOk(c)
}

// ListContentsComments godoc
// @Summary 评论列表
// @Description 根据内容Id获取评论
// @Tags contents
// @Produce json
// @Security ApiKeyAuth
// @Param contentId path string true "Content ID"
// @Success 200 {object} res.CreateContentCommentRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/content/comments/{contentId} [get]
func ListContentsComments(c *gin.Context) {
	contentID, err := uuid.Parse(c.Param("contentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid contentId"})
		return
	}

	comments, err := biz.GetCommentsByContentID(contentID)
	if err != nil {
		util.ResErrorWithMsg(c, err.Error())
		return
	}

	util.ResOkWithData(c, comments)
}
