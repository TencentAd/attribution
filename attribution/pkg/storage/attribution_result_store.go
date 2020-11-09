/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  yuanweishi@tencent.com
 * Last Modify: 9/16/20, 4:26 PM
 */

package storage

import (
	"flag"
	"strings"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/storage/http"
	"github.com/TencentAd/attribution/attribution/pkg/storage/native"
	"github.com/TencentAd/attribution/attribution/proto/conv"
)

var (
	attributionStoresNames = flag.String("attribution_result_storage", "stdout", "multiple storage split by comma")

	attributionStoreFactory = factory.NewFactory("attribution_result")
)

type AttributionStore interface {
	Store(conv *conv.ConversionLog) error
}

func init() {
	attributionStoreFactory.Register("ams", http.NewAmsAttributionForward)
	attributionStoreFactory.Register("stdout", native.NewStdoutAttributionStore)
}

func CreateAttributionStore() ([]AttributionStore, error) {
	names := strings.Split(*attributionStoresNames, ",")
	ret := make([]AttributionStore, 0, len(names))
	for _, name := range names {
		if s, err := createSingleAttributionStore(name); err != nil {
			return nil, err
		} else {
			ret = append(ret, s)
		}
	}
	return ret, nil
}

func createSingleAttributionStore(name string) (AttributionStore, error) {
	if s, err := attributionStoreFactory.Create(name); err != nil {
		return nil, err
	} else {
		return s.(AttributionStore), nil
	}
}
