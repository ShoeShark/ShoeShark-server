package dao

import (
	"github.com/shoe-shark/shoe-shark-service/mods/points/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/points/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	"gorm.io/gorm"
	"time"
)

func GetUnSyncedLogs(rp *db.Repository) ([]schema.UserPointsLog, error) {
	var logs []schema.UserPointsLog
	result := rp.Where("is_sync_link = ?", false).Find(&logs)
	return logs, result.Error
}

func AddPointsLog(rp *db.Repository, log *schema.UserPointsLog) error {
	result := rp.Create(&log)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func VerifyIsSignIn(rp *db.Repository, account string) (bool, error) {
	var log schema.UserPointsLog
	today := time.Now().Format("2006-01-02")

	tx := rp.Where("account_address = ? AND source = ? AND DATE(created_at) = ?", account, constants.SIGN_IN, today).First(&log)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return true, tx.Error
	}

	return true, nil
}

func VerifyPublishContent(rp *db.Repository, account string) (bool, error) {
	// TODO
	return false, nil
}

func UpdateIsSyncLink(rp *db.Repository, ids []uint, newValue bool) error {
	result := rp.Model(&schema.UserPointsLog{}).Where("id IN ?", ids).Update("is_sync_link", newValue)
	return result.Error
}
