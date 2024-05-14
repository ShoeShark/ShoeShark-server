package schema

import (
	"github.com/google/uuid"
	"time"
)

type Content struct {
	ContentID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"content_id"`
	UserId         uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title          string    `gorm:"type:varchar(255);not null" json:"title"`
	Body           string    `gorm:"type:text;not null" json:"body"`
	AccountAddress string    `gorm:"type:varchar(42);not null" json:"account_address"`
	Location       string    `gorm:"type:varchar(255)" json:"location"`
	IsPublic       bool      `gorm:"default:true" json:"is_public"`
	IsDelete       bool      `gorm:"default:false" json:"is_delete"`
	CreatedAt      time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}

func (a *Content) TableName() string {
	return "sst.sst_contents"
}
