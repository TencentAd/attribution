package leveldb

import (
	"encoding/json"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/impression/kv/opt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	dbopt "github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type item struct {
	Time  time.Time
	Value string
}

func item2bytes(v *item) []byte {
	b, _ := json.Marshal(v)
	return b
}

func bytes2item(bytes []byte) *item {
	v := &item{}
	_ = json.Unmarshal(bytes, v)
	return v
}

type LevelDb struct {
	db     *leveldb.DB
	option *opt.Option
}

func (s *LevelDb) Get(key string) (string, error) {
	value, err := s.db.Get([]byte(opt.Prefix+key), nil)
	if err != nil {
		return "", err
	}

	v := bytes2item(value)
	return v.Value, nil
}

func (s *LevelDb) Set(key string, value string) error {
	v := &item{Time: time.Now(), Value: value}
	return s.db.Put([]byte(opt.Prefix+key), item2bytes(v), nil)
}

func (s *LevelDb) cronjob() {
	iter := s.db.NewIterator(util.BytesPrefix([]byte(opt.Prefix)), nil)
	for iter.Next() {
		v := bytes2item(iter.Value())
		if time.Since(v.Time) > s.option.Expiration {
			key := iter.Key()
			_ = s.db.Delete(key, nil)
		}
	}
}

func New(option *opt.Option) (*LevelDb, error) {
	db, err := leveldb.OpenFile(option.Address, &dbopt.Options{
		Filter: filter.NewBloomFilter(32),
	})

	if err != nil {
		return nil, err
	}

	return &LevelDb{
		db:     db,
		option: option,
	}, nil
}
