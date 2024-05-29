package job

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shoe-shark/shoe-shark-service/config"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi_repository"
	"github.com/shoe-shark/shoe-shark-service/repository"
	"testing"
)

func TestGrantPointsJob(t *testing.T) {

	config.InitConfig("../../../resources/")
	repository.Init()
	eth.InitClient(nil)

	client, privateKey := eth.GetClient(), eth.GetPrivateKey()

	job := &ShoeSharkContractJob{
		client:     client,
		privateKey: privateKey,
	}

	job.startGrantPointsJob()

}

func TestSetRoot(t *testing.T) {
	eth.InitClient(nil)

	client := eth.GetClient()
	privateKey := eth.GetPrivateKey()

	nftBiz, err := abi_repository.NewShoeSharkNftRepository(client, "0x87f949c651338d1c8561f29dabef90f8dfb3dfda", privateKey)
	if err != nil {
		panic(err)
	}

	accounts := make([]string, 0)
	accounts = append(accounts, "0xc44cE209A984135B8C34b1B8408C43e2FDbB282E")
	accounts = append(accounts, "0x15Bb92E888cFF5ab5b557D4E75cE01eAEE1a9caA")
	accounts = append(accounts, "0xaEA9D21C154aF3027d7bdb7c4EF74FF3A70D077D")

	root, laylers, _ := eth.BuildMerkleTree(accounts)
	fmt.Println("root: 0x", hex.EncodeToString(root[:]))
	for i, layler := range laylers {
		fmt.Println("layler:", i)
		for l := range layler {
			fmt.Println("layler", hex.EncodeToString(layler[l][:]))
		}
	}
	//err = nftBiz.SetMerkleRoot(root)
	//if err != nil {
	//	panic(err)
	//}

	//proof, _ := eth.BuildMerkleProof("0xc44cE209A984135B8C34b1B8408C43e2FDbB282E", accounts)
	//
	//for _, p := range proof {
	//	fmt.Println("Proof part: 0x", hex.EncodeToString(p[:]))
	//}
	//
	//err = nftBiz.MintWhitelist(common.HexToAddress("0xc44cE209A984135B8C34b1B8408C43e2FDbB282E"), proof)
	//if err != nil {
	//	panic(err)
	//}

	//minted, err := nftBiz.Contract.SHasMinted(&bind.CallOpts{}, common.HexToAddress("0xc44cE209A984135B8C34b1B8408C43e2FDbB282E"))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println("minted: ", minted)
	//
	//sRoot, err := nftBiz.Contract.SRoot(&bind.CallOpts{})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println("sRoot: ", hex.EncodeToString(sRoot[:]))

	owner, err := nftBiz.Contract.GetOwner(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println("owner:", owner.Hex())
}
