package dao

import (
	"github.com/google/uuid"
	"github.com/shoe-shark/shoe-shark-service/mods/content/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func InsertComment(rp *db.Repository, comment *schema.ContentsComment) error {
	return rp.Create(comment).Error
}

func GetCommentsByContentID(rp *db.Repository, contentID uuid.UUID) ([]schema.ContentsComment, error) {
	var comments []schema.ContentsComment
	if err := rp.Where("content_id = ?", contentID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
