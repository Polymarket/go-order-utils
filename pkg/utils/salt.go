package utils

import (
	"crypto/rand"
	"math"
	"math/big"
)

func GenerateRandomSalt() int64 {
	maxInt := math.Pow(2, 32)
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(maxInt)))
	return nBig.Int64()
}
