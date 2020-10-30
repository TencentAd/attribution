package kv

import (
    "fmt"
    "time"

    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/aerospike"
    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/leveldb"
    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/redis"
)

const (
    StorageTypeLevelDB   StorageType = "LEVELDB"
    StorageTypeRedis     StorageType = "REDIS"
    StorageTypeAerospike StorageType = "AEROSPIKE"
)

var (
    DefaultExpiration = 7*24*time.Hour
    Prefix = "impression::"
)

type KV interface {
    Has(string) (bool, error)
    Set(string) error
}

type StorageType string

type Option struct {
    Address  string          `json:"address"`
    Password string          `json:"password"`
    Expiration time.Duration `json:"expiration"`
}

func CreateKV(storageType StorageType, option *Option) (KV, error) {
    if option.Expiration == 0 {
        option.Expiration = DefaultExpiration
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

