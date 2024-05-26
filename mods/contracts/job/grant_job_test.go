package job

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/config"
	"github.com/shoe-shark/shoe-shark-service/logger"
	"github.com/shoe-shark/shoe-shark-service/repository"
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

func TestGrantPointsJob(t *testing.T) {

	config.InitConfig("../../../resources/application.dev.yml")
	logger.InitLogger()
	repository.Init()

	client, privateKey, err := GetClientAndPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	job := &ShoeSharkContractJob{
		client:     client,
		privateKey: privateKey,
	}

	job.startGrantPointsJob()

}
