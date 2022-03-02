package main

import (
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func main() {
	key, _ := crypto.GenerateKey()

	address := crypto.PubkeyToAddress(key.PublicKey).Hex()

	privateKey := hex.EncodeToString(key.D.Bytes())

	fmt.Println(address)
	fmt.Println(privateKey)

	key.D, _ = new(big.Int).SetString(privateKey, 16)
	key.PublicKey.Curve = elliptic.P256()
	key.PublicKey.X, key.PublicKey.Y = key.PublicKey.Curve.ScalarBaseMult(key.D.Bytes())
	pubKey := fmt.Sprintf("%x", elliptic.Marshal(secp256k1.S256(), key.X, key.Y))
	fmt.Println(pubKey)
}

// func setZeroes(matrix [][]int) {
// 	m, n := len(matrix), len(matrix[0])

// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == 0 {
// 				for k := 0; k < n; k++ {
// 					matrix[i][k] = 0
// 				}
// 			}
// 		}
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if matrix[j][i] == 0 {
// 				for k := 0; k < m; k++ {
// 					matrix[k][i] = 0
// 				}
// 			}
// 		}
// 	}
// }

// func checkZero(row []int) []int {
// 	n := len(row)
// 	var ans []int
// 	for i := 0; i < n; i++ {
// 		if row[i] == 0 {
// 			ans = append(ans, i)
// 		}
// 	}

// 	return ans
// }

// func setRowZero(m [][]int, rowIdx int) {
// 	for i := 0; i < len(m[0]); i++ {
// 		m[rowIdx][i] = 0
// 	}
// }

// func setColZero(m [][]int, colIdx int) {
// 	for i := 0; i < len(m); i++ {
// 		m[i][colIdx] = 0
// 	}
// }
