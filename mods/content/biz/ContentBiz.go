package biz

import (
	"context"
	"errors"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/res"
	"github.com/shoe-shark/shoe-shark-service/mods/content/schema"
	pointsBiz "github.com/shoe-shark/shoe-shark-service/mods/points/biz"
	"github.com/shoe-shark/shoe-shark-service/mods/points/constants"
	"github.com/shoe-shark/shoe-shark-service/repository"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	log "github.com/sirupsen/logrus"
)

func CreateContent(ctx *context.Context, req *req.CreateContentReq) error {
	accountAddress := (*ctx).Value("accountAddress").(string)
	content := schema.Content{
		Title:          req.Title,
		Description:    req.Description,
		Location:       req.Location,
		IsPublic:       req.IsPublic,
		AccountAddress: accountAddress,
	}

	rp := repository.GetPGRepository()

	if err := rp.Create(&content).Error; err != nil {
		return err
	}

	err := pointsBiz.AddPoints(accountAddress, 20, constants.PUBLISH_CONTENT)
	if err != nil {
		log.Error("[创建文章][记录积分失败] accounts: ", accountAddress, " ContentTitle: ", content.Title)
	}

	return nil
}

func DeleteContent(ctx *context.Context, contentID string) error {
	accountAddress := (*ctx).Value("accountAddress").(string)

	rp := repository.GetPGRepository()

	var content schema.Content
	if err := rp.Where("content_id = ?", contentID).First(&content).Error; err != nil {
		return err
	}

	if content.AccountAddress != accountAddress {
		return errors.New("not authorized to delete this content")
	}

	content.IsDelete = true
	if err := rp.Save(&content).Error; err != nil {
		return err
	}
	return nil
}

func UpdateContent(ctx *context.Context, req *req.UpdateContentReq) error {
	accountAddress := (*ctx).Value("accountAddress").(string)

	rp := repository.GetPGRepository()
	var content schema.Content
	if err := rp.Where("content_id = ?", req.ContentID).First(&content).Error; err != nil {
		return err
	}

	if content.AccountAddress != accountAddress {
		return errors.New("not authorized to update this content")
	}

	var updateContent = schema.Content{
		Title:          req.Title,
		Description:    req.Description,
		AccountAddress: accountAddress,
		Location:       req.Location,
		IsPublic:       req.IsPublic,
		ContentID:      content.ContentID,
	}

	if err := rp.SaveOne(&updateContent); err != nil {
		return err
	}
	return nil
}

func GetContent(ctx *context.Context, contentID string) (*res.ContentInfoRes, error) {
	rp := repository.GetPGRepository()
	var content schema.Content
	if err := rp.Where("content_id = ? AND is_delete = FALSE", contentID).First(&content).Error; err != nil {
		return nil, err
	}

	result := res.ConvertToContentInfoRes(&content)
	return &result, nil
}

func ListContent(queryReq *req.QueryContentReq) (*db.Page, error) {
	// 构建查询条件
	rp := repository.GetPGRepository()
	dbQuery := rp.Model(&schema.Content{}).Where("is_delete = FALSE")
	if queryReq.Title != "" {
		dbQuery = dbQuery.Where("title LIKE ?", "%"+queryReq.Title+"%")
	}
	if queryReq.Description != "" {
		dbQuery = dbQuery.Where("description LIKE ?", "%"+queryReq.Description+"%")
	}
	if queryReq.AccountAddress != "" {
		dbQuery = dbQuery.Where("account_address = ?", queryReq.AccountAddress)
	}
	if queryReq.Location != "" {
		dbQuery = dbQuery.Where("location = ?", queryReq.Location)
	}

	var total int64
	var records []schema.Content
	var page = &db.Page{
		Page:  queryReq.Page,
		Size:  queryReq.Size,
		Total: total,
	}

	dbQuery.Count(&total)
	if total < 1 {
		return page, nil
	}

	offset := (queryReq.Page - 1) * queryReq.Size
	err := dbQuery.Order("created_at DESC").Offset(offset).Limit(queryReq.Size).Find(&records).Error
	if err != nil {
		return nil, err
	}

	var contentInfos []res.ContentInfoRes
	for _, record := range records {
		contentInfos = append(contentInfos, res.ConvertToContentInfoRes(&record))
	}
	page.Records = &contentInfos

	return page, nil
}
