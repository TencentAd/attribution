package decrypt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/define"
	"github.com/golang/glog"

	"github.com/TencentAd/attribution/attribution/pkg/common/metricutil"
	"github.com/TencentAd/attribution/attribution/pkg/crypto"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/protocal"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/decrypt/metrics"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/decrypt/safeguard"
)

type HttpHandle struct {
	safeguard *safeguard.DecryptSafeguard
}

func NewHttpHandle() *HttpHandle {
	return &HttpHandle{}
}

func (handle *HttpHandle) WithSafeguard(sg *safeguard.DecryptSafeguard) *HttpHandle {
	handle.safeguard = sg
	return handle
}

func (handle *HttpHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	defer func() {
		metrics.DecryptHttpCost.Observe(metricutil.CalcTimeUsedMilli(startTime))
	}()
	var resp *protocal.CryptoResponse
	var err error
	if resp, err = handle.process(r); err != nil {
		resp = protocal.CreateErrCryptoResponse(err)
		glog.Errorf("failed to decrypt, err: %v", err)
		metrics.DecryptErrCount.Add(1)
	}

	handle.writeResponse(w, resp)
}

func (handle *HttpHandle) process(r *http.Request) (*protocal.CryptoResponse, error) {
	var err error
	var body []byte
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if glog.V(define.VLogLevel) {
		glog.V(define.VLogLevel).Infof("decrypt body: %s", string(body))
	}

	var req protocal.CryptoRequest
	if err = json.Unmarshal(body, &req); err != nil {
		return nil, err
	}

	err = handle.safeguard.Against(req.CampaignId, int64(len(req.Data)))
	if err != nil {
		return nil, err
	}

	groupId := strconv.FormatInt(req.CampaignId, 10)
	var resp protocal.CryptoResponse
	for _, reqData := range req.Data {
		respData, err := protocal.ProcessData(groupId, reqData, crypto.Decrypt)
		if err != nil {
			return nil, err
		}

		resp.Data = append(resp.Data, respData)
	}
	return &resp, nil
}

func (handle *HttpHandle) writeResponse(w http.ResponseWriter, resp *protocal.CryptoResponse) {
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
	} else {
		_, _ = w.Write(data)
	}
}
