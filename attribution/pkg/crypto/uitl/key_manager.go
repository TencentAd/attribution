package uitl

import "math/big"

// 接口，实现类有redis的KeyManager和hdfs的KeyManager
type KeyManager interface {
	GetEncryptKey(string) (*big.Int, error)
}