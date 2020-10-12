/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 2:37 PM
 */

package action

import (
	"attribution/pkg/association"
	"attribution/pkg/association/validation"
	"attribution/pkg/handler/http/conv/data"
	"attribution/pkg/storage"

	"github.com/golang/glog"
)

type ClickAssocAction struct {
	index storage.ClickIndex
	assoc *association.ClickAssociation
}

func NewClickAssocAction(index storage.ClickIndex) *ClickAssocAction {
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
