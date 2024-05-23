package schema

import "time"

type UserPoints struct {
	AccountAddress string    `gorm:"primaryKey;type:varchar(42);not null"`
	Points         int64     `gorm:"not null"`
	UpdateAt       time.Time `gorm:"default:current_timestamp"`
}

func (UserPoints) TableName() string {
	return "sst.sst_user_points"
}
