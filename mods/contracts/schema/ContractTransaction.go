package schema

import (
	"time"
)

type ContractTransaction struct {
	ID             uint      `gorm:"primaryKey"`
	AccountAddress string    `gorm:"type:varchar(42);not null"`
	TxnHash        string    `gorm:"type:varchar(66);not null"`
	Method         string    `gorm:"type:varchar(100)"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (ContractTransaction) TableName() string {
	return "sst.sst_contract_transactions"
}
