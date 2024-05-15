package dao

import (
	"errors"
	"github.com/shoe-shark/shoe-shark-service/mods/user/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	"gorm.io/gorm"
)

func GetNonceByAccountAddress(rp *db.Repository, accountAddress string) (string, error) {
	user := schema.User{}
	err := rp.Where("account_address = ?", accountAddress).Select([]string{"nonce"}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "nil", err
	}

	return user.Nonce, nil
}

func UpdateNonceByAccountAddress(rp *db.Repository, accountAddress string, nonce string) error {
	err := rp.Model(&schema.User{}).Where("account_address = ?", accountAddress).Update("nonce", nonce).Error
	if err != nil {
		return err
	}
	return nil
}

func GetByAccountAddress(rp *db.Repository, accountAddress string) (*schema.User, error) {
	user := schema.User{}
	err := rp.Where("account_address = ?", accountAddress).Select([]string{"nonce"}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
