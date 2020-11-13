/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/12/20, 5:14 PM
 */

package loader

import (
	"flag"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/colinmarc/hdfs"
	"github.com/golang/glog"
)

var (
	hdfsAddress = flag.String("hdfs_address", "", "comma separated address")
	hdfsUser    = flag.String("hdfs_user", "", "")

	defaultClient *hdfs.Client
	once          sync.Once
)

type HdfsHotLoader struct {
	client   *hdfs.Client
	filePath string
	fileLoad fileLoadFunc
	alloc    allocFunc

	detectModify time.Time
	lastModify   time.Time
}

type fileLoadFunc func(reader io.Reader, i interface{}) error

func NewHdfsHotLoader(client *hdfs.Client, filePath string, loader fileLoadFunc, alloc allocFunc) *HdfsHotLoader {
	return &HdfsHotLoader{
		client:   client,
		filePath: filePath,
		fileLoad: loader,
		alloc:    alloc,
	}
}

func NewHdfsHotLoaderWithDefaultClient(filePath string, loader fileLoadFunc, alloc allocFunc) *HdfsHotLoader {
	var err error
	once.Do(
		func() {
			defaultClient, err = hdfs.NewClient(hdfs.ClientOptions{
				Addresses: strings.Split(*hdfsAddress, ","),
				User:      *hdfsUser,
			})
			if err != nil {
				glog.Errorf("failed to create default hdfs client, err: %v", err)
			}
		},
	)
	if err != nil {
		return nil
	}
	return NewHdfsHotLoader(defaultClient, filePath, loader, alloc)
}

func (l *HdfsHotLoader) Alloc() interface{} {
	return l.alloc()
}

func (l *HdfsHotLoader) Reset() {
	l.lastModify = time.Unix(0, 0)
}

func (l *HdfsHotLoader) DetectNewFile() (string, bool) {
	statInfo, err := l.client.Stat(l.filePath)
	if err != nil {
		glog.Errorf("failed to stat[%s], err: %v", l.filePath, err)
		return "", false
	}

	l.detectModify = statInfo.ModTime()
	return l.filePath, l.detectModify.After(l.lastModify)
}

func (l *HdfsHotLoader) Load(filePath string, i interface{}) error {
	reader, err := l.client.Open(filePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	err = l.fileLoad(reader, i)
	if err != nil {
		return err
	}

	l.lastModify = l.detectModify
	return nil
}
