/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 2:37 PM
 */

package action

import (
	"github.com/TencentAd/attribution/attribution/pkg/association"
	"github.com/TencentAd/attribution/attribution/pkg/association/validation"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/conv/data"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"

	"github.com/golang/glog"
)

type ClickAssocAction struct {
	index clickindex.ClickIndex
	assoc *association.ClickAssociation
}

func NewClickAssocAction(index clickindex.ClickIndex) *ClickAssocAction {
	return &ClickAssocAction{
		index: index,
		assoc: association.NewClickAssociation().
			WithClickIndex(index).
			WithValidation(&validation.DefaultClickLogValidation{},
			),
	}
}

func (action *ClickAssocAction) Run(i interface{}) {
	c := i.(*data.ConvContext)
	if err := action.runInternal(c); err != nil {
		c.StopWithError(err)
		glog.Errorf("err: %v", err)
		return
	}
}

func (action *ClickAssocAction) runInternal(c *data.ConvContext) error {
	return action.assoc.Association(c.AssocContext)
}
