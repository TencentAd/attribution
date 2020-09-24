/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  yuanweishi@tencent.com
 * Last Modify: 9/14/20, 4:26 PM
 */

package hbase

import (
	"attribution/pkg/common/define"
	"attribution/proto/click"
	"attribution/proto/user"
	"flag"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

var (
	hbase_zk_addr   = flag.String("hbase_zk_addr", "", "")
	hbase_zk_path    = flag.String("hbase_zk_path", "", "")
	hbase_table_name     = flag.String("hbase_table_name", "", "")
)

const hbase_column_family_name string = "click_family"
const hbase_column_name string = "click"

type ClickIndexHbase struct {
	hbase_client *ClientX
	table_name string
}

func NewClickIndeHbase() *ClickIndexHbase {
	hbase_client, err := GetInstance("hbase_click", *hbase_zk_addr, *hbase_zk_path)
	if err != nil {
		return nil
	}

	click_index_base := &ClickIndexHbase{
		hbase_client: hbase_client,
		table_name: *hbase_table_name,
	}
	return click_index_base
}

func (kv *ClickIndexHbase) Set(t user.IdType, key string, value *click.ClickLog) error {
	bytes_click_log, err := proto.Marshal(value);
	if err != nil {
		glog.V(define.VLogLevel).Infof("marshal [%d] [%s] failed", t, key)
		return err
	}

	combine_key := user.IdType_name[int32(t)] + "_" + key
	click_value := map[string]map[string][]byte {
		hbase_column_family_name : map[string][]byte {
			hbase_column_name: bytes_click_log,
		},
	}

	if rval := kv.hbase_client.Put(kv.table_name, combine_key, click_value); rval != nil {
		glog.V(define.VLogLevel).Infof("put [%d] [%s] failed, [%s]", t, key, rval)
		return rval
	}
	return nil
}

func (kv *ClickIndexHbase) Get(t user.IdType, key string) (*click.ClickLog, error) {
	combine_key := user.IdType_name[int32(t)] + "_" + key
	row, err := kv.hbase_client.Get(kv.table_name, combine_key)
	if err != nil {
		glog.V(define.VLogLevel).Infof("Get value failed [%d] [%s] failed, [%s]", t, key, err)
		return nil, err
	}

	byte_value := row.GetValue(hbase_column_family_name, hbase_column_name)
	if len(byte_value) < 5 {
		glog.V(define.VLogLevel).Infof("Get value failed [%d] [%s] failed", t, key)
		return nil, nil
	}

	click_log := &click.ClickLog{}
	if err := proto.Unmarshal(byte_value, click_log); err != nil {
		glog.V(define.VLogLevel).Infof("unmarshal value failed [%d] [%s] failed", t, key)
		return nil, nil
	}
	return click_log, nil
}

func (kv *ClickIndexHbase) Remove(t user.IdType, key string) error {
	combine_key := user.IdType_name[int32(t)] + "_" + key
	kv.hbase_client.Delete(kv.table_name, combine_key, nil)
	return nil
}
