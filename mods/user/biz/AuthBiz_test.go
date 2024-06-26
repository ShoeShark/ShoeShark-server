package biz

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shoe-shark/shoe-shark-service/pkg/ethers"
	"testing"
)

func TestGenerateAndVerifySignature(t *testing.T) {
	privateKeyHex := "7111ec7d38f35eaa460c85991cd269ee4f39f567ffb587d4e6aa474b38dccb7e"

	// 将十六进制字符串转换为私钥对象
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		t.Fatalf("Failed to convert hex string to private key: %v", err)
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// 将公钥转换为以太坊地址
	accountAddress := crypto.PubkeyToAddress(*publicKey).Hex()
	fmt.Println("accountAddress:", accountAddress)

	// 生成签名
	nonce := "19594ac9f53c19eb0af408287426635b"

	signature, err := GenerateSignature(privateKey, nonce)
	fmt.Println("signature:", signature)
	if err != nil {
		t.Fatalf("Failed to generate signature: %v", err)
	}

	// 验证签名
	isVaild := ethers.VerifySignature(accountAddress, signature, nonce)
	if !isVaild {
		t.Fatalf("Signature verification failed")
	}
}

func GenerateSignature(privateKey *ecdsa.PrivateKey, nonce string) (string, error) {
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonce), nonce)
	messageHash := crypto.Keccak256Hash([]byte(message))

	signature, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %v", err)
	}

	// Adjust the V value (0 or 1) to 27 or 28
	signature[64] += 27

	return hexutil.Encode(signature), nil
}
