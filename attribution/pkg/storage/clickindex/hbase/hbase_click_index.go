/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  yuanweishi@tencent.com
 * Last Modify: 9/14/20, 4:26 PM
 */

package hbase

import (
	"flag"

	"github.com/TencentAd/attribution/attribution/pkg/common/define"
	"github.com/TencentAd/attribution/attribution/proto/click"
	"github.com/TencentAd/attribution/attribution/proto/user"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

var (
	hbaseZkAddr    = flag.String("hbase_zk_addr", "", "")
	hbaseZkPath    = flag.String("hbase_zk_path", "", "")
	hbaseTableName = flag.String("hbase_table_name", "", "")
)

const hbaseColumnFamilyName string = "click_family"
const hbaseColumnName string = "click"

type ClickIndexHbase struct {
	hbaseClient *ClientX
	tableName   string
}

func NewClickIndexHbase() interface{} {
	hbaseClient, err := GetInstance("hbase_click", *hbaseZkAddr, *hbaseZkPath)
	if err != nil {
		return nil
	}

	clickIndexBase := &ClickIndexHbase{
		hbaseClient: hbaseClient,
		tableName:   *hbaseTableName,
	}
	return clickIndexBase
}

func (kv *ClickIndexHbase) Set(t user.IdType, key string, value *click.ClickLog) error {
	bytesClickLog, err := proto.Marshal(value)
	if err != nil {
		glog.V(define.VLogLevel).Infof("marshal [%d] [%s] failed", t, key)
		return err
	}

	combineKey := user.IdType_name[int32(t)] + "_" + key
	clickValue := map[string]map[string][]byte{
		hbaseColumnFamilyName: {
			hbaseColumnName: bytesClickLog,
		},
	}

	if rval := kv.hbaseClient.Put(kv.tableName, combineKey, clickValue); rval != nil {
		glog.V(define.VLogLevel).Infof("put [%d] [%s] failed, [%s]", t, key, rval)
		return rval
	}
	return nil
}

func (kv *ClickIndexHbase) Get(t user.IdType, key string) (*click.ClickLog, error) {
	combineKey := user.IdType_name[int32(t)] + "_" + key
	row, err := kv.hbaseClient.Get(kv.tableName, combineKey)
	if err != nil {
		glog.V(define.VLogLevel).Infof("Get value failed [%d] [%s] failed, [%s]", t, key, err)
		return nil, err
	}

	byteValue := row.GetValue(hbaseColumnFamilyName, hbaseColumnName)
	if len(byteValue) < 5 {
		glog.V(define.VLogLevel).Infof("Get value failed [%d] [%s] failed", t, key)
		return nil, nil
	}

	clickLog := &click.ClickLog{}
	if err := proto.Unmarshal(byteValue, clickLog); err != nil {
		glog.V(define.VLogLevel).Infof("unmarshal value failed [%d] [%s] failed", t, key)
		return nil, nil
	}
	return clickLog, nil
}

func (kv *ClickIndexHbase) Remove(t user.IdType, key string) error {
	combineKey := user.IdType_name[int32(t)] + "_" + key
	kv.hbaseClient.Delete(kv.tableName, combineKey, nil)
	return nil
}
