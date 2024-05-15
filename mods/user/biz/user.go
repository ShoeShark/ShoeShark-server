package biz

import (
	"github.com/shoe-shark/shoe-shark-service/mods/user/dao"
	"github.com/shoe-shark/shoe-shark-service/mods/user/schema"
	"github.com/shoe-shark/shoe-shark-service/repository"
)

func GetUserByAccountAddress(accountAddress string) (*schema.User, error) {

	rp := repository.GetPGRepository()
	user, err := dao.GetByAccountAddress(rp, accountAddress)
	if err != nil {
		return nil, err
	}

	return user, nil
}
