package schema

import "github.com/google/uuid"

type ContentTag struct {
	ContentID uuid.UUID `gorm:"type:uuid;not null" json:"content_id"`
	TagID     int       `gorm:"not null" json:"tag_id"`
}

func (a *ContentTag) TableName() string {
	return "sst.sst_content_tags"
}
