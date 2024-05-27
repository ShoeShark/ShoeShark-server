package dao

import (
	"fmt"
	"github.com/shoe-shark/shoe-shark-service/mods/points/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	"gorm.io/gorm"
	"math/big"
	"time"
)

// AddPoints 增加或更新账户的积分
func AddPoints(rp *db.Repository, account string, points int64) error {
	var userPoints schema.UserPoints
	result := rp.First(&userPoints, "account_address = ?", account)
	if result.Error == gorm.ErrRecordNotFound {
		// 记录不存在，创建新记录
		newRecord := schema.UserPoints{
			AccountAddress: account,
			Points:         points,
			CreatedAt:      time.Now(),
		}
		return rp.Create(&newRecord).Error
	} else if result.Error != nil {
		// 数据库错误
		return result.Error
	}

	// 记录存在，更新积分
	userPoints.Points += points
	return rp.Save(&userPoints).Error
}

func AddPointsBatch(rp *db.Repository, accountsHex []string, points []*big.Int) error {
	if len(accountsHex) != len(points) {
		return fmt.Errorf("accounts and points arrays must be of the same length")
	}

	for i, account := range accountsHex {
		if points[i].IsInt64() {
			err := AddPoints(rp, account, points[i].Int64())
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("point value too large for int64 for account: %s", account)
		}
	}
	return nil
}
