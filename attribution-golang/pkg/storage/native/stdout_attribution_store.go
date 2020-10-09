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

	"attribution/proto/click"
	"attribution/proto/conv"
)

type StdoutAttributionStore struct {
}

func NewStdoutAttributionStore() *StdoutAttributionStore {
	return &StdoutAttributionStore{}
}

func (s *StdoutAttributionStore) Store(conv *conv.ConversionLog, click *click.ClickLog) error {
	var convContent []byte
	var clickContent []byte
	var err error
	convContent, err = json.Marshal(conv)
	if err != nil {
		return err
	}

	if click != nil {
		clickContent, err = json.Marshal(click)
		if err != nil {
			return err
		}
	}

	fmt.Printf("conv: %s\nclick:%s\n", string(convContent), string(clickContent))
	return nil
}
