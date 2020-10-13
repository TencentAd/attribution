/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/10/20, 11:18 AM
 */

package jsonline

import (
	"encoding/json"

	"github.com/TencentAd/attribution/attribution/proto/conv"
)

// 按照json格式解析http的body
type ConvParser struct{}

func NewConvParser() *ConvParser {
	return &ConvParser{}
}

func (p *ConvParser) Parse(input interface{}) ([]*conv.ConversionLog, error) {
	line := input.(string)
	convLog := new(conv.ConversionLog)
	err := json.Unmarshal([]byte(line), convLog)
	if err != nil {
		return nil, err
	}
	return []*conv.ConversionLog{convLog}, nil
}
