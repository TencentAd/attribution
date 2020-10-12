/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 4:29 PM
 */
package data

import (
	"attribution/pkg/association"
	"attribution/pkg/data"
	"attribution/proto/conv"
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
