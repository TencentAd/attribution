package conf

import (
	"fmt"
	"strings"
)

type addrs []string

type Value interface {
	String() string
	Set(string) error
}

func (r *addrs) String() string {
	return fmt.Sprint(*r)
}

func (r *addrs) Set(value string) error {
	split := strings.Split(value, ",")
	*r = split
	return nil
}
