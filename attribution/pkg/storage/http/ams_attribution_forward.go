/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/12/20, 3:55 PM
 */

package http

import (
	"net/http"

	"attribution/pkg/storage"
	"attribution/proto/conv"
)

func init()  {
	storage.AttributionStoreFactory.Register("ams", NewAmsAttributionForward)
}

type AmsAttributionForward struct {
	client *http.Client
}

func NewAmsAttributionForward() interface{} {
	return &AmsAttributionForward{
		client: &http.Client{},
	}
}

func (f *AmsAttributionForward) Store(convLog *conv.ConversionLog) error {

}

