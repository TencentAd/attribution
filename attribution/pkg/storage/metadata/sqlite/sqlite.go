package sqlite

import (
	"github.com/TencentAd/attribution/attribution/pkg/storage/metadata/opt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	db *gorm.DB
}

func (s *Sqlite) Get(key string) (string, error) {
	if v, ok := s.db.Get(opt.Prefix + key); ok {
		return v.(string), nil
	} else {
		return "", nil
	}
}

func (s *Sqlite) Set(key string, value string) error {
	return s.db.Set(opt.Prefix + key, value).Error
}

func New(config map[string]interface{}) (*Sqlite, error) {
	dsn := config["dsn"].(string)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Sqlite{db: db}, nil
}
