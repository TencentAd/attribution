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

func (s *StdoutAttributionStore) Store(message interface{}) error {
	conversionLog := message.(*conv.ConversionLog)
	var convContent []byte
	var clickContent []byte
	var err error
	convContent, err = json.Marshal(conversionLog)
	if err != nil {
		return err
	}

	if conversionLog.MatchClick != nil {
		clickContent, err = json.Marshal(conversionLog.MatchClick)
		if err != nil {
			return err
		}
	}

	fmt.Printf("conversionLog: %s\nclick:%s\n", string(convContent), string(clickContent))
	return nil
}
