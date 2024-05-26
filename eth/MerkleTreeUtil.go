package eth

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
)

// Keccak256 hashes using Keccak-256 algorithm, returning a 32-byte hash.
func Keccak256(data []byte) [32]byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	var hash [32]byte
	copy(hash[:], hasher.Sum(nil))
	return hash
}

// BuildMerkleTree constructs a Merkle tree and returns the root and layers.
func BuildMerkleTree(elements []string) ([32]byte, [][][32]byte, error) {
	var leaves [][32]byte
	for _, elem := range elements {
		data, err := hex.DecodeString(elem[2:])
		if err != nil {
			return [32]byte{}, nil, fmt.Errorf("error decoding string: %v", err)
		}
		hashed := Keccak256(data)
		leaves = append(leaves, hashed)
	}

	var layers [][][32]byte
	layers = append(layers, leaves)

	for len(leaves) > 1 {
		var nextLevel [32]byte
		var newLevel [][32]byte

		for i := 0; i < len(leaves); i += 2 {
			if i+1 < len(leaves) {
				left := leaves[i]
				right := leaves[i+1]
				if bytes.Compare(left[:], right[:]) > 0 {
					left, right = right, left
				}
				combined := append(left[:], right[:]...)
				nextLevel = Keccak256(combined)
				newLevel = append(newLevel, nextLevel)
			} else {
				newLevel = append(newLevel, leaves[i])
			}
		}
		leaves = newLevel
		layers = append(layers, newLevel)
	}

	if len(leaves) == 0 {
		return [32]byte{}, layers, nil
	}
	return leaves[0], layers, nil
}

// BuildMerkleProof constructs a Merkle proof for the given account string from the list of accounts.
func BuildMerkleProof(account string, accounts []string) ([][32]byte, error) {
	_, layers, err := BuildMerkleTree(accounts)
	if err != nil {
		return nil, err
	}

	index := -1
	for i, acc := range accounts {
		if acc == account {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, fmt.Errorf("account not found")
	}

	var proof [][32]byte
	for _, layer := range layers[:len(layers)-1] {
		pairIndex := index ^ 1
		if pairIndex < len(layer) {
			proof = append(proof, layer[pairIndex])
		}
		index /= 2
	}
	return proof, nil
}
