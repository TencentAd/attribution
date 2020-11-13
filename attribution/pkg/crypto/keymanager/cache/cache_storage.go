package cache

import (
	"sync"

	"github.com/TencentAd/attribution/attribution/pkg/crypto/structure"
)

type CryptoCache struct {
	m sync.Map
}

func NewCryptoCache() *CryptoCache {
	return &CryptoCache{}
}

func (c *CryptoCache) Store(groupId string, pair *structure.CryptoPair) {
	m.Store(groupId, pair)
}

// 该方法用于将数据存入本地的缓存中
func (c *CryptoCache) Load(groupId string) (*structure.CryptoPair, bool) {
	pair, ok := m.Load(groupId)
	if !ok {
		return nil, ok
	}
	return pair.(*structure.CryptoPair), ok
}


var m sync.Map

// 该方法用于从本地缓存中拿到数据
func Store(groupId string, encryptKey string) {
	m.Store(groupId, encryptKey)
}

// 该方法用于将数据存入本地的缓存中
func Fetch(groupId string) (string, bool) {
	encryptKey, ok := m.Load(groupId)
	return encryptKey.(string), ok
}
