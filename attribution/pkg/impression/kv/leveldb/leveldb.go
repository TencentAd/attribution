package leveldb

import (
    "time"

    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/opt"
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/filter"
    dbopt "github.com/syndtr/goleveldb/leveldb/opt"
    "github.com/syndtr/goleveldb/leveldb/util"
)

type LevelDb struct {
    db *leveldb.DB
    option *opt.Option
}

func (s *LevelDb) Has(key string) (bool, error) {
    value, err := s.db.Get([]byte(opt.Prefix + key), nil)
    if err != nil {
        return false, err
    }

    t := getBytesTime(value)
    return time.Since(t) <= s.option.Expiration, nil
}

func (s *LevelDb) Set(key string) error {
    return s.db.Put([]byte(opt.Prefix + key), getTimeBytes(time.Now()), nil)
}

func (s *LevelDb) cronjob() {
    iter := s.db.NewIterator(util.BytesPrefix([]byte(opt.Prefix)), nil)
    for iter.Next() {
        t := getBytesTime(iter.Value())
        if time.Since(t) > s.option.Expiration {
            key := iter.Key()
            _ = s.db.Delete(key, nil)
        }
    }
}

func getTimeBytes(t time.Time) []byte {
    return []byte(t.Format(time.RFC3339))
}

func getBytesTime(bytes []byte) time.Time {
    t, _ := time.Parse(time.RFC3339, string(bytes))
    return t
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