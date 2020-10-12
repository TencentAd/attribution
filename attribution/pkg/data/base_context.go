/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 4:17 PM
 */

package data

import (
	"context"
)

type BaseContext struct {
	Ctx    context.Context
	cancel context.CancelFunc
	Error  error
}

func NewBaseContext() *BaseContext {
	ctx, cancel := context.WithCancel(context.Background())
	return &BaseContext{
		Ctx:    ctx,
		cancel: cancel,
	}
}

func (c *BaseContext) StopWithError(err error) {
	c.Error = err
	c.cancel()
}
