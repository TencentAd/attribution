package ia

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/workflow"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/action"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/data"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/ia/metrics"
	"github.com/TencentAd/attribution/attribution/pkg/impression/kv"
	"github.com/TencentAd/attribution/attribution/pkg/oauth"
	"github.com/TencentAd/attribution/attribution/pkg/parser"
	"github.com/TencentAd/attribution/attribution/pkg/protocal/parse"
	"github.com/golang/glog"
)

var (
	ErrEmptyAction = errors.New("empty action")
)

type ImpAttributionHandle struct {
	convParser parser.ConvParserInterface

	jobQueue workflow.JobQueue

	encryptAction      *action.EncryptAction
	encryptTwiceAction *action.EncryptTwiceAction
	firstDecryptAction *action.FirstDecryptAction
	attributionAction  *action.AttributionAction
	finalDecryptAction *action.FinalDecryptAction
	sendAction         *action.SendAction
}

func NewImpAttributionHandle() *ImpAttributionHandle {
	return &ImpAttributionHandle{
		encryptAction:      action.NewEncryptAction(),
		encryptTwiceAction: action.NewEncryptTwiceAction(),
		firstDecryptAction: action.NewFirstDecryptAction(),
		attributionAction:  action.NewAttributionAction(),
		finalDecryptAction: action.NewFinalDecryptAction(),
		sendAction:         action.NewSendAction(),
	}
}

func (handle *ImpAttributionHandle) WithConvParser(convParser parser.ConvParserInterface) *ImpAttributionHandle {
	handle.convParser = convParser
	return handle
}

func (handle *ImpAttributionHandle) WithImpStorage(impStorage kv.KV) *ImpAttributionHandle {
	handle.attributionAction.WithImpStorage(impStorage)
	return handle
}

func (handle *ImpAttributionHandle) WithJobQueue(queue workflow.JobQueue) *ImpAttributionHandle {
	handle.jobQueue = queue
	return handle
}

func (handle *ImpAttributionHandle) WithToken(token *oauth.Token) *ImpAttributionHandle {
	handle.sendAction.WithToken(token)
	return handle
}


func (handle *ImpAttributionHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var err error
	defer func() {
		metrics.CollectHandleMetrics(startTime, err)
	}()

	resp := SuccessResponse
	if err = handle.process(r); err != nil {
		glog.Errorf("failed to process request, err: %v", err)
		resp = CreateErrResponse(err)
	}
	handle.writeResponse(w, resp)
}

func (handle *ImpAttributionHandle) process(r *http.Request) error {
	var err error
	var pr *parse.ConvParseResult
	pr, err = handle.convParser.Parse(r)
	if err != nil {
		return err
	}
	if len(pr.ConvLogs) == 0 {
		return ErrEmptyAction
	}

	var c *data.ImpAttributionContext
	c, err = data.NewImpAttributionContext(pr)
	if err != nil {
		return err
	}

	wf := handle.buildWorkflow()
	wf.StartWithJobQueue(handle.jobQueue, c.Ctx, c)
	wf.WaitDone()

	return c.Error
}

func (handle *ImpAttributionHandle) buildWorkflow() *workflow.WorkFlow {
	wf := workflow.NewWorkFlow()

	encryptTask := workflow.NewTaskNode(handle.encryptAction)
	encryptTwiceTask := workflow.NewTaskNode(handle.encryptTwiceAction)
	firstDecryptTask := workflow.NewTaskNode(handle.firstDecryptAction)
	attributionTask := workflow.NewTaskNode(handle.attributionAction)
	finalDecryptTask := workflow.NewTaskNode(handle.finalDecryptAction)
	sendTask := workflow.NewTaskNode(handle.sendAction)

	wf.AddStartNode(encryptTask)
	wf.AddEdge(encryptTask, encryptTwiceTask)
	wf.AddEdge(encryptTwiceTask, firstDecryptTask)
	wf.AddEdge(firstDecryptTask, attributionTask)
	wf.AddEdge(attributionTask, finalDecryptTask)
	wf.AddEdge(finalDecryptTask, sendTask)
	wf.ConnectToEnd(sendTask)

	return wf
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var SuccessResponse = &Response{
	Code:    0,
	Message: "success",
}

func CreateErrResponse(err error) *Response {
	return &Response{
		Code:    -1,
		Message: err.Error(),
	}
}

func (handle *ImpAttributionHandle) writeResponse(w http.ResponseWriter, resp *Response) {
	bytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
	} else {
		_, _ = w.Write(bytes)
	}
}
