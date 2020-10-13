/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 4:29 PM
 */
package data

import (
	"github.com/TencentAd/attribution/attribution/pkg/association"
	"github.com/TencentAd/attribution/attribution/pkg/data"
	"github.com/TencentAd/attribution/attribution/proto/conv"
)

type ConvContext struct {
	*data.BaseContext
	AssocContext *association.AssocContext
}

func NewConvContext() *ConvContext {
	return &ConvContext{
		BaseContext: data.NewBaseContext(),
	}
}

func (c *ConvContext) SetConvLog(log *conv.ConversionLog) {
	c.AssocContext = association.NewAssocContext(log)
}
