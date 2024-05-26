package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/config"
	"math/big"
)

var (
	clientInstance *ethclient.Client
	privateKey     *ecdsa.PrivateKey
)

func InitClient(cfg *config.Config) {
	var err error
	clientInstance, err = ethclient.DialContext(context.Background(), "https://sepolia.infura.io/v3/599c8e1c92a54659b339ecbaad80c39c")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the Ethereum client: %v", err))
	}

	privateKey, err = crypto.HexToECDSA("7111ec7d38f35eaa460c85991cd269ee4f39f567ffb587d4e6aa474b38dccb7e")
	if err != nil {
		panic(fmt.Sprintf("Failed to  init privateKey: %v", err))
	}
}

func GetClient() *ethclient.Client {
	return clientInstance
}

func GetPrivateKey() *ecdsa.PrivateKey {
	return privateKey
}

// NewTransactOpts creates a new instance of TransactOpts with the given private key
func NewTransactOpts(privateKey *ecdsa.PrivateKey, client *ethclient.Client) (*bind.TransactOpts, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // default value
	auth.GasLimit = uint64(3000000) // default gas limit
	auth.GasPrice = gasPrice

	return auth, nil
}
