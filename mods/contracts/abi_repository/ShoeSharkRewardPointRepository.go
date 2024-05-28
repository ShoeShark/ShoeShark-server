package abi_repository

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type ShoeSharkRewardPointRepository struct {
	client     *ethclient.Client
	contract   *abi.ShoeSharkRewardPoint
	privateKey *ecdsa.PrivateKey
}

func NewShoeSharkRewardPointRepository(client *ethclient.Client, contractAddress string, privateKey *ecdsa.PrivateKey) (*ShoeSharkRewardPointRepository, error) {
	contract, err := abi.NewShoeSharkRewardPoint(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	return &ShoeSharkRewardPointRepository{
		client:     client,
		contract:   contract,
		privateKey: privateKey,
	}, nil
}

func (biz *ShoeSharkRewardPointRepository) GetPoints(account common.Address) *big.Int {
	contract := biz.contract

	points, err := contract.GetUserPoints(&bind.CallOpts{}, account)
	if err != nil {
		log.Error("Get user:", account, " points error:", err)
		return big.NewInt(0)
	}

	return points
}

func (biz *ShoeSharkRewardPointRepository) SetPoints(accounts []common.Address, points []*big.Int) error {
	contract := biz.contract

	transactOpts, err := eth.NewTransactOpts(biz.privateKey, biz.client)
	if err != nil {
		return err
	}

	transaction, err := contract.SetPoints(transactOpts, accounts, points)
	if err != nil {
		return err
	}

	log.Info("SetPoints transaction successfully txHash: ", transaction.Hash().Hex())

	fmt.Println("setPoints transaction successfully txHash: ", transaction.Hash().Hex())
	return nil
}
