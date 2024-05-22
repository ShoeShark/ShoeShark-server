package biz

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/contracts"
)

type SimpleStorageService struct {
	client      *ethclient.Client
	contract    *contracts.ShoeSharkRewardPoint
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
}

func NewSimpleStorageService(client *ethclient.Client, contractAddress string, privateKey string) (*SimpleStorageService, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)

	contract, err := contracts.NewShoeSharkRewardPoint(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}

	return &SimpleStorageService{
		client:      client,
		contract:    contract,
		privateKey:  pk,
		fromAddress: fromAddress,
	}, nil
}
