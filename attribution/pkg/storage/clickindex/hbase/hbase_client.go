/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  yuanweishi@tencent.com
 * Last Modify: 9/14/20, 4:26 PM
 */

package hbase

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

var topCtx = context.Background()

type Scanner struct {
	S       hrpc.Scanner
	Cancel  context.CancelFunc
	m       sync.Mutex
}

//每次Next返回下一行
//当所有行已经遍历完毕时，err会返回io.EOF
//当出现其他错误时，err会返回详细错误信息，并且在下一次调用Next时返回io.EOF
func (o *Scanner) Next() (r *Row, err error) {
	//startTime := time.Now()
	o.m.Lock()
	defer m.Unlock()
	r = new(Row)
	r.R, err = o.S.Next()
	if err == io.EOF {
		o.Close()
	}
	return
}

func (o *Scanner) Close() error {
	o.Cancel()
	return o.S.Close()
}

type ClientX struct {
	C       gohbase.Client
	Ctx     context.Context
	Cancel  context.CancelFunc
	Timeout time.Duration

}

//根据config中的配置生成一个client
//address为zk地址，格式为 {ip}:{port},{ip}:{port},...  ，当没有写port时，zk默认port为2181
//root为hbase根节点位置，zookeeper.znode.parent
//timeout为部分交互操作的超时时间，其中Get，Put，Delete均使用该超时时间，Scan和ScanRange没有超时时间
//mod_id和int_id分别为模调的模块id和接口id
func newClient(name string, quorum string, zkparent string) (c *ClientX, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("%v", e)
		}
	}()

	timeout := 60* time.Second

	ctx, cancel := context.WithCancel(topCtx)

	c = &ClientX{
		gohbase.NewClient(quorum, gohbase.ZookeeperRoot(zkparent),
			gohbase.RegionLookupTimeout(timeout),
			gohbase.RegionReadTimeout(timeout),
			gohbase.ZookeeperTimeout(timeout)),
		ctx,
		cancel,
		timeout,
	}
	return
}

func (o *ClientX) Close() {
	o.Cancel()
	o.C.Close()
}

//指定key查询某一行的信息
//options详见https://godoc.org/github.com/tsuna/gohbase/hrpc
func (o *ClientX) Get(table, key string, options ...func(hrpc.Call) error) (r *Row, err error) {
	//startTime := time.Now()
	var (
		req         *hrpc.Get
		ctx, cancel = context.WithTimeout(o.Ctx, o.Timeout)
		//ctx = context.Background()
	)
	defer cancel()
	if req, err = hrpc.NewGetStr(ctx, table, key, options...); err != nil {
		glog.Errorf("创建请求失败，err = %v\n", err)
		return
	}

	var rsp *hrpc.Result
	if rsp, err = o.C.Get(req); err != nil {
		glog.Errorf("HBASE-GET请求返回失败，err = %v", err)
		return
	}

	r = new(Row)
	r.R = rsp
	return
}

//指定key更新某一行的信息
//value的结构为 {ColumnFamily: {Qualifier: Value} }
//options详见https://godoc.org/github.com/tsuna/gohbase/hrpc
func (o *ClientX) Put(table, key string, value map[string]map[string][]byte, options ...func(hrpc.Call) error) (err error) {
	//startTime := time.Now()
	var (
		req         *hrpc.Mutate
		ctx, cancel = context.WithTimeout(o.Ctx, o.Timeout)
	)
	defer cancel()
	if req, err = hrpc.NewPutStr(ctx, table, key, value, options...); err != nil {
		glog.Errorf("创建请求失败，err = %v\n", err)
		return
	}

	if _, err = o.C.Put(req); err != nil {
		glog.Errorf("HBASE-PUT请求返回失败，err = %v", err)
		return
	}
	return
}

// Check符合条件的指定行是否存在，是则执行Put，否则退出
// 通过table，key，family，qualifier，checkValue定位到需要更新的行
// putValue为执行更新的内容，格式见上文Put中的描述
// 当找到指定行并更新成功时，返回ret = true，err = nil
// 找不到指定行时，ret = false
func (o *ClientX) CheckAndPut(table, key, family, qualifier string, checkValue []byte,
	putValue map[string]map[string][]byte, putOptions ...func(hrpc.Call) error) (ret bool, err error) {
	//startTime := time.Now()
	var (
		req         *hrpc.Mutate
		ctx, cancel = context.WithTimeout(o.Ctx, o.Timeout)
	)
	defer cancel()
	if req, err = hrpc.NewPutStr(ctx, table, key, putValue, putOptions...); err != nil {
		glog.Errorf("创建请求失败，err = %v\n", err)
		return
	}

	if ret, err = o.C.CheckAndPut(req, family, qualifier, checkValue); err != nil {
		glog.Errorf("HBASE-CHECKPUT请求返回失败，err = %v", err)
		return
	}
	return
}

