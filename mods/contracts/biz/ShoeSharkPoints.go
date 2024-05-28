package biz

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi_repository"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/dao"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	log "github.com/sirupsen/logrus"
	"math/big"
)

func GetPoints(ctx *context.Context) (*big.Int, error) {
	accountAddress := (*ctx).Value("accountAddress").(string)

	nftContractAddress, err := dao.GetContractAddressByContractName(db.GetPGRepository(), constants.ShoeSharkRewardPoint)
	if err != nil {
		log.Error("un find nft contracts entity", "error", err)
		return big.NewInt(0), err
	}

	pointsAbiRp, err := abi_repository.NewShoeSharkRewardPointRepository(eth.GetClient(), nftContractAddress, eth.GetPrivateKey())
	if err != nil {
		log.Error("un find abi reward points entity", "error", err)
		return big.NewInt(0), err
	}

	points := pointsAbiRp.GetPoints(common.HexToAddress(accountAddress))
	return points, nil

}
