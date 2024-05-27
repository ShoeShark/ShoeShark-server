package res

import "time"

type CreateContentCommentRes struct {
	Description    string    `json:"description"`
	AccountAddress string    `json:"accountAddress"`
	ContentId      string    `json:"contentId"`
	CreatedAt      time.Time `json:"createdAt"`
}
