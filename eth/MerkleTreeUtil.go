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

// buildMerkleTree constructs a Merkle tree and returns the root and layers.
func BuildMerkleTree(elements []string) ([32]byte, [][][32]byte) {
	var leaves [][32]byte
	for _, elem := range elements {
		data, err := hex.DecodeString(elem[2:])
		if err != nil {
			fmt.Println("Error decoding string:", err)
			continue
		}
		hashed := Keccak256(data)
		leaves = append(leaves, hashed)
	}

	var layers [][][32]byte
	layers = append(layers, leaves)

	for len(leaves) > 1 {
		var nextLevel [][32]byte
		for i := 0; i < len(leaves); i += 2 {
			if i+1 < len(leaves) {
				// 对两个节点进行排序
				if bytes.Compare(leaves[i][:], leaves[i+1][:]) > 0 {
					leaves[i], leaves[i+1] = leaves[i+1], leaves[i]
				}
				combined := append(leaves[i][:], leaves[i+1][:]...)
				combinedHash := Keccak256(combined)
				nextLevel = append(nextLevel, combinedHash)
			} else {
				nextLevel = append(nextLevel, leaves[i])
			}
		}
		leaves = nextLevel
		layers = append(layers, nextLevel)
	}

	if len(leaves) == 0 {
		return [32]byte{}, layers
	}
	return leaves[0], layers
}

// 根据给定的账户字符串和账户数组构建Merkle树并生成证明
func GenerateMerkleProof(account string, accounts []string) [][32]byte {
	_, layers := BuildMerkleTree(accounts)
	index := -1
	for i, acc := range accounts {
		if acc == account {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Account not found")
		return nil
	}

	var proof [][32]byte
	for _, layer := range layers[:len(layers)-1] {
		pairIndex := index ^ 1
		if pairIndex < len(layer) {
			if bytes.Compare(layer[index][:], layer[pairIndex][:]) > 0 {
				pairIndex = index // 如果当前节点的哈希大于兄弟节点的哈希，取当前节点为证明节点
				index = index ^ 1
			}
			proof = append(proof, layer[pairIndex])
		}
		index /= 2
	}
	return proof // 返回证明, Merkle根和所有层
}
