/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 3:49 PM
 */

package validation

import (
	"github.com/TencentAd/attribution/attribution/proto/click"
	"github.com/TencentAd/attribution/attribution/proto/conv"
)

type ClickLogValidation interface {
	Check(*conv.ConversionLog, *click.ClickLog) bool
}

// 默认点击日志过滤器
//  1. 点击事件时间必须小于转化事件时间
//  2. 转化事件必须在点击发生的一定的时间范围内
//
// 根据广告系统的不同以及不同的场景，需要支持不同的策略
type DefaultClickLogValidation struct {
}

const ConvEventDelayThreshold = 86400 * 7

func (f *DefaultClickLogValidation) Check(convLog *conv.ConversionLog, clickLog *click.ClickLog) bool {
	if convLog.AppId != clickLog.AppId {
		return false
	}

	if clickLog.ClickTime >= convLog.EventTime {
		return false
	}

	if convLog.EventTime-clickLog.ClickTime > ConvEventDelayThreshold {
		return false
	}

	return true
}
