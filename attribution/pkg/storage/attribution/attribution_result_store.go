/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/11/20, 4:30 PM
 */

package attribution

import (
	"flag"
	"github.com/TencentAd/attribution/attribution/pkg/storage/attribution/kafka"
	"strings"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/storage/attribution/http"
	"github.com/TencentAd/attribution/attribution/pkg/storage/attribution/native"
)

var (
	attributionStoresNames = flag.String("attribution_result_storage", "stdout", "multiple storage split by comma")

	attributionStoreFactory = factory.NewFactory("attribution_result")
)

type Storage interface {
	//Store(conv *conv.ConversionLog) error
	Store(message interface{}) error
}

func init() {
	attributionStoreFactory.Register("ams", http.NewAmsAttributionForward)
	attributionStoreFactory.Register("stdout", native.NewStdoutAttributionStore)
	attributionStoreFactory.Register("kafka_click", kafka.NewAmsKafkaClickStore)
	attributionStoreFactory.Register("kafka_conversion", kafka.NewAmsKafkaConversionStore)
	attributionStoreFactory.Register("kafka_attribution", kafka.NewAmsKafkaAttributionStore)
}

func CreateAttributionStore() ([]Storage, error) {
	names := strings.Split(*attributionStoresNames, ",")
	ret := make([]Storage, 0, len(names))
	for _, name := range names {
		if s, err := createSingleAttributionStore(name); err != nil {
			return nil, err
		} else {
			ret = append(ret, s)
		}
	}
	return ret, nil
}

func createSingleAttributionStore(name string) (Storage, error) {
	if s, err := attributionStoreFactory.Create(name); err != nil {
		return nil, err
	} else {
		return s.(Storage), nil
	}
}
