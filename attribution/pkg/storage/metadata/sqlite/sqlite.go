package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	db *gorm.DB
}

type meta struct {
	Key   string `gorm:"primaryKey"`
	Value string
}

func (s *Sqlite) Get(key string) (string, error) {
	m := &meta{Key: key}
	err := s.db.FirstOrCreate(m).Error
	return m.Value, err
}

func (s *Sqlite) Set(key string, value string) error {
	return s.db.Updates(&meta{Key: key, Value: value}).Error
}

func New(config map[string]interface{}) (*Sqlite, error) {
	dsn := config["dsn"].(string)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Sqlite{db: db}, db.AutoMigrate(&meta{})
}
