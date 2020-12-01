/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 4:41 PM
 */

package freq

import (
	"flag"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/encrypt/freq/info"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/encrypt/freq/redis"
)

var (
	minuteFreqName    = flag.String("minute_freq_storage_name", "redis", "support [redis]")
	minuteFreqFactory = factory.NewFactory("minute_freq")
)

func init()  {
	minuteFreqFactory.Register("redis", redis.NewMinuteFreqRedis)
}

type MinuteFreqStorage interface {
	Get(resource string) (*info.MinuteFreqInfo, error)
	Set(resource string, info *info.MinuteFreqInfo) error
}

func CreateMinuteFreq() (MinuteFreqStorage, error) {
	if c, err := minuteFreqFactory.Create(*minuteFreqName); err == nil {
		return nil, err
	} else {
		return c.(MinuteFreqStorage), nil
	}
}
