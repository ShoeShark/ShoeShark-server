package dao

import (
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func GetUnSyncedLogs(rp *db.Repository) ([]schema.UserPointsLog, error) {
	var logs []schema.UserPointsLog
	result := rp.Where("is_sync_link = ?", false).Find(&logs)
	return logs, result.Error
}

func UpdateIsSyncLink(rp *db.Repository, ids []uint, newValue bool) error {
	result := rp.Model(&schema.UserPointsLog{}).Where("id IN ?", ids).Update("is_sync_link", newValue)
	return result.Error
}
