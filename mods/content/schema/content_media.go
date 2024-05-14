package schema

import (
	"github.com/google/uuid"
	"time"
)

type ContentMedia struct {
	MediaID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"media_id"`
	ContentID uuid.UUID `gorm:"type:uuid;not null" json:"content_id"`
	MediaType string    `gorm:"type:varchar(10);not null" json:"media_type"`
	MediaURL  string    `gorm:"type:text;not null" json:"media_url"`
	SortOrder int       `gorm:"not null" json:"sort_order"`
	IsDelete  bool      `gorm:"default:false" json:"is_delete"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

func (a *ContentMedia) TableName() string {
	return "sst.sst_content_media"
}
