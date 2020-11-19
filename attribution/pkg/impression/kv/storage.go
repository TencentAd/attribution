package kv

import (
    "fmt"
    
    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/aerospike"
    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/leveldb"
    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/opt"
    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/redis"
)

const (
    StorageTypeLevelDB   StorageType = "LEVELDB"
    StorageTypeRedis     StorageType = "REDIS"
    StorageTypeAerospike StorageType = "AEROSPIKE"
)

type KV interface {
    Get(string) (string, error)
    Set(string, string) error
}

type StorageType string

func CreateKV(storageType StorageType, option *opt.Option) (KV, error) {
    if option.Expiration == 0 {
        option.Expiration = opt.DefaultExpiration
    }

    switch storageType {
    case StorageTypeLevelDB:
        return leveldb.New(option)
    case StorageTypeRedis:
        return redis.New(option)
    case StorageTypeAerospike:
        return aerospike.New(option)
    default:
        return nil, fmt.Errorf("not support storage type")
    }
}

