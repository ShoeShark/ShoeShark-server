package biz

import (
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/dao"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func InsertContractTransaction(accountAddress, txnHash, method string) error {
	return dao.InsertContractTransaction(db.GetPGRepository(), accountAddress, txnHash, method)
}
