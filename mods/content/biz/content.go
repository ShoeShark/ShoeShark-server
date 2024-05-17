package biz

import (
	"context"
	"errors"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/res"
	"github.com/shoe-shark/shoe-shark-service/mods/content/schema"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
	"github.com/shoe-shark/shoe-shark-service/repository"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
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

func GetContent(ctx *context.Context, contentID string) (*schema.Content, error) {
	rp := repository.GetPGRepository()
	var content schema.Content
	if err := rp.Where("content_id = ? AND is_delete = FALSE", contentID).First(&content).Error; err != nil {
		return nil, err
	}
	return &content, nil
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

	var total int64
	var records []schema.Content
	var page = &db.Page{
		Page:    queryReq.Page,
		Size:    queryReq.Size,
		Records: &records,
		Total:   total,
	}

	dbQuery.Count(&total)
	if total < 1 {
		return page, nil
	}

	offset := (queryReq.Page - 1) * queryReq.Size
	err := dbQuery.Offset(offset).Limit(queryReq.Size).Find(&records).Error
	if err != nil {
		return nil, err
	}

	var contentInfos []res.ContentInfoRes
	err = util.GenericConvert(&records, &contentInfos)

	return page, nil
}
