/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:30 AM
 */

package click

import (
	"encoding/json"
	"errors"
	"flag"
	"net/http"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/workflow"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/click/action"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/click/data"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/click/response"
	"github.com/TencentAd/attribution/attribution/pkg/parser"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"
	"github.com/TencentAd/attribution/attribution/proto/click"

	"github.com/golang/glog"
)

var (
	clickHandleWorkerCount    = flag.Int("click_http_handle_worker_count", 50, "")
	clickHandleQueueSize      = flag.Int("click_http_handle_queue_size", 200, "")
	clickHandleQueueTimeoutMS = flag.Int("click_http_handle_queue_timeout_ms", 10, "")
)

type HttpHandle struct {
	parser     parser.ClickParserInterface
	ClickIndex clickindex.ClickIndex

	jobQueue workflow.JobQueue

	// 定义所有的action，如增加id mapping等完善id信息
	clickLogIndexAction *action.ClickLogIndexAction
}

func NewClickHttpHandle() *HttpHandle {
	jq := workflow.NewDefaultJobQueue(
		&workflow.QueueOption{
			WorkerCount: *clickHandleWorkerCount,
			QueueSize:   *clickHandleQueueSize,
			PushTimeout: time.Duration(*clickHandleQueueTimeoutMS) * time.Millisecond,
		})
	jq.Start()

	return &HttpHandle{
		jobQueue: jq,
	}
}

func (handle *HttpHandle) WithParser(parser parser.ClickParserInterface) *HttpHandle {
	handle.parser = parser
	return handle
}

func (handle *HttpHandle) WithClickIndex(clickIndex clickindex.ClickIndex) *HttpHandle {
	handle.ClickIndex = clickIndex
	return handle
}

func (handle *HttpHandle) Init() error {
	if handle.parser == nil {
		return errors.New("parser should be set")
	}
	if handle.ClickIndex == nil {
		return errors.New("index should be set")
	}

	handle.clickLogIndexAction = action.NewClickLogIndexAction(handle.ClickIndex)

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
	var clickLog *click.ClickLog
	clickLog, err = handle.parser.Parse(r)
	if err != nil {
		return err
	}

	c := data.NewClickContext()
	c.SetClickLog(clickLog)

	handle.run(c)
	err = c.Error
	return err
}

func (handle *HttpHandle) run(c *data.ClickContext) {
	wf := workflow.NewWorkFlow()

	clickLogIndexTask := workflow.NewTaskNode(handle.clickLogIndexAction)
	wf.AddStartNode(clickLogIndexTask)
	wf.ConnectToEnd(clickLogIndexTask)

	wf.StartWithJobQueue(handle.jobQueue, c.Ctx, c)
	wf.WaitDone()
}

func (handle *HttpHandle) serveResponse(w http.ResponseWriter, err error) {
	var resp *response.ClickHttpResponse
	if err != nil {
		resp = &response.ClickHttpResponse{
			Code:    -1,
			Message: err.Error(),
		}
		w.WriteHeader(400)
	} else {
		resp = &response.ClickHttpResponse{
			Message: "success",
		}
	}

	bytes, err := json.Marshal(resp)
	if err == nil {
		w.Write(bytes)
	} else {
		glog.Errorf("failed to marshal response, err: %v", err)
	}
}
