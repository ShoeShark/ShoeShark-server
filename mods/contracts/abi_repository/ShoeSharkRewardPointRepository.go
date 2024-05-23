package abi_repository

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type ShoeSharkRewardPointRepository struct {
	client      *ethclient.Client
	contract    *abi.ShoeSharkRewardPoint
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
}

func NewShoeSharkRewardPointRepository(client *ethclient.Client, contractAddress string, privateKey *ecdsa.PrivateKey) (*ShoeSharkRewardPointRepository, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	contract, err := abi.NewShoeSharkRewardPoint(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	return &ShoeSharkRewardPointRepository{
		client:      client,
		contract:    contract,
		privateKey:  privateKey,
		fromAddress: fromAddress,
	}, nil
}

func (biz *ShoeSharkRewardPointRepository) SetPoints(accounts []*common.Address, points []*big.Int) error {
	contract := biz.contract

	transactOpts, err := eth.NewTransactOpts(biz.privateKey, biz.client)
	if err != nil {
		return err
	}

	accountValues := make([]common.Address, len(accounts))
	for i, acc := range accounts {
		accountValues[i] = *acc
	}

	transaction, err := contract.SetPoints(transactOpts, accountValues, points)
	if err != nil {
		return err
	}

	log.Info("SetPoints transaction successfully", transaction)

	return nil
}
