/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 4:27 PM
 */

package data

import (
	"github.com/TencentAd/attribution/attribution/pkg/data"
	"github.com/TencentAd/attribution/attribution/proto/click"
)

type ClickContext struct {
	*data.BaseContext
	ClickLog *click.ClickLog
}

func NewClickContext() *ClickContext {
	return &ClickContext{
		BaseContext: data.NewBaseContext(),
	}
}

func (c *ClickContext) SetClickLog(log *click.ClickLog) {
	c.ClickLog = log
}
