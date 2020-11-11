/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:41 AM
 */

package action

import (
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/click/data"
	"github.com/TencentAd/attribution/attribution/pkg/logic"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"

	"github.com/golang/glog"
)

type ClickLogIndexAction struct {
	index clickindex.ClickIndex
}

func NewClickLogIndexAction(index clickindex.ClickIndex) *ClickLogIndexAction {
	return &ClickLogIndexAction{
		index: index,
	}
}

func (action *ClickLogIndexAction) Run(i interface{}) {
	c := i.(*data.ClickContext)

	if err := action.runInternal(c); err != nil {
		c.StopWithError(err)
		glog.Errorf("err: %v", err)
	}
}

func (action *ClickLogIndexAction) runInternal(c *data.ClickContext) error {
	return logic.ProcessClickLog(c.ClickLog, action.index)
}
