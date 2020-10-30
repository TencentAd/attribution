package aerospike

import (
    "net"
    "strconv"
    "time"

    "github.com/TencentAd/attribution/attribution/pkg/impression/kv"
    "github.com/aerospike/aerospike-client-go"
)

var (
    Namespace = "attribution"
)

type Aerospike struct {
    client *aerospike.Client
    option *kv.Option
}

func (a *Aerospike) Has(key string) (bool, error) {
    k, _ := aerospike.NewKey(Namespace, kv.Prefix, key)
    r, err := a.client.Get(nil, k)
    if r == nil || err != err {
        return false, err
    }
    return true, nil
}

func (a *Aerospike) Set(key string) error {
    k, _ := aerospike.NewKey(Namespace, kv.Prefix, key)
    aerospike.NewWritePolicy(0, uint32(a.option.Expiration / time.Second))
    return a.client.Put(nil, k, nil)
}

func New(option *kv.Option) (*Aerospike, error) {
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