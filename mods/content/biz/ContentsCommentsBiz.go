package biz

import (
	"context"
	"github.com/google/uuid"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/req"
	"github.com/shoe-shark/shoe-shark-service/mods/content/api/res"
	"github.com/shoe-shark/shoe-shark-service/mods/content/dao"
	"github.com/shoe-shark/shoe-shark-service/mods/content/schema"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func InsertComment(ctx *context.Context, req *req.CreateContentCommentReq) error {
	accountAddress := (*ctx).Value("accountAddress").(string)
	contentID, err := uuid.Parse(req.ContentId)
	if err != nil {
		return err
	}
	// 插入评论
	comment := schema.ContentsComment{
		Description:    req.Description,
		AccountAddress: accountAddress,
		ContentID:      contentID,
	}
	if err := dao.InsertComment(db.GetPGRepository(), &comment); err != nil {
		return err
	}

	return nil
}

func GetCommentsByContentID(contentID uuid.UUID) ([]res.CreateContentCommentRes, error) {
	// 查询评论
	comments, err := dao.GetCommentsByContentID(db.GetPGRepository(), contentID)
	if err != nil {
		return nil, err
	}

	var contentCommentInfos []res.CreateContentCommentRes
	if len(comments) > 0 {

		for _, comment := range comments {
			contentCommentInfos = append(contentCommentInfos, res.CreateContentCommentRes{
				ContentId:      comment.ContentID,
				Description:    comment.Description,
				CreatedAt:      util.CustomTime{Time: comment.CreatedAt},
				AccountAddress: comment.AccountAddress,
			})
		}
		return contentCommentInfos, nil
	}

	return contentCommentInfos, nil
}
