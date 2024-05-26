package abi_repository

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi"
	log "github.com/sirupsen/logrus"
)

type ShoeSharkNftRepository struct {
	client     *ethclient.Client
	Contract   *abi.ShoeSharkNft
	privateKey *ecdsa.PrivateKey
}

func NewShoeSharkNftRepository(client *ethclient.Client, contractAddress string, privateKey *ecdsa.PrivateKey) (*ShoeSharkNftRepository, error) {
	contract, err := abi.NewShoeSharkNft(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	return &ShoeSharkNftRepository{
		client:     client,
		Contract:   contract,
		privateKey: privateKey,
	}, nil
}

func (biz *ShoeSharkNftRepository) SetMerkleRoot(merkleRoot [32]byte) error {
	contract := biz.Contract

	transactOpts, err := eth.NewTransactOpts(biz.privateKey, biz.client)
	if err != nil {
		return err
	}

	transaction, err := contract.SetRoot(transactOpts, merkleRoot)
	if err != nil {
		return err
	}

	log.Info("txHash: ", transaction.Hash().Hex())

	log.Info("Set Root transaction", transaction)

	return nil
}

func (biz *ShoeSharkNftRepository) MintWhitelist(account common.Address, proof [][32]byte) error {
	contract := biz.Contract

	transactOpts, err := eth.NewTransactOpts(biz.privateKey, biz.client)
	if err != nil {
		return err
	}

	if transactOpts == nil {
		return errors.New("transation is null")
	}

	transaction, err := contract.MintWhitelist(transactOpts, account, proof)
	if err != nil {
		return err
	}

	log.Info("MintWhitelist success: ", transaction)

	return nil
}
