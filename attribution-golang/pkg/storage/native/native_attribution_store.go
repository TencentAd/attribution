/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  yuanweishi@tencent.com
 * Last Modify: 9/16/20, 4:26 PM
 */

package native

import (
	"attribution/proto/click"
	"attribution/proto/conv"
	"bufio"
	"encoding/json"
	"flag"
	"os"

	"github.com/golang/glog"
	"attribution/pkg/association"
	"attribution/pkg/common/define"
)

var (
	attribution_store_filename  = flag.String("attribution_store_filename", "", "")
)

type NativeAttributionStore struct {
	asyn_writer *AsynNativeFileWriter
}

func NewNativeAttributionStore() *NativeAttributionStore {
	native_attribution_store := &NativeAttributionStore{
		asyn_writer: NewAsynNativeWriter(),
	}
	return native_attribution_store
}

func (kv *NativeAttributionStore) Store(conv *conv.ConversionLog, click *click.ClickLog) error {
	ass := &association.AssocContext {
		ConvLog: conv,
		SelectClick: click,
	}
	kv.asyn_writer.as_channel <- ass
	return nil
}

type AsynNativeFileWriter struct {
	as_channel chan *association.AssocContext
}

func NewAsynNativeWriter() *AsynNativeFileWriter{
	asyn_native_writer := new(AsynNativeFileWriter)
	asyn_native_writer.as_channel = make(chan *association.AssocContext, 10000)
	asyn_native_writer.Start()
	return asyn_native_writer

}

func (asyn_writer *AsynNativeFileWriter) Start() {
	go func() {
		store_item := 0
		file, err := os.OpenFile(*attribution_store_filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0664)
		if err != nil {
			glog.V(define.VLogLevel).Infof("open file [%s] failed", *attribution_store_filename)
			return
		}
		writer := bufio.NewWriter(file)

		for {
			as := <- asyn_writer.as_channel

			if as.SelectClick == nil {
				continue;
			}

			var convContent []byte
			var clickContent []byte
			var err error
			convContent, err = json.Marshal(as.ConvLog)
			if err != nil {
				glog.V(define.VLogLevel).Infof("Store Marshal conv [%s] failed", err)
				continue;
			}

			clickContent, err = json.Marshal(as.SelectClick)
			if err != nil {
				glog.V(define.VLogLevel).Infof("Store Marshal click [%s] failed", err)
			}

			var ss string = string(convContent)
			ss += "###"
			ss += string(clickContent)
			ss += "\n"
			if _, err1 := writer.WriteString(ss); err1 != nil {
				glog.V(define.VLogLevel).Infof("Fatal error store to disk [%s] failed", err1)
			}

			store_item++
			if store_item > 100 {
				writer.Flush()
				store_item = 0
			}
		}
		file.Close()
	}()
}