// 指定表名和RowKey，对该行的数据进行删除
// 支持整行删除，部分列族删除，部分列删除，以及对以上各情况的全版本删除和部分版本删除
// 整行删除，value填nil
// 部分列族删除，value填 {"cf": nil} ，将删除名为cf的列族
// 部分列删除，value填 {"cf": {"q": nil} } ，将删除列族cf下名为q的列
// 指定版本删除通过options控制
// 传入hrpc.Timestamp(ts)将删除ts之前的所有版本
// 传入hrpc.DeleteOneVersion()将删除最新的版本
// 具体细节可以查看https://godoc.org/github.com/tsuna/gohbase/hrpc#NewDel
func (o *ClientX) Delete(table, key string, value map[string]map[string][]byte, options ...func(hrpc.Call) error) (err error) {
	//startTime := time.Now()
	var (
		req         *hrpc.Mutate
		ctx, cancel = context.WithTimeout(o.Ctx, o.Timeout)
	)
	defer cancel()
	if req, err = hrpc.NewDelStr(ctx, table, key, value, options...); err != nil {
		glog.Errorf("创建请求失败，err = %v\n", err)
		return
	}

	if _, err = o.C.Delete(req); err != nil {
		glog.Errorf("HBASE-DELETE请求返回失败，err = %v", err)
		return
	}
	return
}

// 全表遍历
// options详见https://godoc.org/github.com/tsuna/gohbase/hrpc
func (o *ClientX) Scan(table string, options ...func(hrpc.Call) error) (s *Scanner, err error) {
	//startTime := time.Now()
	var (
		req         *hrpc.Scan
		ctx, cancel = context.WithCancel(o.Ctx)
	)
	if req, err = hrpc.NewScanStr(ctx, table, options...); err != nil {
		glog.Errorf("创建请求失败，err = %v\n", err)
		return
	}

	s = new(Scanner)
	s.S = o.C.Scan(req)
	s.Cancel = cancel
	return
}

// 通过startRow和endRow指定范围进行遍历
// startRow和endRow可以是前缀匹配
// 未指定时即为从头开始遍历或者一直遍历到结束
// 区间为[startRow, endRow)，即遍历时会包含startRow，但是不会包含endRow
// options详见https://godoc.org/github.com/tsuna/gohbase/hrpc
func (o *ClientX) ScanRange(table, startRow, endRow string, options ...func(hrpc.Call) error) (s *Scanner, err error) {
	//startTime := time.Now()
	var (
		req         *hrpc.Scan
		ctx, cancel = context.WithCancel(o.Ctx)
	)
	if req, err = hrpc.NewScanRangeStr(ctx, table, startRow, endRow, options...); err != nil {
		glog.Errorf("创建请求失败，err = %v\n", err)
		return
	}

	s = new(Scanner)
	s.S = o.C.Scan(req)
	s.Cancel = cancel
	return
}

var (
	clientMap = make(map[string]*ClientX)
	m         sync.Mutex
	initMap   sync.Map
)

func getOnce(name string) *sync.Once {
	if v, ok := initMap.Load(name); ok {
		return v.(*sync.Once)
	} else {
		once := new(sync.Once)
		initMap.Store(name, once)
		return once
	}
}

func GetInstance(name string, quorum string, zkparent string) (c *ClientX, err error) {
	getOnce(name).Do(func() {
		m.Lock()
		defer m.Unlock()
		if c, err = newClient(name, quorum, zkparent); err == nil {
			clientMap[name] = c
		}
	})
	if err != nil {
		return
	}
	c = clientMap[name]
	return
}

func NewInstance(name string, quorum string, zkparent string) (c *ClientX) {
	var err error
	if c, err = GetInstance(name, quorum, zkparent); err != nil {
		panic(err)
	}
	return
}
