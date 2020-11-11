package uitl

import (
	"crypto/rand"
	"math/big"
)

const defaultLength = 1024

func GenerateEncKey(length int) (*big.Int, error) {
	l := defaultLength
	if length > 0 {
		l = length
	}
	encKey, err := rand.Prime(rand.Reader, l)
	if err != nil {
		return nil, nil
	}
	// 确保prime是一个奇数
	tmp := big.NewInt(0)
	tmp.Add(tmp, encKey)
	if tmp.Mod(tmp, big.NewInt(2)) == big.NewInt(1) {
		encKey.Add(encKey, big.NewInt(1))
	}
	return encKey, nil
}
