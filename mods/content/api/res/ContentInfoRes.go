package res

import (
	"github.com/google/uuid"
	"time"
)

type ContentInfoRes struct {
	ContentID      uuid.UUID `json:"contentID"`
	UID            uuid.UUID `json:"userID"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	AccountAddress string    `json:"accountAddress"`
	Location       string    `json:"location"`
	IsPublic       bool      `json:"isPublic"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
