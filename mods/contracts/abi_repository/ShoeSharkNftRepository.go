package abi_repository

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi"
	log "github.com/sirupsen/logrus"
)

type ShoeSharkNftRepository struct {
	client     *ethclient.Client
	contract   *abi.ShoeSharkNft
	privateKey *ecdsa.PrivateKey
}

func NewShoeSharkNftRepository(client *ethclient.Client, contractAddress string, privateKey *ecdsa.PrivateKey) (*ShoeSharkNftRepository, error) {
	contract, err := abi.NewShoeSharkNft(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	return &ShoeSharkNftRepository{
		client:     client,
		contract:   contract,
		privateKey: privateKey,
	}, nil
}

func (biz *ShoeSharkNftRepository) SetMerkleRoot(merkleRoot [32]byte) error {
	contract := biz.contract

	transactOpts, err := eth.NewTransactOpts(biz.privateKey, biz.client)
	if err != nil {
		return err
	}

	transaction, err := contract.SetRoot(transactOpts, merkleRoot)
	if err != nil {
		return err
	}

	log.Info("Set Root transaction", transaction)

	return nil
}
