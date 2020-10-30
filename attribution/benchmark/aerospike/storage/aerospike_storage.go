/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/27/20, 7:37 PM
 */

package main

import (
	"flag"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/handler/file/line"
	"github.com/aerospike/aerospike-client-go"
	"github.com/golang/glog"
)

var (
	filePath      = flag.String("file_path", "", "")
	aerospikeHost = flag.String("aerospike_address", "", "")
	aerospikePort = flag.Int("aerospike_port", 3000, "")
)

type StorageKeyHandle struct {
	client   *aerospike.Client
	filename string
}

func NewClickFileHandle(filename string) (*StorageKeyHandle, error) {
	client, err := aerospike.NewClient(*aerospikeHost, *aerospikePort)
	if err != nil {
		return nil, err
	}
	return &StorageKeyHandle{
		filename: filename,
		client:   client,
	}, nil
}

func (p *StorageKeyHandle) Run() error {
	lp := line.NewLineProcess(p.filename, p.processLine, func(line string, err error) {
		glog.Errorf("failed to handle line[%s], err[%v]", line, err)
	}).WithQueueTimeout(time.Second * 10).
		WithParallelism(100)


	if err := lp.Run(); err != nil {
		return err
	}

	lp.WaitDone()
	return nil
}

func (p *StorageKeyHandle) processLine(line string) error {
	key, err := aerospike.NewKey("test", "aerospike", line)
	if err != nil {
		return err
	}
	bins := aerospike.BinMap{"a": ""}
	err = p.client.Put(nil, key, bins)
	if err != nil {
		return err
	}
	return err
}

func main() {
	flag.Parse()
	handle, err := NewClickFileHandle(*filePath)
	if err != nil {
		panic(err)
	}

	if err := handle.Run(); err != nil {
		panic(err)
	}
}
