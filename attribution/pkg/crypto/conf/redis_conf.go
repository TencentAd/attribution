package conf

import (
	"errors"
	"flag"
	"github.com/golang/glog"
)

type RedisMode int

var (
	ErrNotSupportedRedisMode = errors.New("un supported redis mode")
)

const (
	RedisSingleNode RedisMode = 0
	RedisCluster    RedisMode = 1
)

// 这个结构体用于保存所有redis连接的相关配置
type RedisConf struct {
	// 这里变量名首字母要大写
	NodeAddresses []string
	RedisMode     RedisMode
}

func NewRedisConf(mode RedisMode, addresses ...string) (*RedisConf, error) {
	if mode != RedisCluster && mode != RedisSingleNode {
		glog.Errorf("unsupported redis mode")
		return nil, ErrNotSupportedRedisMode
	}
	return &RedisConf{
		NodeAddresses: addresses,
		RedisMode:     mode,
	}, nil
}

// 刚函数允许用户使用flag来传入redis参数
// 默认使用的是redis的单点模式
func (r *RedisConf) InitConf() {
	var nodeAddresses addrs
	flag.Var(&nodeAddresses, "redis_addresses", "")
	mode := flag.Int("redis_mode", 0, "")
	flag.Parse()
	r.NodeAddresses = nodeAddresses
	r.RedisMode = RedisMode(*mode)
}
