/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:30 AM
 */

package conv

import (
	"encoding/json"
	"errors"
	"flag"
	"net/http"
	"sync"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/workflow"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/conv/action"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/conv/data"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/conv/response"
	"github.com/TencentAd/attribution/attribution/pkg/parser"
	"github.com/TencentAd/attribution/attribution/pkg/protocal/parse"
	"github.com/TencentAd/attribution/attribution/pkg/storage/attribution"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"
	"github.com/TencentAd/attribution/attribution/proto/conv"

	"github.com/golang/glog"
)

var (
	convHandleWorkerCount    = flag.Int("conv_http_handle_worker_count", 50, "")
	convHandleQueueSize      = flag.Int("conv_http_handle_queue_size", 200, "")
	convHandleQueueTimeoutMS = flag.Int("conv_http_handle_queue_timeout_ms", 10, "")
)

type HttpHandle struct {
	parser            parser.ConvParserInterface
	ClickIndex        clickindex.ClickIndex
	attributionStores []attribution.Storage
	jobQueue          workflow.JobQueue

	// 定义所有的action，如增加id mapping等完善id信息
	clickAssocAction *action.ClickAssocAction
}

func NewConvHttpHandle() *HttpHandle {
	jq := workflow.NewDefaultJobQueue(
		&workflow.QueueOption{
			WorkerCount: *convHandleWorkerCount,
			QueueSize:   *convHandleQueueSize,
			PushTimeout: time.Duration(*convHandleQueueTimeoutMS) * time.Millisecond,
		})
	jq.Start()

	return &HttpHandle{
		jobQueue: jq,
	}
}

func (handle *HttpHandle) WithParser(parser parser.ConvParserInterface) *HttpHandle {
	handle.parser = parser
	return handle
}

func (handle *HttpHandle) WithClickIndex(clickIndex clickindex.ClickIndex) *HttpHandle {
	handle.ClickIndex = clickIndex
	return handle
}

func (handle *HttpHandle) WithAttributionStore(attributionStores []attribution.Storage) *HttpHandle {
	handle.attributionStores = attributionStores
	return handle
}

func (handle *HttpHandle) Init() error {
	if handle.parser == nil {
		return errors.New("parser should be set")
	}
	if handle.ClickIndex == nil {
		return errors.New("index should be set")
	}

	handle.clickAssocAction = action.NewClickAssocAction(handle.ClickIndex)

	return nil
}

func (handle *HttpHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO 监控
	if err := handle.doServeHTTP(w, r); err != nil {
		glog.Errorf("failed to serve, err: %v", err)
	}
}

func (handle *HttpHandle) doServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var err error
	defer func() {
		handle.serveResponse(w, err)
	}()
	var pr *parse.ConvParseResult
	pr, err = handle.parser.Parse(r)
	if err != nil {
		return err
	}

	// 并行处理所有的转化
	var wg sync.WaitGroup
	wg.Add(len(pr.ConvLogs))
	for _, convLog := range pr.ConvLogs {
		go func(convLog *conv.ConversionLog) {
			defer wg.Add(1)

			c := data.NewConvContext()
			c.SetConvLog(convLog)
			handle.run(c)
			for _, s := range handle.attributionStores {
				if err = s.Store(c.AssocContext.ConvLog); err != nil {
					glog.Errorf("failed to store, err: %v", err)
				}
			}

			if c.Error != nil {
				err = c.Error
			}
		}(convLog)
	}

	return err
}

func (handle *HttpHandle) run(c *data.ConvContext) {
	wf := workflow.NewWorkFlow()

	clickLogIndexTask := workflow.NewTaskNode(handle.clickAssocAction)
	wf.AddStartNode(clickLogIndexTask)
	wf.ConnectToEnd(clickLogIndexTask)

	wf.StartWithJobQueue(handle.jobQueue, c.Ctx, c)
	wf.WaitDone()
}

func (handle *HttpHandle) serveResponse(w http.ResponseWriter, err error) {
	var resp *response.ConvHttpResponse
	if err != nil {
		resp = &response.ConvHttpResponse{
			Code:    -1,
			Message: err.Error(),
		}
	} else {
		resp = &response.ConvHttpResponse{
			Code:    0,
			Message: "success",
		}
	}

	bytes, err := json.Marshal(resp)
	if err == nil {
		_, _ = w.Write(bytes)
	} else {
		glog.Errorf("failed to marshal response, err: %v", err)
	}
}
