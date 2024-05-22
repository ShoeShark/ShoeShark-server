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
)

func InitClient(cfg *config.Config) {
	var err error
	clientInstance, err = ethclient.DialContext(context.Background(), "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the Ethereum client: %v", err))
	}
}

func GetClient() *ethclient.Client {
	return clientInstance
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
