/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 4:26 PM
 */

package native

import (
	"github.com/TencentAd/attribution/attribution/pkg/common/define"
	"github.com/TencentAd/attribution/attribution/proto/click"
	"github.com/TencentAd/attribution/attribution/proto/user"

	"github.com/golang/glog"
	cmap "github.com/orcaman/concurrent-map"
)

type ClickIndexNative struct {
	data []cmap.ConcurrentMap
}

func NewClickIndexNative() interface{} {
	length := len(user.IdType_value)

	kv := &ClickIndexNative{
		data: make([]cmap.ConcurrentMap, len(user.IdType_value)),
	}

	for i := 0; i < length; i++ {
		kv.data[i] = cmap.New()
	}
	return kv
}

func (kv *ClickIndexNative) Set(t user.IdType, key string, value *click.ClickLog) error {
	if glog.V(define.VLogLevel) {
		glog.V(define.VLogLevel).Infof("set [%d] [%s] [%v]", t, key, value)
	}
	kv.data[t].Set(key, value)
	return nil
}

func (kv *ClickIndexNative) Get(t user.IdType, key string) (*click.ClickLog, error) {
	if glog.V(define.VLogLevel) {
		glog.V(define.VLogLevel).Infof("get [%d] [%s]", t, key)
	}
	if v, ok := kv.data[t].Get(key); ok {
		return v.(*click.ClickLog), nil
	} else {
		return nil, nil
	}
}

func (kv *ClickIndexNative) Remove(t user.IdType, key string) error {
	if glog.V(define.VLogLevel) {
		glog.V(define.VLogLevel).Infof("remove [%d] [%s]", t, key)
	}
	kv.data[t].Remove(key)
	return nil
}