package abi_repository

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
	"testing"
)

func TestGrantPoints(t *testing.T) {
	eth.InitClient(nil)

	client := eth.GetClient()
	privateKey := eth.GetPrivateKey()

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
	eth.InitClient(nil)

	client := eth.GetClient()
	privateKey := eth.GetPrivateKey()

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
