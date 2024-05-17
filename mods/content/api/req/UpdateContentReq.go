package req

import "github.com/google/uuid"

type UpdateContentReq struct {
	ContentID   uuid.UUID `json:"contentId" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"`
	IsPublic    bool      `json:"isPublic"`
}
