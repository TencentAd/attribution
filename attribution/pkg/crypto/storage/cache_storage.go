package storage

import (
	"sync"
)

var CacheStorage sync.Map

// 该方法用于从本地缓存中拿到数据
func Storage(groupId string, encryptKey string) {
	CacheStorage.Store(groupId, encryptKey)
}

// 该方法用于将数据存入本地的缓存中
func Fetch(groupId string) (string, bool) {
	encryptKey, ok := CacheStorage.Load(groupId)
	return encryptKey.(string), ok
}
