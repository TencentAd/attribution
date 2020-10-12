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
	"fmt"
	"strings"

	"attribution/pkg/common/factory"
	"attribution/proto/conv"
)

var (
	attributionStoresNames = flag.String("attribution_result_storage", "stdout", "multiple storage split by comma")

	AttributionStoreFactory = factory.NewFactory()
)

type AttributionStore interface {
	Store(conv *conv.ConversionLog) error
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
	s := AttributionStoreFactory.Create(name)
	if s == nil {
		return nil, fmt.Errorf("attribution store[%s] not register", name)
	}
	return s.(AttributionStore), nil
}