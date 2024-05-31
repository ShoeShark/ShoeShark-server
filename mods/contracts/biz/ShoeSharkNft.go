package biz

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/dao"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	log "github.com/sirupsen/logrus"
)

func MintWhiteList(ctx *context.Context) ([]string, error) {
	accountAddress := (*ctx).Value("accountAddress").(string)

	rp := db.GetPGRepository()

	accounts, err := dao.GetLatestAccounts(rp)
	if err != nil {
		log.Error("un find nft contracts entity", "error", err)
		return nil, err
	}

	//account := common.HexToAddress(accountAddress)

	//eth.InitClient(nil)
	//client := eth.GetClient()
	//privateKey := eth.GetPrivateKey()

	//nftContractAddress, err := dao.GetContractAddressByContractName(rp, constants.ShoeSharkNft)
	//if err != nil {
	//	log.Error("un find nft contracts entity", "error", err)
	//	return err
	//}

	//nftRepository, err := abi_repository.NewShoeSharkNftRepository(client, nftContractAddress, privateKey)
	//if err != nil {
	//	log.Error("un find nft contracts entity", "error", err)
	//	return err
	//}

	proof, _ := eth.BuildMerkleProof(accountAddress, accounts)
	if proof == nil {
		log.Error("un find nft contracts entity", "error:  ", err)
		return nil, errors.New("not in white list")
	}

	//生成并打印指定地址的 Merkle Proof
	//for _, p := range proof {
	//	fmt.Println("Proof part:", hex.EncodeToString(p[:]))
	//}

	//err = nftRepository.MintWhitelist(account, proof)
	//if err != nil {
	//	log.Error("un find nft contracts entity", "error", err)
	//
	//	return err
	//}

	var r = make([]string, len(proof))

	for i, bytes := range proof {
		r[i] = hex.EncodeToString(bytes[:])
	}

	return r, nil
}
