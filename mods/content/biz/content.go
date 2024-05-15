package biz

import (
	"context"
	"github.com/google/uuid"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/content/schema"
	"github.com/shoe-shark/shoe-shark-service/repository"
)

func CreateContent(ctx context.Context, req *req.CreateContentReq) error {
	content := schema.Content{
		//ContentID:      uuid.New(),
		UserId:         uuid.New(), // Assuming UID is generated here; adjust as needed
		Title:          req.Title,
		Body:           req.Body,
		AccountAddress: req.AccountAddress,
		Location:       req.Location,
		IsPublic:       req.IsPublic,
		//CreatedAt:      time.Now(),
		//UpdatedAt:      time.Now(),
	}

	rp := repository.GetPGRepository()

	if err := rp.Create(&content).Error; err != nil {
		return err
	}
	return nil
}

func DeleteContent(ctx context.Context, contentID string) error {
	rp := repository.GetPGRepository()

	var content schema.Content
	if err := rp.Where("content_id = ?", contentID).First(&content).Error; err != nil {
		return err
	}
	content.IsDelete = true
	if err := rp.Save(&content).Error; err != nil {
		return err
	}
	return nil
}

func UpdateContent(ctx context.Context, req *req.UpdateContentReq) error {
	var updateContent = schema.Content{
		Title:          req.Title,
		Body:           req.Body,
		AccountAddress: req.AccountAddress,
		Location:       req.Location,
		IsPublic:       req.IsPublic,
	}

	rp := repository.GetPGRepository()
	var content schema.Content
	if err := rp.Where("content_id = ?", req.ContentID).First(&content).Error; err != nil {
		return err
	}

	if err := rp.SaveOne(&updateContent); err != nil {
		return err
	}
	return nil
}

func GetContent(ctx context.Context, contentID string) (*schema.Content, error) {
	rp := repository.GetPGRepository()
	var content schema.Content
	if err := rp.Where("content_id = ? AND is_delete = FALSE", contentID).First(&content).Error; err != nil {
		return nil, err
	}
	return &content, nil
}
