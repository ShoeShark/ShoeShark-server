package schema

import "time"

type MerkleTreeLog struct {
	ID        uint      `gorm:"primaryKey"`
	Accounts  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
}

func (MerkleTreeLog) TableName() string {
	return "sst.sst_merkle_tree_log"
}
