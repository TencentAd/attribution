/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/3/20, 2:26 PM
 */

package crypto

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/TencentAd/attribution/attribution/pkg/common/workflow"
	"github.com/TencentAd/attribution/attribution/pkg/crypto"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/crypto/protocal"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/crypto/safeguard"
)

type HttpHandle struct {
	jobQueue             workflow.JobQueue
	convEncryptSafeguard *safeguard.ConvEncryptSafeguard
}

func NewHttpHandle() *HttpHandle {
	return &HttpHandle{}
}

func (h *HttpHandle) WithJobQueue(jq workflow.JobQueue) *HttpHandle {
	h.jobQueue = jq
	return h
}

func (h *HttpHandle) WithSafeguard(guard *safeguard.ConvEncryptSafeguard) *HttpHandle {
	h.convEncryptSafeguard = guard
	return h
}

func (h *HttpHandle) ServeHttp(w http.ResponseWriter, r *http.Request) {
}

func (h *HttpHandle) doServeHttp(w http.ResponseWriter, r *http.Request) error {
	var err error
	var body []byte
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var req protocal.CryptoRequest
	if err = json.Unmarshal(body, &req); err != nil {
		return err
	}

	err = h.convEncryptSafeguard.Against(req.CampaignId)
	if err != nil {
		return err
	}

	groupId := strconv.FormatInt(req.CampaignId, 10)
	var resp protocal.CryptoResponse
	for _, reqData := range req.Data {
		respData, err := ProcessData(groupId, reqData, crypto.Encrypt)
		if err != nil {
			return err
		}

		resp.Data = append(resp.Data, respData)
	}

	return nil
}

func ProcessData(groupId string, reqData *protocal.RequestData,
	cryptoFunc func(string, string) (string, error)) (*protocal.ResponseData, error) {

	var err error
	var resp protocal.ResponseData
	resp.Imei, err = cryptoFunc(groupId, reqData.Imei)
	if err != nil {
		return nil, err
	}
	resp.Idfa, err = cryptoFunc(groupId, reqData.Idfa)
	if err != nil {
		return nil, err
	}
	resp.AndroidId, err = cryptoFunc(groupId, reqData.AndroidId)
	if err != nil {
		return nil, err
	}
	resp.Oaid, err = cryptoFunc(groupId, reqData.Oaid)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}