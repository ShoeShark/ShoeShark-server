package schema

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserId         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"user_id"`
	Username       string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email          string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	AccountAddress string    `gorm:"type:varchar(42);not null" json:"account_address"`
	IsDelete       bool      `gorm:"default:false" json:"is_delete"`
	CreatedAt      time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
