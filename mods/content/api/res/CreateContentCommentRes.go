package res

import (
	"github.com/google/uuid"
	"github.com/shoe-shark/shoe-shark-service/pkg/util"
)

type CreateContentCommentRes struct {
	Description    string          `json:"description"`
	AccountAddress string          `json:"accountAddress"`
	ContentId      uuid.UUID       `json:"contentId"`
	CreatedAt      util.CustomTime `json:"createdAt"`
}
