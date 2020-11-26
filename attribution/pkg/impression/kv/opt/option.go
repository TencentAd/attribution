package opt

import (
	"time"
)

var (
	// DefaultExpiration : default kv storage expiration
	DefaultExpiration = 7 * 24 * time.Hour
	// Prefix : default key's prefix for kv storage
	Prefix = "impression::"
)

// Option : option to kv storage configuration
type Option struct {
	Address    string        `json:"address"`
	Password   string        `json:"password"`
	Expiration time.Duration `json:"expiration"`
}
