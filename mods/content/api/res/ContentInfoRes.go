package res

import (
	"github.com/google/uuid"
	"github.com/shoe-shark/shoe-shark-service/mods/content/schema"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

type ContentInfoRes struct {
	ContentID      uuid.UUID       `json:"contentId"`
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	AccountAddress string          `json:"accountAddress"`
	Location       string          `json:"location"`
	IsPublic       bool            `json:"isPublic"`
	CreatedAt      util.CustomTime `json:"createdAt"`
}

func ConvertToContentInfoRes(sc *schema.Content) ContentInfoRes {
	return ContentInfoRes{
		ContentID:      sc.ContentID,
		Title:          sc.Title,
		Description:    sc.Description,
		AccountAddress: sc.AccountAddress,
		Location:       sc.Location,
		IsPublic:       sc.IsPublic,
		CreatedAt:      util.CustomTime{Time: sc.CreatedAt},
	}
}
