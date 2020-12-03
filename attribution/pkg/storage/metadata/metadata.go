package metadata

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/TencentAd/attribution/attribution/pkg/storage/metadata/redis"
	"github.com/TencentAd/attribution/attribution/pkg/storage/metadata/sqlite"
)

const (
	StoreTypeSqlite StoreType = "SQLITE"
	StoreTypeMysql  StoreType = "MYSQL"
	StoreTypeRedis  StoreType = "REDIS"
)

var (
	instance      Store
	once          = &sync.Once{}
	DefaultOption = &Option{StoreTypeSqlite, map[string]interface{}{"dsn": "sqlite.db"}}
)

type Store interface {
	Get(string) (string, error)
	Set(string, string) error
}

type StoreType string

type Option struct {
	Type   StoreType              `json:"type"`
	Config map[string]interface{} `json:"config"`
}

func create(option *Option) (Store, error) {
	switch option.Type {
	case StoreTypeSqlite:
		return sqlite.New(option.Config)
	case StoreTypeRedis:
		return redis.New(option.Config)
	default:
		return nil, fmt.Errorf("not support metadata store type")
	}
}

func GetStore(storeType *string, storeOptionStr *string) Store {
	storeOption := make(map[string]interface{})

	err := json.Unmarshal([]byte(*storeOptionStr), &storeOption)
	var option *Option = nil

	if err != nil {
		option = DefaultOption
	} else {
		option = &Option{StoreType(*storeType), storeOption}
	}

	once.Do(func() {
		store, err := create(option)
		if err != nil {
			log.Panic(err)
		}

		instance = store
	})
	return instance
}
