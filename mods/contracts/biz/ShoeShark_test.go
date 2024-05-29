package biz

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shoe-shark/shoe-shark-service/config"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi_repository"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/dao"
	"github.com/shoe-shark/shoe-shark-service/repository"
	"github.com/shoe-shark/shoe-shark-service/repository/db"
	"math/big"
	"testing"
)

func TestMintWhiteList(t *testing.T) {

	config.InitConfig("../../../resources/")
	repository.Init()

	eth.InitClient(config.GetConfig())

	valueCtx := context.WithValue(context.Background(), "accountAddress", "0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69")
	err := MintWhiteList(&valueCtx)
	if err != nil {
		t.Error(err)
	}
}

func TestUpNft(t *testing.T) {

	config.InitConfig("../../../resources/")
	//logger.InitLogger()
	repository.Init()

	eth.InitClient(config.GetConfig())

	accountAddress := "0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69"
	client := eth.GetClient()
	privateKey := eth.GetPrivateKey()

	nftContractAddress, err := dao.GetContractAddressByContractName(db.GetPGRepository(), constants.ShoeSharkNft)
	if err != nil {
		return
	}

	nftRepository, err := abi_repository.NewShoeSharkNftRepository(client, nftContractAddress, privateKey)
	if err != nil {
		return
	}

	opts, err := eth.NewTransactOpts(privateKey, client)
	if err != nil {
		t.Error(err)
		return
	}

	//data := "https://white-left-chameleon-515.mypinata.cloud/ipfs/QmThBYjgDMGxwFnxxWH6WpQT8rM2RjphcXSvJcQWgRnJfF/3"
	//price := 20011

	data, err := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000004c60000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000006568747470733a2f2f77686974652d6c6566742d6368616d656c656f6e2d3531352e6d7970696e6174612e636c6f75642f697066732f516d546842596a67444d477877466e78785748365770515438724d32526a7068635853764a63515767526e4a66462f31000000000000000000000000000000000000000000000000000000")
	from0, err := nftRepository.Contract.SafeTransferFrom0(
		opts,
		common.HexToAddress(accountAddress), // 所属nft账户地址
		common.HexToAddress("0x09cE8D7E2eC68C4445c8a7D1B922ebc0EFCe2366"), // marketNft合约地址
		big.NewInt(2), // tokenId
		data,          // 价格+metaJson地址 hash256
	)
	if err != nil {
		t.Error(err)

	}

	fmt.Println("aa: ", from0.Hash().Hex())

}

func TestGetAllNftMarket(t *testing.T) {

	config.InitConfig("../../../resources/")
	//logger.InitLogger()
	repository.Init()

	eth.InitClient(config.GetConfig())

	//accountAddress := "0x7888b7B844B4B16c03F8daCACef7dDa0F5188645"
	client := eth.GetClient()
	//privateKey := eth.GetPrivateKey()

	nftContractAddress, err := dao.GetContractAddressByContractName(db.GetPGRepository(), constants.ShoeSharkNftMarket)
	if err != nil {
		return
	}

	contract, err := abi.NewShoeSharkNftMarket(common.HexToAddress(nftContractAddress), client)
	if err != nil {
		t.Error(err)
		return
	}

	ts, err := contract.GetAllNFTs(&bind.CallOpts{})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("ts: ", ts)

}
