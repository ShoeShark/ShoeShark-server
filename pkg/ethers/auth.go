package ethers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
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
	messageHash := crypto.Keccak256Hash([]byte(nonce))
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		fmt.Printf("Failed to decode signature: %v", err)
		return false
	}

	recoveredPubKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		fmt.Printf("Failed to recover public key: %v", err)
		return false
	}

	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey)
	return recoveredAddress.Hex() == addressHex
}
