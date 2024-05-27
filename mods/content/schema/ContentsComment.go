package schema

import (
	"github.com/google/uuid"
	"time"
)

type ContentsComment struct {
	ID             uint      `gorm:"primaryKey"`
	Description    string    `gorm:"type:text;not null"`
	AccountAddress string    `gorm:"type:varchar(42);not null"`
	ContentID      uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (a *ContentsComment) TableName() string {
	return "sst.sst_contents_comment"
}
