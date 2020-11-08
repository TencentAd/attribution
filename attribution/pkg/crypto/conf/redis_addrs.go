package conf

import (
	"fmt"
	"strings"
)

type redisAddrs []string

type Value interface {
	String() string
	Set(string) error
}

func (r *redisAddrs) String() string {
	return fmt.Sprint(*r)
}

func (r *redisAddrs) Set(value string) error {
	split := strings.Split(value, ",")
	*r = split
	return nil
}
