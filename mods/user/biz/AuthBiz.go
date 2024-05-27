package biz

import (
	"errors"
	"github.com/shoe-shark/shoe-shark-service/middleware"
	"github.com/shoe-shark/shoe-shark-service/mods/user/dao"
	"github.com/shoe-shark/shoe-shark-service/mods/user/schema"
	"github.com/shoe-shark/shoe-shark-service/pkg/ethers"
	"github.com/shoe-shark/shoe-shark-service/repository"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func GetNonce(accountAddress string) (string, error) {
	rp := repository.GetPGRepository()
	nonce, err := dao.GetNonceByAccountAddress(rp, accountAddress)
	if err != nil {
		return "", err
	}

	if nonce == "" {
		generateNonce, err := ethers.GenerateNonce()
		if err != nil {
			return "", err
		}

		// 创建用户
		var createUser = schema.User{
			AccountAddress: accountAddress,
			Nonce:          generateNonce,
		}

		err = rp.SaveOne(&createUser)
		if err != nil {
			return "", err
		}

		return generateNonce, nil
	}

	return nonce, nil
}

func VerifySignature(accountAddress string, signature string) (string, error) {
	rp := repository.GetPGRepository()
	nonce, err := dao.GetNonceByAccountAddress(rp, accountAddress)
	if err != nil {
		return "", err
	}

	signValid := ethers.VerifySignature(accountAddress, signature, nonce)
	if !signValid {
		return "", errors.New("invalid signature")
	}

	// 更新nonce
	err = generateNonceAndUpdate(rp, accountAddress)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(accountAddress)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateNonceAndUpdate(rp *db.Repository, accountAddress string) error {
	// 更新nonce
	generateNonce, err := ethers.GenerateNonce()
	if err != nil {
		return err
	}
	err = dao.UpdateNonceByAccountAddress(rp, accountAddress, generateNonce)
	if err != nil {
		return err
	}
	return err
}
