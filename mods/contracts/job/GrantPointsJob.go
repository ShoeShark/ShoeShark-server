package job

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi_repository"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/dao"
	"github.com/shoe-shark/shoe-shark-service/repository"
	log "github.com/sirupsen/logrus"
)

type ShoeSharkContractJob struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
}

func (cj *ShoeSharkContractJob) startGrantPointsJob() {
	rp := repository.GetPGRepository()
	rewardPointContractAddress, err := dao.GetContractAddressByContractName(rp, constants.ShoeSharkRewardPoint)
	if err != nil {
		log.Error("un find reward point contract entity", "error", err)
	}
	nftContractAddress, err := dao.GetContractAddressByContractName(rp, constants.ShoeSharkNft)
	if err != nil {
		log.Error("un find nft contract entity", "error", err)
	}

	pointBiz, err := abi_repository.NewShoeSharkRewardPointRepository(cj.client, rewardPointContractAddress, cj.privateKey)
	if err != nil {
		log.Error("init ShoeSharkRewardPointRepository error: ", err)
	}

	nftBiz, err := abi_repository.NewShoeSharkNftRepository(cj.client, nftContractAddress, cj.privateKey)
	if err != nil {
		log.Error("init ShoeSharkNftRepository error: ", err)
	}

	// TODO
	pointBiz.SetPoints(nil, nil)

	var merkleRoot [32]byte

	err = nftBiz.SetMerkleRoot(merkleRoot)
}
