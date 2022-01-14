package utils

import (
	"crypto/rand"
	"math/big"
)

const maxInt = int64(^uint64(0) >> 1)

func GenerateRandomSalt() int64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(maxInt))
	return nBig.Int64()
}
