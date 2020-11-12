package util

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/TencentAd/attribution/attribution/pkg/common/define"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/structure"
)

const defaultLength = 224

func GenerateEncKey(length int) (*big.Int, error) {
	l := defaultLength
	if length > 0 {
		l = length
	}
	encKey, err := rand.Prime(rand.Reader, l)
	if err != nil {
		return nil, err
	}
	// 确保prime是一个奇数
	tmp := big.NewInt(0)
	tmp.Add(tmp, encKey)
	if tmp.Mod(tmp, big.NewInt(2)) == big.NewInt(1) {
		encKey.Add(encKey, big.NewInt(1))
	}
	return encKey, nil
}

func Hex2BigInt(hex string) (*big.Int, error) {
	bi, ok := big.NewInt(0).SetString(hex, 16)
	if !ok {
		return nil, fmt.Errorf("not valid big integer[%s]", hex)
	}

	return bi, nil
}

func HexMust2BigInt(hex string) *big.Int {
	bi, ok := big.NewInt(0).SetString(hex, 16)
	if !ok {
		panic(fmt.Errorf("hex[%s] not valid big integer", hex))
	}

	return bi
}

func BigInt2Hex(bi *big.Int) string {
	return bi.Text(16)
}

func GenerateCryptoPair(length int) (*structure.CryptoPair, error) {
	encKey, err := GenerateEncKey(0)
	if err != nil {
		return nil, err
	}

	return &structure.CryptoPair{
		EncKey: encKey,
		DecKey: FindDecKey(encKey, define.Prime),
	}, nil
}

func CreateCryptoPair(encKey *big.Int) *structure.CryptoPair {
	return &structure.CryptoPair{
		EncKey: encKey,
		DecKey: FindDecKey(encKey, define.Prime),
	}
}