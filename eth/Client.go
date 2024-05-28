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
	"os"
)

var (
	clientInstance *ethclient.Client
	privateKey     *ecdsa.PrivateKey
)

func InitClient(cfg *config.Config) {
	var err error
	//clientInstance, err = ethclient.DialContext(context.Background(), "https://sepolia.infura.io/v3/599c8e1c92a54659b339ecbaad80c39c")
	clientInstance, err = ethclient.DialContext(context.Background(), "https://avalanche-fuji.infura.io/v3/599c8e1c92a54659b339ecbaad80c39c")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the Ethereum client: %v", err))
	}

	privateKey, err = crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
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

	// 获取最新的区块以确定合适的 baseFee
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	baseFee := header.BaseFee

	// 假设我们愿意支付基础费用的2倍，加上1 Gwei的小费
	tip := big.NewInt(1e9) // 1 Gwei
	gasFeeCap := new(big.Int).Mul(baseFee, big.NewInt(2))
	gasFeeCap.Add(gasFeeCap, tip)

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
	auth.GasFeeCap = gasFeeCap
	auth.GasTipCap = tip

	return auth, nil
}
