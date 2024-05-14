package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/content/biz"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

// CreateContent godoc
// @Summary Create a new content
// @Description Create a new content
// @Tags contents
// @Accept json
// @Produce json
// @Param content body req.CreateContentReq true "Content to create"
// @Success 200 {object} util.Response{Msg=string}
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/content/save [post]
func CreateContent(c *gin.Context) {
	var contentReq req.CreateContentReq

	if err := c.ShouldBindJSON(&contentReq); err != nil {
		util.ResErrorWithMsg(c, "Serialize Body Error")
		return
	}

	err := biz.CreateContent(c.Request.Context(), &contentReq)
	if err != nil {
		util.ResErrorWithMsg(c, "Save Error")
		return
	}

	util.ResOk(c)
}

// DeleteContent godoc
// @Summary Delete a content
// @Description Delete a content by ContentID
// @Tags contents
// @Produce json
// @Param id path string true "schema.Content ID"
// @Success 200 {object} util.Response{Msg=string}
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/content/{contentId} [delete]
func DeleteContent(c *gin.Context) {
	contentID := c.Param("contentId")

	err := biz.DeleteContent(c.Request.Context(), contentID)
	if err != nil {
		util.ResErrorWithMsg(c, "Delete Content Error")
		return
	}

	util.ResOk(c)
}

// UpdateContent godoc
// @Summary Update a content
// @Description Update a content by ContentID
// @Tags contents
// @Accept json
// @Produce json
// @Param content body req.UpdateContentReq true "Content to update"
// @Success 200 {object} util.Response{Msg=string}
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/content/edit [put]
func UpdateContent(c *gin.Context) {
	var contentReq req.UpdateContentReq

	if err := c.ShouldBindJSON(&contentReq); err != nil {
		util.ResErrorWithMsg(c, "Serialize Body Error")
		return
	}

	err := biz.UpdateContent(c.Request.Context(), &contentReq)
	if err != nil {
		util.ResErrorWithMsg(c, "Update Content Error")
		return
	}
	util.ResOk(c)
}

// GetContent godoc
// @Summary Get a content
// @Description Get a content by ID
// @Tags contents
// @Produce json
// @Param id path string true "schema.Content ContentID"
// @Success 200 {object} res.ContentInfoRes
// @Failure 500 {object} util.Response{Msg=string}
// @Router /api/v1/content/{contentId} [get]
func GetContent(c *gin.Context) {
	contentID := c.Param("contentId")
	content, err := biz.GetContent(c.Request.Context(), contentID)
	if err != nil {
		util.ResErrorWithMsg(c, "Content not found")
		return
	}
	util.ResOkWithData(c, content)
}
