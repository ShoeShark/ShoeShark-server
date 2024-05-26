package biz

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi_repository"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/dao"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	log "github.com/sirupsen/logrus"
)

func MintWhiteList(ctx *context.Context) error {
	accountAddress := (*ctx).Value("accountAddress").(string)

	rp := db.GetPGRepository()

	accounts, err := dao.GetLatestAccounts(rp)
	if err != nil {
		log.Error("un find nft contracts entity", "error", err)
		return err
	}

	account := common.HexToAddress(accountAddress)

	eth.InitClient(nil)
	client := eth.GetClient()
	privateKey := eth.GetPrivateKey()

	nftContractAddress, err := dao.GetContractAddressByContractName(rp, constants.ShoeSharkNft)
	if err != nil {
		log.Error("un find nft contracts entity", "error", err)
		return err
	}

	nftRepository, err := abi_repository.NewShoeSharkNftRepository(client, nftContractAddress, privateKey)
	if err != nil {
		log.Error("un find nft contracts entity", "error", err)
		return err
	}

	proof := eth.GenerateMerkleProof(accountAddress, accounts)
	if proof == nil {
		log.Error("un find nft contracts entity")
		return errors.New("un find nft contracts entity")
	}

	// 生成并打印指定地址的 Merkle Proof
	for _, p := range proof {
		fmt.Println("Proof part:", hex.EncodeToString(p[:]))
	}

	err = nftRepository.MintWhitelist(account, proof)
	if err != nil {
		log.Error("un find nft contracts entity", "error", err)

		return err
	}

	return nil
}
