package storage

import "errors"

type Cache struct {
	CacheStorage map[string]string
}

func NewCacheStorage() *Cache {
	return &Cache{CacheStorage: make(map[string]string)}
}

// 该方法用于从本地缓存中拿到数据
func (cache *Cache) Storage(groupId string, encryptKey string) error {
	if _, ok := cache.CacheStorage[groupId]; ok {
		return errors.New("key already exists")
	}
	cache.CacheStorage[groupId] = encryptKey
	return nil
}

// 该方法用于将数据存入本地的缓存中
func (cache *Cache) Fetch(groupId string) (string, error) {
	if encryptKey, ok := cache.CacheStorage[groupId]; ok {
		return encryptKey, nil
	}
	return "", errors.New("encryptKey does not exist")
}
