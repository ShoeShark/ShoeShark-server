package dao

import (
	"encoding/json"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func InsertAccounts(rp *db.Repository, accounts []string) error {
	// 将账户数组转换为 JSON 字符串
	accountsJSON, err := json.Marshal(accounts)
	if err != nil {
		return err
	}

	// 创建 MerkleTreeLog 记录并插入数据库
	logEntry := schema.MerkleTreeLog{
		Accounts: string(accountsJSON),
	}
	if result := rp.Create(&logEntry); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetLatestAccounts 查询最近一次插入的记录并返回转换后的账户信息
func GetLatestAccounts(rp *db.Repository) ([]string, error) {
	var logEntry schema.MerkleTreeLog
	// 查询最近一次插入的记录，按创建时间降序排序并获取第一条记录
	if result := rp.Order("created_at DESC").First(&logEntry); result.Error != nil {
		return nil, result.Error
	}

	// 将 JSON 字符串转换回 []string
	var accounts []string
	if err := json.Unmarshal([]byte(logEntry.Accounts), &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}
