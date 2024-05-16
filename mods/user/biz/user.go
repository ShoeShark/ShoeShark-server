package biz

import (
	"github.com/shoe-shark/shoe-shark-service/mods/user/api/res"
	"github.com/shoe-shark/shoe-shark-service/mods/user/dao"
	"github.com/shoe-shark/shoe-shark-service/repository"
)

func GetUserByAccountAddress(accountAddress string) (*res.UserInfoRes, error) {

	rp := repository.GetPGRepository()
	user, err := dao.GetByAccountAddress(rp, accountAddress)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &res.UserInfoRes{
		Email:          user.Email,
		Username:       user.Username,
		AccountAddress: user.AccountAddress,
	}, nil
}
