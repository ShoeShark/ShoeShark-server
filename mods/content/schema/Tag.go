package schema

import "time"

type Tag struct {
	TagID     int       `gorm:"primary_key;auto_increment" json:"tag_id"`
	TagText   string    `gorm:"type:varchar(255);unique;not null" json:"tag_text"`
	IsDelete  bool      `gorm:"default:false" json:"is_delete"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

func (a *Tag) TableName() string {
	return "sst.sst_tags"
}
