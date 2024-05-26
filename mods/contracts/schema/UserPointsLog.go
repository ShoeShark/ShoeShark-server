package schema

import "time"

type UserPointsLog struct {
	ID             uint      `gorm:"primaryKey"`
	AccountAddress string    `gorm:"not null;size:42"`
	Points         int64     `gorm:"not null"`
	IsSyncLink     bool      `gorm:"default:false"`
	CreatedAt      time.Time `gorm:"default:current_timestamp"`
}

func (UserPointsLog) TableName() string {
	return "sst.sst_user_points_log"
}
