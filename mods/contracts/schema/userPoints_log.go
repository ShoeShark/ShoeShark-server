package schema

import "time"

type UserPointsLog struct {
	AccountAddress string    `gorm:"primaryKey;type:varchar(42);not null"`
	Points         int64     `gorm:"not null"`
	CreatedAt      time.Time `gorm:"default:current_timestamp"`
}

func (UserPointsLog) TableName() string {
	return "sst.sst_user_points_log"
}
