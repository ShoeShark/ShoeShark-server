package schema

import (
	"time"
)

type User struct {
	AccountAddress string    `gorm:"type:varchar(42);not null;primary_key"`
	Username       string    `gorm:"type:varchar(255);unique;not null"`
	Email          string    `gorm:"type:varchar(255);unique;not null"`
	Nonce          string    `gorm:"type:varchar(32);not null"`
	IsDelete       bool      `gorm:"default:false"`
	CreatedAt      time.Time `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp"`
}

func (User) TableName() string {
	return "sst.sst_user"
}
