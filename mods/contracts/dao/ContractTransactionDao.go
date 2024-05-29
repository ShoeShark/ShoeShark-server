package dao

import (
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func InsertContractTransaction(rp *db.Repository, accountAddress, txnHash, method string) error {
	txn := schema.ContractTransaction{
		AccountAddress: accountAddress,
		TxnHash:        txnHash,
		Method:         method,
	}
	return rp.Create(&txn).Error
}
