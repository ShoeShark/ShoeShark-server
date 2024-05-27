package res

import (
	"github.com/google/uuid"
	"time"
)

type CreateContentCommentRes struct {
	Description    string    `json:"description"`
	AccountAddress string    `json:"accountAddress"`
	ContentId      uuid.UUID `json:"contentId"`
	CreatedAt      time.Time `json:"createdAt"`
}
