/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 10:20 AM
 */

package counter

import (
	"flag"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/encrypt/counter/redis"
)

var (
	counterName    = flag.String("counter_name", "redis", "support [redis]")
	counterFactory = factory.NewFactory("counter")
)

func init() {
	counterFactory.Register("redis", redis.NewCounter)
}

// MinuteFreqRedis 用于对请求次数进行记录
type Counter interface {
	Inc(resource string) (int64, error)
	Get(resource string) (int64, error)
}

func CreateCounter() (Counter, error) {
	if c, err := counterFactory.Create(*counterName); err == nil {
		return nil, err
	} else {
		return c.(Counter), nil
	}
}
