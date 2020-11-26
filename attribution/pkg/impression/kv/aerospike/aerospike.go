package aerospike

import (
	"net"
	"strconv"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/impression/kv/opt"
	"github.com/aerospike/aerospike-client-go"
)

var (
	Namespace = "attribution"
	binName   = "impression_time"
)

type Aerospike struct {
	client *aerospike.Client
	option *opt.Option
}

func (a *Aerospike) Get(key string) (string, error) {
	k, _ := aerospike.NewKey(Namespace, opt.Prefix, key)
	r, err := a.client.Get(nil, k)
	if r == nil || err != nil {
		return "", err
	}

	return r.Bins[binName].(string), nil
}

func (a *Aerospike) Set(key string, value string) error {
	k, _ := aerospike.NewKey(Namespace, opt.Prefix, key)
	wp := aerospike.NewWritePolicy(0, uint32(a.option.Expiration/time.Second))
	return a.client.Put(wp, k, aerospike.BinMap{
        binName: value,
    })
}

func New(option *opt.Option) (*Aerospike, error) {
	ip, port1, _ := net.SplitHostPort(option.Address)
	port, _ := strconv.Atoi(port1)
	client, err := aerospike.NewClient(ip, port)
	if err != nil {
		return nil, err
	}

	return &Aerospike{
		client: client,
		option: option,
	}, nil
}
