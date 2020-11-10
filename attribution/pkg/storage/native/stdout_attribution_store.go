/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  yuanweishi@tencent.com
 * Last Modify: 9/16/20, 4:26 PM
 */

package native

import (
	"encoding/json"
	"fmt"

	"github.com/TencentAd/attribution/attribution/proto/conv"
)

type StdoutAttributionStore struct {
}

func NewStdoutAttributionStore() interface{} {
	return &StdoutAttributionStore{}
}

func (s *StdoutAttributionStore) Store(conv *conv.ConversionLog) error {
	var convContent []byte
	var clickContent []byte
	var err error
	convContent, err = json.Marshal(conv)
	if err != nil {
		return err
	}

	if conv.MatchClick != nil {
		clickContent, err = json.Marshal(conv.MatchClick)
		if err != nil {
			return err
		}
	}

	fmt.Printf("conv: %s\nclick:%s\n", string(convContent), string(clickContent))
	return nil
}
