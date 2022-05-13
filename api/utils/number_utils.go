package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateNumberInRange(min, max int64) int64 {
	bg := big.NewInt(max - min)
	n, err := rand.Int(rand.Reader, bg)

	if err != nil {
		GenerateNumberInRange(min, max)
	}

	return n.Int64() + min
}
