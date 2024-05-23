package dao

import (
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/schema"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
)

func GetContractAddressByContractName(rp *db.Repository, contractName constants.ShoeSharkContractName) (string, error) {
	contract := schema.Contracts{}
	err := rp.Where("contract_name = ?", string(contractName)).Select([]string{"contract_address"}).First(&contract).Error
	if err != nil {
		return "", err
	}

	return contract.ContractAddress, nil
}
