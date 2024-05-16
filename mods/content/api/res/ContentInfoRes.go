package res

import (
	"github.com/google/uuid"
	"time"
)

type ContentInfoRes struct {
	ContentID      uuid.UUID `json:"contentID"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	AccountAddress string    `json:"accountAddress"`
	Location       string    `json:"location"`
	IsPublic       bool      `json:"isPublic"`
	CreatedAt      time.Time `json:"created_at"`
}
