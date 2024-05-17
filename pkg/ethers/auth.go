package ethers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateNonce() (string, error) {
	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	return hex.EncodeToString(nonce), nil
}

func VerifySignature(addressHex string, signatureHex string, nonce string) bool {
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonce), nonce)
	messageHash := crypto.Keccak256Hash([]byte(message))

	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		fmt.Printf("Failed to decode signature: %v", err)
		return false
	}

	// Ensure the signature length is correct (65 bytes)
	if len(signature) != 65 {
		fmt.Println("Invalid signature length")
		return false
	}
	// Recover the public key
	signature[64] -= 27 // Adjust the V value (27 or 28) to 0 or 1
	recoveredPubKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		fmt.Printf("Failed to recover public key: %v\n", err)
		return false
	}

	// Convert the recovered public key to an Ethereum address
	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey)

	// Compare the recovered address with the provided address
	return recoveredAddress.Hex() == addressHex
}
