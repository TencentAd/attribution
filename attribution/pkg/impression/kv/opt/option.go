package opt

import (
    "time"
)

var (
    DefaultExpiration = 7*24*time.Hour
    Prefix = "impression::"
)

type Option struct {
    Address  string          `json:"address"`
    Password string          `json:"password"`
    Expiration time.Duration `json:"expiration"`
}
