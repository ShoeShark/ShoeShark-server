package abi_repository

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
	"testing"
)

func GetClientAndPrivateKey() (*ethclient.Client, *ecdsa.PrivateKey, error) {
	client, err := ethclient.DialContext(context.Background(), "https://sepolia.infura.io/v3/599c8e1c92a54659b339ecbaad80c39c")
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA("7111ec7d38f35eaa460c85991cd269ee4f39f567ffb587d4e6aa474b38dccb7e")
	if err != nil {
		return nil, nil, err
	}

	return client, privateKey, nil
}

func TestGrantPoints(t *testing.T) {

	client, privateKey, err := GetClientAndPrivateKey()
	if err != nil {
		t.Fatal("gen client error: ", err)
	}

	var rewardPointContractAddress = common.HexToAddress("0xCA6b7c0cBE1134565555Eebf50294af4316cB295")
	rewardPointContract, err := abi.NewShoeSharkRewardPoint(rewardPointContractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a SimpleStorage contracts: %v", err)
	}

	var pointRp = &ShoeSharkRewardPointRepository{
		client:     client,
		privateKey: privateKey,
		contract:   rewardPointContract,
	}

	accounts := make([]common.Address, 0)
	accounts = append(accounts, common.HexToAddress("0xc44cE209A984135B8C34b1B8408C43e2FDbB282E"))
	accounts = append(accounts, common.HexToAddress("0x15Bb92E888cFF5ab5b557D4E75cE01eAEE1a9caA"))

	points := make([]*big.Int, 0)
	points = append(points, big.NewInt(100))
	points = append(points, big.NewInt(100))

	err = pointRp.SetPoints(accounts, points)
	if err != nil {
		log.Fatal("Set points error", err)
	}
}

func TestSetMerkleRoot(t *testing.T) {
	client, privateKey, err := GetClientAndPrivateKey()
	if err != nil {
		log.Error("gen client error: ", err)
	}

	var nftContractAddress = common.HexToAddress("0xCA6b7c0cBE1134565555Eebf50294af4316cB295")
	nftContract, err := abi.NewShoeSharkNft(nftContractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a SimpleStorage contracts: %v", err)
	}

	var nftRp = &ShoeSharkNftRepository{
		client:     client,
		privateKey: privateKey,
		Contract:   nftContract,
	}

	accounts := make([]string, 0)
	accounts = append(accounts, "0xc44cE209A984135B8C34b1B8408C43e2FDbB282E")
	accounts = append(accounts, "0x15Bb92E888cFF5ab5b557D4E75cE01eAEE1a9caA")

	root, _, _ := eth.BuildMerkleTree(accounts)

	err = nftRp.SetMerkleRoot(root)
	if err != nil {
		log.Fatal("Set merkle root error", err)
	}
}

func TestGenSetMerkleRoot(t *testing.T) {
	accounts := make([]string, 0)
	accounts = append(accounts, "0xc44cE209A984135B8C34b1B8408C43e2FDbB282E")
	accounts = append(accounts, "0x15Bb92E888cFF5ab5b557D4E75cE01eAEE1a9caA")

	root, _, _ := eth.BuildMerkleTree(accounts)

	fmt.Println(root)
	// 将 [32]byte 转换为十六进制字符串
	rootHex := hex.EncodeToString(root[:])

	// 打印十六进制字符串，适合直接使用在 Solidity 中
	fmt.Println("Merkle Tree Root for Solidity:", "0x"+rootHex)

	// 生成并打印指定地址的 Merkle Proof
	proof, _ := eth.BuildMerkleProof("0xc44cE209A984135B8C34b1B8408C43e2FDbB282E", accounts) // 为第一个地址生成 proof
	for _, p := range proof {
		fmt.Println("Proof part:", hex.EncodeToString(p[:]))
	}
}
